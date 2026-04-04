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

type OAuthProviderService struct {
	cfg        *config.ServerConfig
	db         *xorm.Engine
	httpClient *http.Client
}

type OAuthProviderMeta struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Type        string `json:"type"`
}

type oauthMetadata struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
}

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type OAuthUserClaims struct {
	Subject string
	Email   string
	Name    string
	Picture string
}

type oauthStateEntry struct {
	ProviderName string
	RedirectTo   string
	CallbackURL  string
	ExpiresAt    time.Time
}

type oauthSignedStatePayload struct {
	ProviderName string `json:"providerName"`
	RedirectTo   string `json:"redirectTo"`
	CallbackURL  string `json:"callbackUrl"`
	ExpiresAt    int64  `json:"expiresAt"`
	Nonce        string `json:"nonce"`
}

type oauthTicketEntry struct {
	Token     string
	ExpiresAt time.Time
}

type oauthMetadataEntry struct {
	Value     oauthMetadata
	ExpiresAt time.Time
}

type oauthRuntimeStore struct {
	mu       sync.Mutex
	states   map[string]oauthStateEntry
	tickets  map[string]oauthTicketEntry
	metadata map[string]oauthMetadataEntry
}

var globalOAuthRuntimeStore = &oauthRuntimeStore{
	states:   map[string]oauthStateEntry{},
	tickets:  map[string]oauthTicketEntry{},
	metadata: map[string]oauthMetadataEntry{},
}

