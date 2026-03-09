package service

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/util"
	"strings"
	"sync"
	"time"

	"xorm.io/xorm"
)

type OIDCAuthService struct {
	cfg        *config.ServerConfig
	db         *xorm.Engine
	httpClient *http.Client
}

type oidcMetadata struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
}

type oidcTokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type OIDCUserClaims struct {
	Subject string
	Email   string
	Name    string
	Picture string
}

type oidcStateEntry struct {
	RedirectTo  string
	CallbackURL string
	ExpiresAt   time.Time
}

type oidcTicketEntry struct {
	Token     string
	ExpiresAt time.Time
}

type oidcMetadataEntry struct {
	Value     oidcMetadata
	ExpiresAt time.Time
}

type oidcRuntimeStore struct {
	mu       sync.Mutex
	states   map[string]oidcStateEntry
	tickets  map[string]oidcTicketEntry
	metadata map[string]oidcMetadataEntry
}

var globalOIDCRuntimeStore = &oidcRuntimeStore{
	states:   map[string]oidcStateEntry{},
	tickets:  map[string]oidcTicketEntry{},
	metadata: map[string]oidcMetadataEntry{},
}

func NewOIDCAuthService(cfg *config.ServerConfig, db *xorm.Engine) *OIDCAuthService {
	return &OIDCAuthService{
		cfg: cfg,
		db:  db,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (s *OIDCAuthService) IsEnabled() bool {
	if s.cfg == nil || s.cfg.OIDC == nil {
		return false
	}
	if !s.cfg.OIDC.Enabled {
		return false
	}
	return strings.TrimSpace(s.cfg.OIDC.Issuer) != "" &&
		strings.TrimSpace(s.cfg.OIDC.ClientID) != "" &&
		strings.TrimSpace(s.cfg.OIDC.ClientSecret) != ""
}

func (s *OIDCAuthService) BuildAdminAuthURL(requestBaseURL, redirectTo string) (string, bool, error) {
	if !s.IsEnabled() {
		return "", false, nil
	}

	metadata, err := s.getMetadata()
	if err != nil {
		return "", true, err
	}

	callbackURL, err := s.resolveCallbackURL(requestBaseURL)
	if err != nil {
		return "", true, err
	}

	state := randomToken(32)
	if state == "" {
		return "", true, errors.New("failed to generate state")
	}

	redirect := s.normalizeSuccessRedirect(redirectTo)
	s.setState(state, oidcStateEntry{
		RedirectTo:  redirect,
		CallbackURL: callbackURL,
		ExpiresAt:   time.Now().Add(s.stateTTL()),
	})

	query := url.Values{}
	query.Set("client_id", s.cfg.OIDC.ClientID)
	query.Set("response_type", "code")
	query.Set("scope", strings.Join(s.scopes(), " "))
	query.Set("redirect_uri", callbackURL)
	query.Set("state", state)
	if p := strings.TrimSpace(s.cfg.OIDC.Prompt); p != "" {
		query.Set("prompt", p)
	}

	return metadata.AuthorizationEndpoint + "?" + query.Encode(), true, nil
}

func (s *OIDCAuthService) ConsumeAdminCallback(code, state string) (string, string, error) {
	failureRedirect := s.normalizeFailureRedirect("")
	if !s.IsEnabled() {
		return "", failureRedirect, errors.New("oidc disabled")
	}
	if strings.TrimSpace(code) == "" || strings.TrimSpace(state) == "" {
		return "", failureRedirect, errors.New("missing code or state")
	}

	stored, ok := s.popState(state)
	if !ok {
		return "", failureRedirect, errors.New("state invalid or expired")
	}

	tokenResp, err := s.exchangeCode(code, stored.CallbackURL)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	claims, err := s.fetchUserClaims(tokenResp)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	user, err := s.resolveAdminUser(claims)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	token, err := s.issueAdminToken(user)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	ticket := randomToken(24)
	if ticket == "" {
		return "", stored.RedirectTo, errors.New("failed to generate ticket")
	}
	s.setTicket(ticket, oidcTicketEntry{
		Token:     token,
		ExpiresAt: time.Now().Add(s.ticketTTL()),
	})

	return ticket, stored.RedirectTo, nil
}

func (s *OIDCAuthService) ExchangeAdminTicket(ticket string) (string, error) {
	if strings.TrimSpace(ticket) == "" {
		return "", errors.New("ticket required")
	}

	item, ok := s.popTicket(ticket)
	if !ok {
		return "", errors.New("ticket invalid or expired")
	}

	return item.Token, nil
}

func (s *OIDCAuthService) getMetadata() (*oidcMetadata, error) {
	issuer := strings.TrimRight(strings.TrimSpace(s.cfg.OIDC.Issuer), "/")
	if issuer == "" {
		return nil, errors.New("oidc issuer required")
	}

	if meta, ok := s.getCachedMetadata(issuer); ok {
		return &meta, nil
	}

	discoveryURL := issuer + "/.well-known/openid-configuration"
	req, _ := http.NewRequest(http.MethodGet, discoveryURL, nil)
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("oidc discovery failed with status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var metadata oidcMetadata
	if err = json.Unmarshal(body, &metadata); err != nil {
		return nil, err
	}
	if metadata.AuthorizationEndpoint == "" || metadata.TokenEndpoint == "" {
		return nil, errors.New("oidc metadata missing required endpoints")
	}

	s.setCachedMetadata(issuer, metadata)
	return &metadata, nil
}

func (s *OIDCAuthService) exchangeCode(code, callbackURL string) (*oidcTokenResponse, error) {
	metadata, err := s.getMetadata()
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", callbackURL)
	form.Set("client_id", s.cfg.OIDC.ClientID)
	form.Set("client_secret", s.cfg.OIDC.ClientSecret)

	req, _ := http.NewRequest(http.MethodPost, metadata.TokenEndpoint, bytes.NewBufferString(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("token exchange failed with status %d", resp.StatusCode)
	}

	var tokenResp oidcTokenResponse
	if err = json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}
	if tokenResp.AccessToken == "" && tokenResp.IDToken == "" {
		return nil, errors.New("oidc token response missing token")
	}
	return &tokenResp, nil
}

func (s *OIDCAuthService) fetchUserClaims(tokenResp *oidcTokenResponse) (*OIDCUserClaims, error) {
	metadata, err := s.getMetadata()
	if err != nil {
		return nil, err
	}

	claims := map[string]interface{}{}
	if metadata.UserinfoEndpoint != "" && tokenResp.AccessToken != "" {
		if err = s.fillClaimsByUserinfo(metadata.UserinfoEndpoint, tokenResp.AccessToken, claims); err != nil {
			return nil, err
		}
	}
	if len(claims) == 0 && tokenResp.IDToken != "" {
		if err = fillClaimsByIDToken(tokenResp.IDToken, claims); err != nil {
			return nil, err
		}
	}

	subjectClaim := s.claimOrDefault(s.cfg.OIDC.SubjectClaim, "sub")
	emailClaim := s.claimOrDefault(s.cfg.OIDC.EmailClaim, "email")
	nameClaim := s.claimOrDefault(s.cfg.OIDC.NameClaim, "name")
	pictureClaim := s.claimOrDefault(s.cfg.OIDC.PictureClaim, "picture")

	userClaims := &OIDCUserClaims{
		Subject: strings.TrimSpace(anyToString(claims[subjectClaim])),
		Email:   strings.TrimSpace(anyToString(claims[emailClaim])),
		Name:    strings.TrimSpace(anyToString(claims[nameClaim])),
		Picture: strings.TrimSpace(anyToString(claims[pictureClaim])),
	}
	if userClaims.Subject == "" {
		return nil, errors.New("oidc subject claim missing")
	}
	if !s.isAllowedEmailDomain(userClaims.Email) {
		return nil, errors.New("email domain not allowed")
	}
	return userClaims, nil
}

func (s *OIDCAuthService) fillClaimsByUserinfo(userinfoEndpoint, accessToken string, claims map[string]interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, userinfoEndpoint, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("userinfo failed with status %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &claims)
}

func fillClaimsByIDToken(idToken string, claims map[string]interface{}) error {
	parts := strings.Split(idToken, ".")
	if len(parts) < 2 {
		return errors.New("invalid id token")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, &claims)
}

func (s *OIDCAuthService) resolveAdminUser(claims *OIDCUserClaims) (*model.User, error) {
	provider := strings.TrimSpace(s.cfg.OIDC.ProviderName)
	if provider == "" {
		provider = "oidc"
	}

	var account model.OAuthAccount
	has, err := s.db.Where("provider = ? and subject = ? and is_admin = 1 and status = 1", provider, claims.Subject).Get(&account)
	if err != nil {
		return nil, err
	}
	if has {
		var user model.User
		ok, err := s.db.Where("id = ? and is_admin = 1 and status > 0", account.UserId).Get(&user)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("bound admin user not available")
		}
		account.Email = claims.Email
		account.Name = claims.Name
		account.Picture = claims.Picture
		account.LastLoginAt = time.Now()
		_, _ = s.db.Where("id = ?", account.Id).Cols("email", "name", "picture", "last_login_at").Update(&account)
		return &user, nil
	}

	user, err := s.matchOrCreateAdminUser(claims)
	if err != nil {
		return nil, err
	}

	newAccount := &model.OAuthAccount{
		UserId:      user.Id,
		Provider:    provider,
		Subject:     claims.Subject,
		Email:       claims.Email,
		Name:        claims.Name,
		Picture:     claims.Picture,
		IsAdmin:     true,
		Status:      1,
		LastLoginAt: time.Now(),
	}
	_, err = s.db.Insert(newAccount)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *OIDCAuthService) matchOrCreateAdminUser(claims *OIDCUserClaims) (*model.User, error) {
	if s.cfg.OIDC.BindByEmail && claims.Email != "" {
		var user model.User
		has, err := s.db.Where("email = ? and is_admin = 1 and status > 0", claims.Email).Get(&user)
		if err != nil {
			return nil, err
		}
		if has {
			return &user, nil
		}
	}

	if !s.cfg.OIDC.AutoCreateAdmin {
		return nil, errors.New("no bindable admin account")
	}

	nameSeed := claims.Email
	if nameSeed == "" {
		nameSeed = claims.Subject
	}
	username := sanitizeUsername(nameSeed)
	if username == "" {
		username = "oidc_admin"
	}
	uniqueUsername, err := s.makeUniqueUsername(username)
	if err != nil {
		return nil, err
	}
	passwordHash, err := util.Password(randomToken(24))
	if err != nil {
		return nil, err
	}
	displayName := strings.TrimSpace(claims.Name)
	if displayName == "" {
		displayName = uniqueUsername
	}

	user := &model.User{
		Username:            uniqueUsername,
		Password:            passwordHash,
		Name:                displayName,
		Email:               claims.Email,
		LoginVerify:         model.LOGIN_ACCESS_TOKEN,
		TwoFactorAuthSecret: "",
		Note:                "auto-created by oidc",
		LicensedDevices:     0,
		Status:              1,
		IsAdmin:             true,
	}
	_, err = s.db.Insert(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *OIDCAuthService) makeUniqueUsername(base string) (string, error) {
	username := base
	for i := 0; i < 30; i++ {
		cnt, err := s.db.Where("username = ?", username).Count(&model.User{})
		if err != nil {
			return "", err
		}
		if cnt == 0 {
			return username, nil
		}
		username = fmt.Sprintf("%s_%d", base, i+1)
	}
	return "", errors.New("failed to allocate unique username")
}

func (s *OIDCAuthService) issueAdminToken(user *model.User) (string, error) {
	_, _ = s.db.Where("user_id = ? and status = 1 and is_admin = 1", user.Id).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})
	signStr := fmt.Sprintf("%d_%s_%s", user.Id, user.Username, time.Now().String())
	token := util.HmacSha256(signStr, s.cfg.SignKey)
	authToken := &model.AuthToken{
		UserId:  user.Id,
		Token:   token,
		Expired: time.Now().Add(2 * time.Hour),
		IsAdmin: true,
		Status:  1,
	}
	_, err := s.db.Insert(authToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *OIDCAuthService) resolveCallbackURL(requestBaseURL string) (string, error) {
	explicit := strings.TrimSpace(s.cfg.OIDC.RedirectURL)
	if explicit != "" {
		return explicit, nil
	}
	base := strings.TrimSpace(requestBaseURL)
	base = strings.TrimRight(base, "/")
	if base == "" {
		return "", errors.New("oidc redirect url missing and request base unavailable")
	}
	return base + "/admin/auth/oidc/callback", nil
}

func (s *OIDCAuthService) normalizeSuccessRedirect(raw string) string {
	target := strings.TrimSpace(raw)
	if target == "" {
		target = strings.TrimSpace(s.cfg.OIDC.SuccessRedirect)
	}
	if target == "" {
		target = "/login"
	}
	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "//") {
		return "/login"
	}
	if !strings.HasPrefix(target, "/") {
		target = "/" + target
	}
	return target
}

func (s *OIDCAuthService) normalizeFailureRedirect(raw string) string {
	target := strings.TrimSpace(raw)
	if target == "" {
		target = strings.TrimSpace(s.cfg.OIDC.FailureRedirect)
	}
	if target == "" {
		target = "/login"
	}
	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "//") {
		return "/login"
	}
	if !strings.HasPrefix(target, "/") {
		target = "/" + target
	}
	return target
}

func (s *OIDCAuthService) scopes() []string {
	if s.cfg == nil || s.cfg.OIDC == nil || len(s.cfg.OIDC.Scopes) == 0 {
		return []string{"openid", "profile", "email"}
	}
	return s.cfg.OIDC.Scopes
}

func (s *OIDCAuthService) stateTTL() time.Duration {
	if s.cfg == nil || s.cfg.OIDC == nil || s.cfg.OIDC.StateTTLSeconds <= 0 {
		return 180 * time.Second
	}
	return time.Duration(s.cfg.OIDC.StateTTLSeconds) * time.Second
}

func (s *OIDCAuthService) ticketTTL() time.Duration {
	if s.cfg == nil || s.cfg.OIDC == nil || s.cfg.OIDC.TicketTTLSeconds <= 0 {
		return 180 * time.Second
	}
	return time.Duration(s.cfg.OIDC.TicketTTLSeconds) * time.Second
}

func (s *OIDCAuthService) claimOrDefault(value, fallback string) string {
	v := strings.TrimSpace(value)
	if v == "" {
		return fallback
	}
	return v
}

func (s *OIDCAuthService) isAllowedEmailDomain(email string) bool {
	if s.cfg == nil || s.cfg.OIDC == nil || len(s.cfg.OIDC.AllowedEmailDomains) == 0 {
		return true
	}
	i := strings.LastIndex(email, "@")
	if i <= 0 {
		return false
	}
	domain := strings.ToLower(strings.TrimSpace(email[i+1:]))
	for _, allowed := range s.cfg.OIDC.AllowedEmailDomains {
		if strings.ToLower(strings.TrimSpace(allowed)) == domain {
			return true
		}
	}
	return false
}

func (s *OIDCAuthService) setState(key string, value oidcStateEntry) {
	now := time.Now()
	globalOIDCRuntimeStore.mu.Lock()
	defer globalOIDCRuntimeStore.mu.Unlock()
	for k, v := range globalOIDCRuntimeStore.states {
		if now.After(v.ExpiresAt) {
			delete(globalOIDCRuntimeStore.states, k)
		}
	}
	globalOIDCRuntimeStore.states[key] = value
}

func (s *OIDCAuthService) popState(key string) (oidcStateEntry, bool) {
	now := time.Now()
	globalOIDCRuntimeStore.mu.Lock()
	defer globalOIDCRuntimeStore.mu.Unlock()
	v, ok := globalOIDCRuntimeStore.states[key]
	if !ok {
		return oidcStateEntry{}, false
	}
	delete(globalOIDCRuntimeStore.states, key)
	if now.After(v.ExpiresAt) {
		return oidcStateEntry{}, false
	}
	return v, true
}

func (s *OIDCAuthService) setTicket(key string, value oidcTicketEntry) {
	now := time.Now()
	globalOIDCRuntimeStore.mu.Lock()
	defer globalOIDCRuntimeStore.mu.Unlock()
	for k, v := range globalOIDCRuntimeStore.tickets {
		if now.After(v.ExpiresAt) {
			delete(globalOIDCRuntimeStore.tickets, k)
		}
	}
	globalOIDCRuntimeStore.tickets[key] = value
}

func (s *OIDCAuthService) popTicket(key string) (oidcTicketEntry, bool) {
	now := time.Now()
	globalOIDCRuntimeStore.mu.Lock()
	defer globalOIDCRuntimeStore.mu.Unlock()
	v, ok := globalOIDCRuntimeStore.tickets[key]
	if !ok {
		return oidcTicketEntry{}, false
	}
	delete(globalOIDCRuntimeStore.tickets, key)
	if now.After(v.ExpiresAt) {
		return oidcTicketEntry{}, false
	}
	return v, true
}

func (s *OIDCAuthService) getCachedMetadata(issuer string) (oidcMetadata, bool) {
	now := time.Now()
	globalOIDCRuntimeStore.mu.Lock()
	defer globalOIDCRuntimeStore.mu.Unlock()
	v, ok := globalOIDCRuntimeStore.metadata[issuer]
	if !ok || now.After(v.ExpiresAt) {
		if ok {
			delete(globalOIDCRuntimeStore.metadata, issuer)
		}
		return oidcMetadata{}, false
	}
	return v.Value, true
}

func (s *OIDCAuthService) setCachedMetadata(issuer string, metadata oidcMetadata) {
	globalOIDCRuntimeStore.mu.Lock()
	defer globalOIDCRuntimeStore.mu.Unlock()
	globalOIDCRuntimeStore.metadata[issuer] = oidcMetadataEntry{
		Value:     metadata,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

func randomToken(byteLen int) string {
	if byteLen <= 0 {
		return ""
	}
	buf := make([]byte, byteLen)
	_, err := rand.Read(buf)
	if err != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(buf)
}

func sanitizeUsername(seed string) string {
	seed = strings.TrimSpace(strings.ToLower(seed))
	if i := strings.Index(seed, "@"); i > 0 {
		seed = seed[:i]
	}
	var b strings.Builder
	for _, r := range seed {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
		case r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '_' || r == '-' || r == '.':
			b.WriteRune('_')
		}
	}
	out := strings.Trim(b.String(), "_")
	if out == "" {
		return ""
	}
	if len(out) > 32 {
		out = out[:32]
	}
	return out
}

func anyToString(v interface{}) string {
	switch t := v.(type) {
	case string:
		return t
	case float64:
		return fmt.Sprintf("%.0f", t)
	case int:
		return fmt.Sprintf("%d", t)
	case int64:
		return fmt.Sprintf("%d", t)
	case json.Number:
		return t.String()
	default:
		return ""
	}
}