func NewOAuthProviderService(cfg *config.ServerConfig, db *xorm.Engine) *OAuthProviderService {
	return &OAuthProviderService{
		cfg: cfg,
		db:  db,
		httpClient: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (s *OAuthProviderService) ListEnabledProviders() []OAuthProviderMeta {
	metas := make([]OAuthProviderMeta, 0)
	if s == nil || s.cfg == nil {
		return metas
	}
	for _, provider := range s.cfg.OAuthProviders() {
		normalized := normalizeOAuthProvider(provider)
		if !s.isProviderEnabled(normalized) {
			continue
		}
		metas = append(metas, OAuthProviderMeta{
			Name:        normalized.Name,
			DisplayName: normalized.DisplayName,
			Type:        normalized.Type,
		})
	}
	return metas
}

func (s *OAuthProviderService) BuildAdminAuthURL(providerName, requestBaseURL, redirectTo string) (string, bool, error) {
	provider, ok := s.getProvider(providerName)
	if !ok {
		return "", false, nil
	}
	if !s.isProviderEnabled(provider) {
		return "", false, nil
	}

	metadata, err := s.getMetadata(provider)
	if err != nil {
		return "", true, err
	}

	callbackURL, err := s.resolveCallbackURL(provider, requestBaseURL)
	if err != nil {
		return "", true, err
	}

	stateEntry := oauthStateEntry{
		ProviderName: provider.Name,
		RedirectTo:   s.normalizeSuccessRedirect(provider, redirectTo),
		CallbackURL:  callbackURL,
		ExpiresAt:    time.Now().Add(s.stateTTL(provider)),
	}

	state := s.buildSignedState(stateEntry)
	if state == "" {
		return "", true, errors.New("failed to generate state")
	}

	s.setState(state, stateEntry)

	query := url.Values{}
	query.Set("client_id", provider.ClientID)
	query.Set("response_type", "code")
	query.Set("scope", strings.Join(s.scopes(provider), " "))
	query.Set("redirect_uri", callbackURL)
	query.Set("state", state)
	if prompt := strings.TrimSpace(provider.Prompt); prompt != "" {
		query.Set("prompt", prompt)
	}

	return metadata.AuthorizationEndpoint + "?" + query.Encode(), true, nil
}

func (s *OAuthProviderService) ConsumeAdminCallback(providerName, code, state string) (string, string, error) {
	provider, ok := s.getProvider(providerName)
	failureRedirect := s.normalizeFailureRedirect(config.OAuthProviderConfig{}, "")
	if !ok {
		return "", failureRedirect, errors.New("provider not found")
	}
	failureRedirect = s.normalizeFailureRedirect(provider, "")
	if !s.isProviderEnabled(provider) {
		return "", failureRedirect, errors.New("provider disabled")
	}
	if strings.TrimSpace(code) == "" || strings.TrimSpace(state) == "" {
		return "", failureRedirect, errors.New("missing code or state")
	}

	stored, ok := s.popState(state)
	if (!ok || stored.ProviderName != provider.Name) && state != "" {
		if decoded, decodeErr := s.parseSignedState(state); decodeErr == nil {
			stored = decoded
			ok = stored.ProviderName == provider.Name
		}
	}
	if !ok || stored.ProviderName != provider.Name {
		return "", failureRedirect, errors.New("state invalid or expired")
	}

	tokenResp, err := s.exchangeCode(provider, code, stored.CallbackURL)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	claims, err := s.fetchUserClaims(provider, tokenResp)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	user, err := s.resolveAdminUser(provider, claims)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	token, err := s.issueAdminToken(user)
	if err != nil {
		return "", stored.RedirectTo, err
	}

	ticket := randomOAuthToken(24)
	if ticket == "" {
		return "", stored.RedirectTo, errors.New("failed to generate ticket")
	}

	s.setTicket(ticket, oauthTicketEntry{
		Token:     token,
		ExpiresAt: time.Now().Add(s.ticketTTL(provider)),
	})

	return ticket, stored.RedirectTo, nil
}

func (s *OAuthProviderService) ExchangeAdminTicket(ticket string) (string, error) {
	if strings.TrimSpace(ticket) == "" {
		return "", errors.New("ticket required")
	}
	item, ok := s.popTicket(ticket)
	if !ok {
		return "", errors.New("ticket invalid or expired")
	}
	return item.Token, nil
}

func (s *OAuthProviderService) getProvider(name string) (config.OAuthProviderConfig, bool) {
	if s == nil || s.cfg == nil {
		return config.OAuthProviderConfig{}, false
	}
	target := strings.TrimSpace(name)
	if target == "" {
		target = "oidc"
	}
	for _, provider := range s.cfg.OAuthProviders() {
		normalized := normalizeOAuthProvider(provider)
		if normalized.Name == target {
			return normalized, true
		}
	}
	return config.OAuthProviderConfig{}, false
}

func (s *OAuthProviderService) isProviderEnabled(provider config.OAuthProviderConfig) bool {
	if !provider.Enabled {
		return false
	}
	if strings.TrimSpace(provider.ClientID) == "" || strings.TrimSpace(provider.ClientSecret) == "" {
		return false
	}
	if provider.Type == "oidc" || provider.Type == "google" {
		return strings.TrimSpace(provider.Issuer) != "" ||
			(strings.TrimSpace(provider.AuthorizationEndpoint) != "" && strings.TrimSpace(provider.TokenEndpoint) != "")
	}
	return strings.TrimSpace(provider.AuthorizationEndpoint) != "" &&
		strings.TrimSpace(provider.TokenEndpoint) != ""
}

func (s *OAuthProviderService) getMetadata(provider config.OAuthProviderConfig) (*oauthMetadata, error) {
	if provider.AuthorizationEndpoint != "" && provider.TokenEndpoint != "" {
		return &oauthMetadata{
			AuthorizationEndpoint: provider.AuthorizationEndpoint,
			TokenEndpoint:         provider.TokenEndpoint,
			UserinfoEndpoint:      provider.UserinfoEndpoint,
		}, nil
	}

	issuer := strings.TrimRight(strings.TrimSpace(provider.Issuer), "/")
	if issuer == "" {
		return nil, errors.New("oauth issuer required")
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
		return nil, fmt.Errorf("oauth discovery failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var metadata oauthMetadata
	if err = json.Unmarshal(body, &metadata); err != nil {
		return nil, err
	}
	if metadata.AuthorizationEndpoint == "" || metadata.TokenEndpoint == "" {
		return nil, errors.New("oauth metadata missing required endpoints")
	}

	s.setCachedMetadata(issuer, metadata)
	return &metadata, nil
}

func (s *OAuthProviderService) exchangeCode(provider config.OAuthProviderConfig, code, callbackURL string) (*oauthTokenResponse, error) {
	metadata, err := s.getMetadata(provider)
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", callbackURL)
	form.Set("client_id", provider.ClientID)
	form.Set("client_secret", provider.ClientSecret)

	req, _ := http.NewRequest(http.MethodPost, metadata.TokenEndpoint, bytes.NewBufferString(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "rustdesk-api-server-pro")

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

	var tokenResp oauthTokenResponse
	if err = json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}
	if tokenResp.AccessToken == "" && tokenResp.IDToken == "" {
		return nil, errors.New("oauth token response missing token")
	}
	return &tokenResp, nil
}

func (s *OAuthProviderService) fetchUserClaims(provider config.OAuthProviderConfig, tokenResp *oauthTokenResponse) (*OAuthUserClaims, error) {
	metadata, err := s.getMetadata(provider)
	if err != nil {
		return nil, err
	}

	claims := map[string]interface{}{}
	if metadata.UserinfoEndpoint != "" && tokenResp.AccessToken != "" {
		if err = s.fillClaimsByUserinfo(metadata.UserinfoEndpoint, tokenResp.AccessToken, claims); err != nil {
			return nil, err
		}
	}
	if provider.Type == "github" && tokenResp.AccessToken != "" {
		if err = s.fillGithubEmail(tokenResp.AccessToken, claims); err != nil {
			return nil, err
		}
	}
	if len(claims) == 0 && tokenResp.IDToken != "" {
		if err = fillClaimsByOAuthIDToken(tokenResp.IDToken, claims); err != nil {
			return nil, err
		}
	}

	userClaims := &OAuthUserClaims{
		Subject: strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.SubjectClaim, "sub")])),
		Email:   strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.EmailClaim, "email")])),
		Name:    strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.NameClaim, "name")])),
		Picture: strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.PictureClaim, "picture")])),
	}

	if userClaims.Subject == "" {
		return nil, errors.New("oauth subject claim missing")
	}
	if !s.isAllowedEmailDomain(provider, userClaims.Email) {
		return nil, errors.New("email domain not allowed")
	}
	return userClaims, nil
}

func (s *OAuthProviderService) fillClaimsByUserinfo(userinfoEndpoint, accessToken string, claims map[string]interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, userinfoEndpoint, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "rustdesk-api-server-pro")

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

func (s *OAuthProviderService) fillGithubEmail(accessToken string, claims map[string]interface{}) error {
	if email := strings.TrimSpace(anyToOAuthString(claims["email"])); email != "" {
		return nil
	}

	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/user/emails", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "rustdesk-api-server-pro")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil
	}

	var emails []struct {
		Email   string `json:"email"`
		Primary bool   `json:"primary"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return err
	}

	for _, item := range emails {
		if item.Primary && strings.TrimSpace(item.Email) != "" {
			claims["email"] = item.Email
			return nil
		}
	}
	for _, item := range emails {
		if strings.TrimSpace(item.Email) != "" {
			claims["email"] = item.Email
			return nil
		}
	}
	return nil
}

func fillClaimsByOAuthIDToken(idToken string, claims map[string]interface{}) error {
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

func (s *OAuthProviderService) resolveAdminUser(provider config.OAuthProviderConfig, claims *OAuthUserClaims) (*model.User, error) {
	var account model.OAuthAccount
	has, err := s.db.Where("provider = ? and subject = ? and is_admin = 1 and status = 1", provider.Name, claims.Subject).Get(&account)
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

	user, err := s.matchOrCreateAdminUser(provider, claims)
	if err != nil {
		return nil, err
	}

	newAccount := &model.OAuthAccount{
		UserId:      user.Id,
		Provider:    provider.Name,
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

func (s *OAuthProviderService) matchOrCreateAdminUser(provider config.OAuthProviderConfig, claims *OAuthUserClaims) (*model.User, error) {
	if provider.BindByEmail && claims.Email != "" {
		var user model.User
		has, err := s.db.Where("email = ? and is_admin = 1 and status > 0", claims.Email).Get(&user)
		if err != nil {
			return nil, err
		}
		if has {
			return &user, nil
		}
	}

	if !provider.AutoCreateAdmin {
		return nil, errors.New("no bindable admin account")
	}

	nameSeed := claims.Email
	if nameSeed == "" {
		nameSeed = claims.Subject
	}
	username := sanitizeOAuthUsername(nameSeed)
	if username == "" {
		username = provider.Name + "_admin"
	}
	uniqueUsername, err := s.makeUniqueUsername(username)
	if err != nil {
		return nil, err
	}
	passwordHash, err := util.Password(randomOAuthToken(24))
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
		Note:                "auto-created by oauth:" + provider.Name,
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

func (s *OAuthProviderService) makeUniqueUsername(base string) (string, error) {
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

func (s *OAuthProviderService) issueAdminToken(user *model.User) (string, error) {
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

func (s *OAuthProviderService) resolveCallbackURL(provider config.OAuthProviderConfig, requestBaseURL string) (string, error) {
	if explicit := strings.TrimSpace(provider.RedirectURL); explicit != "" {
		return explicit, nil
	}
	base := strings.TrimRight(strings.TrimSpace(requestBaseURL), "/")
	if base == "" {
		return "", errors.New("oauth redirect url missing and request base unavailable")
	}
	return fmt.Sprintf("%s/admin/auth/oauth/%s/callback", base, provider.Name), nil
}

func (s *OAuthProviderService) normalizeSuccessRedirect(provider config.OAuthProviderConfig, raw string) string {
	target := strings.TrimSpace(raw)
	if target == "" {
		target = strings.TrimSpace(provider.SuccessRedirect)
	}
	return normalizeOAuthRedirectTarget(target)
}

func (s *OAuthProviderService) normalizeFailureRedirect(provider config.OAuthProviderConfig, raw string) string {
	target := strings.TrimSpace(raw)
	if target == "" {
		target = strings.TrimSpace(provider.FailureRedirect)
	}
	return normalizeOAuthRedirectTarget(target)
}

func normalizeOAuthRedirectTarget(target string) string {
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

func (s *OAuthProviderService) scopes(provider config.OAuthProviderConfig) []string {
	if len(provider.Scopes) > 0 {
		return provider.Scopes
	}
	switch provider.Type {
	case "github":
		return []string{"read:user", "user:email"}
	default:
		return []string{"openid", "profile", "email"}
	}
}

func (s *OAuthProviderService) stateTTL(provider config.OAuthProviderConfig) time.Duration {
	if provider.StateTTLSeconds <= 0 {
		return 180 * time.Second
	}
	return time.Duration(provider.StateTTLSeconds) * time.Second
}

func (s *OAuthProviderService) ticketTTL(provider config.OAuthProviderConfig) time.Duration {
	if provider.TicketTTLSeconds <= 0 {
		return 180 * time.Second
	}
	return time.Duration(provider.TicketTTLSeconds) * time.Second
}

func (s *OAuthProviderService) isAllowedEmailDomain(provider config.OAuthProviderConfig, email string) bool {
	if len(provider.AllowedEmailDomains) == 0 {
		return true
	}
	i := strings.LastIndex(email, "@")
	if i <= 0 {
		return false
	}
	domain := strings.ToLower(strings.TrimSpace(email[i+1:]))
	for _, allowed := range provider.AllowedEmailDomains {
		if strings.ToLower(strings.TrimSpace(allowed)) == domain {
			return true
		}
	}
	return false
}

func (s *OAuthProviderService) setState(key string, value oauthStateEntry) {
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	for k, v := range globalOAuthRuntimeStore.states {
		if now.After(v.ExpiresAt) {
			delete(globalOAuthRuntimeStore.states, k)
		}
	}
	globalOAuthRuntimeStore.states[key] = value
}

func (s *OAuthProviderService) popState(key string) (oauthStateEntry, bool) {
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	v, ok := globalOAuthRuntimeStore.states[key]
	if !ok {
		return oauthStateEntry{}, false
	}
	delete(globalOAuthRuntimeStore.states, key)
	if now.After(v.ExpiresAt) {
		return oauthStateEntry{}, false
	}
	return v, true
}

func (s *OAuthProviderService) setTicket(key string, value oauthTicketEntry) {
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	for k, v := range globalOAuthRuntimeStore.tickets {
		if now.After(v.ExpiresAt) {
			delete(globalOAuthRuntimeStore.tickets, k)
		}
	}
	globalOAuthRuntimeStore.tickets[key] = value
}

func (s *OAuthProviderService) popTicket(key string) (oauthTicketEntry, bool) {
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	v, ok := globalOAuthRuntimeStore.tickets[key]
	if !ok {
		return oauthTicketEntry{}, false
	}
	delete(globalOAuthRuntimeStore.tickets, key)
	if now.After(v.ExpiresAt) {
		return oauthTicketEntry{}, false
	}
	return v, true
}

func (s *OAuthProviderService) getCachedMetadata(issuer string) (oauthMetadata, bool) {
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	v, ok := globalOAuthRuntimeStore.metadata[issuer]
	if !ok || now.After(v.ExpiresAt) {
		if ok {
			delete(globalOAuthRuntimeStore.metadata, issuer)
		}
		return oauthMetadata{}, false
	}
	return v.Value, true
}

func (s *OAuthProviderService) setCachedMetadata(issuer string, metadata oauthMetadata) {
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	globalOAuthRuntimeStore.metadata[issuer] = oauthMetadataEntry{
		Value:     metadata,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
}

func (s *OAuthProviderService) buildSignedState(entry oauthStateEntry) string {
	payload := oauthSignedStatePayload{
		ProviderName: entry.ProviderName,
		RedirectTo:   entry.RedirectTo,
		CallbackURL:  entry.CallbackURL,
		ExpiresAt:    entry.ExpiresAt.Unix(),
		Nonce:        randomOAuthToken(12),
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString(data)
	signature := util.HmacSha256(encodedPayload, s.cfg.SignKey)
	if signature == "" {
		return ""
	}
	return encodedPayload + "." + base64.RawURLEncoding.EncodeToString([]byte(signature))
}

func (s *OAuthProviderService) parseSignedState(state string) (oauthStateEntry, error) {
	parts := strings.Split(state, ".")
	if len(parts) != 2 {
		return oauthStateEntry{}, errors.New("invalid state format")
	}

	expectedSignature := util.HmacSha256(parts[0], s.cfg.SignKey)
	rawSignature, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return oauthStateEntry{}, err
	}
	if string(rawSignature) != expectedSignature {
		return oauthStateEntry{}, errors.New("state signature mismatch")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return oauthStateEntry{}, err
	}

	var payload oauthSignedStatePayload
	if err = json.Unmarshal(payloadBytes, &payload); err != nil {
		return oauthStateEntry{}, err
	}

	entry := oauthStateEntry{
		ProviderName: strings.TrimSpace(payload.ProviderName),
		RedirectTo:   normalizeOAuthRedirectTarget(payload.RedirectTo),
		CallbackURL:  strings.TrimSpace(payload.CallbackURL),
		ExpiresAt:    time.Unix(payload.ExpiresAt, 0),
	}
	if entry.ProviderName == "" || entry.CallbackURL == "" {
		return oauthStateEntry{}, errors.New("state payload incomplete")
	}
	if time.Now().After(entry.ExpiresAt) {
		return oauthStateEntry{}, errors.New("state expired")
	}
	return entry, nil
}

func normalizeOAuthProvider(provider config.OAuthProviderConfig) config.OAuthProviderConfig {
	provider.Type = strings.TrimSpace(strings.ToLower(provider.Type))
	if provider.Type == "" {
		provider.Type = "oidc"
	}
	provider.Name = strings.TrimSpace(provider.Name)
	if provider.Name == "" {
		switch provider.Type {
		case "github":
			provider.Name = "github"
		case "google":
			provider.Name = "google"
		default:
			provider.Name = "oidc"
		}
	}
	if provider.DisplayName == "" {
		switch provider.Type {
		case "github":
			provider.DisplayName = "GitHub"
		case "google":
			provider.DisplayName = "Google"
		default:
			provider.DisplayName = "OIDC"
		}
	}
	switch provider.Type {
	case "google":
		if provider.Issuer == "" {
			provider.Issuer = "https://accounts.google.com"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"openid", "profile", "email"}
		}
	case "github":
		if provider.AuthorizationEndpoint == "" {
			provider.AuthorizationEndpoint = "https://github.com/login/oauth/authorize"
		}
		if provider.TokenEndpoint == "" {
			provider.TokenEndpoint = "https://github.com/login/oauth/access_token"
		}
		if provider.UserinfoEndpoint == "" {
			provider.UserinfoEndpoint = "https://api.github.com/user"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"read:user", "user:email"}
		}
		if provider.SubjectClaim == "" {
			provider.SubjectClaim = "id"
		}
		if provider.EmailClaim == "" {
			provider.EmailClaim = "email"
		}
		if provider.NameClaim == "" {
			provider.NameClaim = "name"
		}
		if provider.PictureClaim == "" {
			provider.PictureClaim = "avatar_url"
		}
	default:
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"openid", "profile", "email"}
		}
		if provider.SubjectClaim == "" {
			provider.SubjectClaim = "sub"
		}
		if provider.EmailClaim == "" {
			provider.EmailClaim = "email"
		}
		if provider.NameClaim == "" {
			provider.NameClaim = "name"
		}
		if provider.PictureClaim == "" {
			provider.PictureClaim = "picture"
		}
	}
	if provider.StateTTLSeconds <= 0 {
		provider.StateTTLSeconds = 180
	}
	if provider.TicketTTLSeconds <= 0 {
		provider.TicketTTLSeconds = 180
	}
	if provider.SuccessRedirect == "" {
		provider.SuccessRedirect = "/login"
	}
	if provider.FailureRedirect == "" {
		provider.FailureRedirect = "/login"
	}
	return provider
}

func randomOAuthToken(byteLen int) string {
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

func sanitizeOAuthUsername(seed string) string {
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

func anyToOAuthString(v interface{}) string {
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

func defaultIfEmpty(value, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
