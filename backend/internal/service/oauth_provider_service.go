package service

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/util"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kataras/iris/v12"
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
	AccountRole string `json:"accountRole"`
}

type oauthMetadata struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
	JWKSURI               string `json:"jwks_uri"`
}

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	OpenID      string `json:"openid"`
	IDToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type OAuthUserClaims struct {
	Subject       string
	Email         string
	Login         string
	Name          string
	Picture       string
	EmailVerified bool
}

type oauthStateEntry struct {
	ProviderName string
	RedirectTo   string
	CallbackURL  string
	ExpiresAt    time.Time
	CodeVerifier string
	RustdeskId   string
	Uuid         string
	DeviceOs     string
	DeviceType   string
	DeviceName   string
	PollToken    string
}

type oauthSignedStatePayload struct {
	ProviderName string `json:"providerName"`
	RedirectTo   string `json:"redirectTo"`
	CallbackURL  string `json:"callbackUrl"`
	ExpiresAt    int64  `json:"expiresAt"`
	Nonce        string `json:"nonce"`
}

type oauthTicketEntry struct {
	UserID    int
	IsAdmin   bool
	ExpiresAt time.Time
	RustdeskId string
	Uuid       string
	DeviceOs   string
	DeviceType string
	DeviceName string
}

type oauthMetadataEntry struct {
	Value     oauthMetadata
	ExpiresAt time.Time
}

type oauthRuntimeStore struct {
	mu       sync.Mutex
	states   map[string]oauthStateEntry
	tickets  map[string]oauthTicketEntry
	polls    map[string]oauthPollEntry
	metadata map[string]oauthMetadataEntry
}

type oauthPollEntry struct {
	Ticket    string
	ExpiresAt time.Time
	Result    string
}

var globalOAuthRuntimeStore = &oauthRuntimeStore{
	states:   map[string]oauthStateEntry{},
	tickets:  map[string]oauthTicketEntry{},
	polls:    map[string]oauthPollEntry{},
	metadata: map[string]oauthMetadataEntry{},
}

func NewOAuthProviderService(cfg *config.ServerConfig, db *xorm.Engine) *OAuthProviderService {
	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ForceAttemptHTTP2:     false,
		MaxIdleConns:          20,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	return &OAuthProviderService{
		cfg: cfg,
		db:  db,
		httpClient: &http.Client{
			Timeout:   20 * time.Second,
			Transport: transport,
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
			AccountRole: normalized.AccountRole,
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

	stateEntry.CodeVerifier = randomOAuthToken(32)
	state := randomOAuthToken(24)
	if state == "" {
		return "", true, errcode.New(errcode.ERR2005.Code, errcode.ERR2005.Message)
	}

	if err = s.setState(state, stateEntry); err != nil {
		return "", true, err
	}

	query := url.Values{}
	if provider.Type == "wechat" {
		query.Set("appid", provider.ClientID)
	} else {
		query.Set("client_id", provider.ClientID)
	}
	query.Set("response_type", "code")
	scopeSeparator := " "
	if provider.Type == "qq" || provider.Type == "wechat" {
		scopeSeparator = ","
	}
	query.Set("scope", strings.Join(s.scopes(provider), scopeSeparator))
	query.Set("redirect_uri", callbackURL)
	query.Set("state", state)
	// QQ Connect, WeChat and Apple do not support PKCE. The persisted one-time state
	// still protects the callback against CSRF and replay attacks.
	if provider.Type != "qq" && provider.Type != "wechat" && provider.Type != "apple" {
		challenge := sha256.Sum256([]byte(stateEntry.CodeVerifier))
		query.Set("code_challenge", base64.RawURLEncoding.EncodeToString(challenge[:]))
		query.Set("code_challenge_method", "S256")
	}
	if prompt := strings.TrimSpace(provider.Prompt); prompt != "" {
		query.Set("prompt", prompt)
	}

	return metadata.AuthorizationEndpoint + "?" + query.Encode(), true, nil
}

func (s *OAuthProviderService) ConsumeAdminCallback(providerName, code, state string) (string, string, error) {
	provider, ok := s.getProvider(providerName)
	failureRedirect := s.normalizeFailureRedirect(config.OAuthProviderConfig{}, "")
	if !ok {
		return "", failureRedirect, errcode.New(errcode.ERR2001.Code, errcode.ERR2001.Message)
	}
	failureRedirect = s.normalizeFailureRedirect(provider, "")
	if !s.isProviderEnabled(provider) {
		return "", failureRedirect, errcode.New(errcode.ERR2002.Code, errcode.ERR2002.Message)
	}
	if strings.TrimSpace(code) == "" || strings.TrimSpace(state) == "" {
		return "", failureRedirect, errcode.New(errcode.ERR2003.Code, errcode.ERR2003.Message)
	}

	stored, ok := s.popState(state)
	if !ok || stored.ProviderName != provider.Name {
		return "", failureRedirect, errcode.New(errcode.ERR2004.Code, errcode.ERR2004.Message)
	}

	tokenResp, err := s.exchangeCode(provider, code, stored.CallbackURL, stored.CodeVerifier)
	if err != nil {
		return "", failureRedirect, err
	}

	claims, err := s.fetchUserClaims(provider, tokenResp)
	if err != nil {
		return "", failureRedirect, err
	}

	user, err := s.resolveOAuthUser(provider, claims)
	if err != nil {
		return "", failureRedirect, err
	}

	ticket := randomOAuthToken(24)
	if ticket == "" {
		return "", failureRedirect, errcode.New(errcode.ERR2006.Code, errcode.ERR2006.Message)
	}

	if err = s.setTicket(ticket, oauthTicketEntry{
		UserID:    user.Id,
		IsAdmin:   user.IsAdmin,
		ExpiresAt: time.Now().Add(s.ticketTTL(provider)),
	}); err != nil {
		return "", failureRedirect, err
	}

	return ticket, stored.RedirectTo, nil
}

func (s *OAuthProviderService) ExchangeAdminTicket(ticket string) (string, error) {
	if strings.TrimSpace(ticket) == "" {
		return "", errcode.New(errcode.ERR2007.Code, errcode.ERR2007.Message)
	}
	item, ok := s.popTicket(ticket)
	if !ok {
		return "", errcode.New(errcode.ERR2008.Code, errcode.ERR2008.Message)
	}
	var user model.User
	has, err := s.db.Where("id = ? and is_admin = ? and status > 0", item.UserID, item.IsAdmin).Get(&user)
	if err != nil {
		return "", err
	}
	if !has {
		return "", errcode.New(errcode.ERR2009.Code, errcode.ERR2009.Message)
	}
	return s.issueOAuthToken(&user)
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
	if strings.TrimSpace(provider.ClientID) == "" {
		return false
	}
	if provider.Type == "apple" {
		return strings.TrimSpace(provider.TeamID) != "" &&
			strings.TrimSpace(provider.KeyID) != "" &&
			strings.TrimSpace(provider.PrivateKey) != ""
	}
	if strings.TrimSpace(provider.ClientSecret) == "" {
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
			JWKSURI:               provider.JWKSURI,
		}, nil
	}

	issuer := strings.TrimRight(strings.TrimSpace(provider.Issuer), "/")
	if issuer == "" {
		return nil, errcode.New(errcode.ERR2010.Code, errcode.ERR2010.Message)
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
		return nil, errcode.Errorf(errcode.ERR2030.Code, errcode.ERR2030.Message, resp.StatusCode)
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
		return nil, errcode.New(errcode.ERR2011.Code, errcode.ERR2011.Message)
	}

	s.setCachedMetadata(issuer, metadata)
	return &metadata, nil
}

func (s *OAuthProviderService) exchangeCode(provider config.OAuthProviderConfig, code, callbackURL, codeVerifier string) (*oauthTokenResponse, error) {
	metadata, err := s.getMetadata(provider)
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set("grant_type", "authorization_code")
	form.Set("code", code)
	form.Set("redirect_uri", callbackURL)
	if provider.Type == "wechat" {
		form.Set("appid", provider.ClientID)
		form.Set("secret", provider.ClientSecret)
	} else {
		form.Set("client_id", provider.ClientID)
		form.Set("client_secret", provider.ClientSecret)
	}
	if provider.Type != "qq" && provider.Type != "wechat" && provider.Type != "apple" {
		form.Set("code_verifier", codeVerifier)
	}
	if provider.Type == "apple" {
		clientSecret, jwtErr := generateAppleClientSecretJWT(provider)
		if jwtErr != nil {
			return nil, errcode.New(errcode.ERR2035.Code, errcode.ERR2035.Message)
		}
		form.Set("client_id", provider.ClientID)
		form.Set("client_secret", clientSecret)
	}
	if provider.Type == "qq" {
		form.Set("fmt", "json")
		form.Set("need_openid", "1")
		endpoint, parseErr := url.Parse(metadata.TokenEndpoint)
		if parseErr != nil {
			return nil, parseErr
		}
		endpoint.RawQuery = form.Encode()
		req, _ := http.NewRequest(http.MethodGet, endpoint.String(), nil)
		return s.doTokenRequest(req)
	}
	if provider.Type == "wechat" {
		endpoint, parseErr := url.Parse(metadata.TokenEndpoint)
		if parseErr != nil {
			return nil, parseErr
		}
		endpoint.RawQuery = form.Encode()
		req, _ := http.NewRequest(http.MethodGet, endpoint.String(), nil)
		return s.doTokenRequest(req)
	}

	req, _ := http.NewRequest(http.MethodPost, metadata.TokenEndpoint, bytes.NewBufferString(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return s.doTokenRequest(req)
}

func (s *OAuthProviderService) doTokenRequest(req *http.Request) (*oauthTokenResponse, error) {
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
		return nil, errcode.Errorf(errcode.ERR2031.Code, errcode.ERR2031.Message, resp.StatusCode)
	}

	var tokenResp oauthTokenResponse
	if err = json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}
	if tokenResp.AccessToken == "" && tokenResp.IDToken == "" {
		return nil, errcode.New(errcode.ERR2012.Code, errcode.ERR2012.Message)
	}
	return &tokenResp, nil
}

func (s *OAuthProviderService) fetchUserClaims(provider config.OAuthProviderConfig, tokenResp *oauthTokenResponse) (*OAuthUserClaims, error) {
	metadata, err := s.getMetadata(provider)
	if err != nil {
		return nil, err
	}

	claims := map[string]interface{}{}
	idTokenClaims := map[string]interface{}{}
	if tokenResp.IDToken != "" {
		if err = s.verifyOAuthIDToken(provider, metadata, tokenResp.IDToken, idTokenClaims); err != nil {
			return nil, err
		}
	}

	if provider.Type == "qq" && tokenResp.AccessToken != "" {
		if err = s.fillQQClaims(provider, tokenResp, claims); err != nil {
			return nil, err
		}
	} else if provider.Type == "wechat" && tokenResp.AccessToken != "" {
		if err = s.fillWechatClaims(provider, tokenResp, claims); err != nil {
			return nil, err
		}
	} else if provider.Type == "apple" {
		if err = s.fillAppleClaims(provider, idTokenClaims, claims); err != nil {
			return nil, err
		}
	} else if metadata.UserinfoEndpoint != "" && tokenResp.AccessToken != "" {
		if err = s.fillClaimsByUserinfo(provider, metadata.UserinfoEndpoint, tokenResp.AccessToken, claims); err != nil {
			return nil, err
		}
	}
	if provider.Type == "github" && tokenResp.AccessToken != "" {
		if err = s.fillGithubEmail(provider, tokenResp.AccessToken, claims); err != nil {
			return nil, err
		}
	}
	if len(claims) == 0 && len(idTokenClaims) > 0 {
		claims = idTokenClaims
	}
	if len(claims) > 0 && len(idTokenClaims) > 0 {
		subjectClaim := defaultIfEmpty(provider.SubjectClaim, "sub")
		userinfoSubject := strings.TrimSpace(anyToOAuthString(claims[subjectClaim]))
		idTokenSubject := strings.TrimSpace(anyToOAuthString(idTokenClaims[subjectClaim]))
		if userinfoSubject != "" && idTokenSubject != "" && userinfoSubject != idTokenSubject {
			return nil, errcode.New(errcode.ERR2013.Code, errcode.ERR2013.Message)
		}
	}

	userClaims := &OAuthUserClaims{
		Subject: strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.SubjectClaim, "sub")])),
		Email:   strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.EmailClaim, "email")])),
		Login:   strings.TrimSpace(anyToOAuthString(claims["login"])),
		Name:    strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.NameClaim, "name")])),
		Picture: strings.TrimSpace(anyToOAuthString(claims[defaultIfEmpty(provider.PictureClaim, "picture")])),
	}
	userClaims.EmailVerified, _ = claims["email_verified"].(bool)

	if userClaims.Subject == "" {
		return nil, errcode.New(errcode.ERR2014.Code, errcode.ERR2014.Message)
	}
	if provider.Type == "github" && provider.BindByEmail && !userClaims.EmailVerified {
		return nil, errcode.New(errcode.ERR2015.Code, errcode.ERR2015.Message)
	}
	if (provider.BindByEmail || len(provider.AllowedEmailDomains) > 0) && userClaims.Email == "" {
		return nil, errcode.New(errcode.ERR2016.Code, errcode.ERR2016.Message)
	}
	if !s.isAllowedEmailDomain(provider, userClaims.Email) {
		return nil, errcode.New(errcode.ERR2017.Code, errcode.ERR2017.Message)
	}
	return userClaims, nil
}

func (s *OAuthProviderService) fillQQClaims(provider config.OAuthProviderConfig, tokenResp *oauthTokenResponse, claims map[string]interface{}) error {
	graphBase := "https://graph.qq.com"
	userinfoEndpoint := strings.TrimSpace(provider.UserinfoEndpoint)
	if parsed, err := url.Parse(userinfoEndpoint); err == nil && parsed.Scheme != "" && parsed.Host != "" {
		graphBase = parsed.Scheme + "://" + parsed.Host
	}
	openid := strings.TrimSpace(tokenResp.OpenID)
	if openid == "" {
		meURL := graphBase + "/oauth2.0/me?fmt=json&access_token=" + url.QueryEscape(tokenResp.AccessToken)
		var me struct {
			OpenID string `json:"openid"`
			Error  int    `json:"error"`
		}
		if err := s.getQQJSON(meURL, &me); err != nil {
			return err
		}
		if me.Error != 0 || strings.TrimSpace(me.OpenID) == "" {
			return errcode.New(errcode.ERR2018.Code, errcode.ERR2018.Message)
		}
		openid = me.OpenID
	}
	if userinfoEndpoint == "" {
		userinfoEndpoint = graphBase + "/user/get_user_info"
	}
	query := url.Values{
		"access_token":       {tokenResp.AccessToken},
		"oauth_consumer_key": {provider.ClientID},
		"openid":             {openid},
		"fmt":                {"json"},
	}
	separator := "?"
	if strings.Contains(userinfoEndpoint, "?") {
		separator = "&"
	}
	var profile struct {
		Ret       int    `json:"ret"`
		Msg       string `json:"msg"`
		Nickname  string `json:"nickname"`
		FigureURL string `json:"figureurl_qq_2"`
		Figure40  string `json:"figureurl_qq_1"`
	}
	if err := s.getQQJSON(userinfoEndpoint+separator+query.Encode(), &profile); err != nil {
		return err
	}
	if profile.Ret != 0 {
		return errcode.Errorf(errcode.ERR2032.Code, errcode.ERR2032.Message, profile.Ret)
	}
	claims[defaultIfEmpty(provider.SubjectClaim, "openid")] = openid
	claims[defaultIfEmpty(provider.NameClaim, "nickname")] = profile.Nickname
	picture := profile.FigureURL
	if picture == "" {
		picture = profile.Figure40
	}
	claims[defaultIfEmpty(provider.PictureClaim, "figureurl_qq_2")] = picture
	return nil
}

func (s *OAuthProviderService) fillWechatClaims(provider config.OAuthProviderConfig, tokenResp *oauthTokenResponse, claims map[string]interface{}) error {
	openid := strings.TrimSpace(tokenResp.OpenID)
	if openid == "" {
		return errcode.New(errcode.ERR2018.Code, errcode.ERR2018.Message)
	}
	userinfoEndpoint := strings.TrimSpace(provider.UserinfoEndpoint)
	if userinfoEndpoint == "" {
		userinfoEndpoint = "https://api.weixin.qq.com/sns/userinfo"
	}
	query := url.Values{
		"access_token": {tokenResp.AccessToken},
		"openid":       {openid},
	}
	separator := "?"
	if strings.Contains(userinfoEndpoint, "?") {
		separator = "&"
	}
	var profile struct {
		OpenID     string `json:"openid"`
		Nickname   string `json:"nickname"`
		HeadImgURL string `json:"headimgurl"`
		ErrCode    int    `json:"errcode"`
	}
	if err := s.getQQJSON(userinfoEndpoint+separator+query.Encode(), &profile); err != nil {
		return err
	}
	if profile.ErrCode != 0 {
		return errcode.Errorf(errcode.ERR2032.Code, errcode.ERR2032.Message, profile.ErrCode)
	}
	claims[defaultIfEmpty(provider.SubjectClaim, "openid")] = openid
	if profile.Nickname != "" {
		claims[defaultIfEmpty(provider.NameClaim, "nickname")] = profile.Nickname
	}
	if profile.HeadImgURL != "" {
		claims[defaultIfEmpty(provider.PictureClaim, "headimgurl")] = profile.HeadImgURL
	}
	return nil
}

func (s *OAuthProviderService) fillAppleClaims(provider config.OAuthProviderConfig, idTokenClaims map[string]interface{}, claims map[string]interface{}) error {
	sub := strings.TrimSpace(anyToOAuthString(idTokenClaims["sub"]))
	if sub == "" {
		return errcode.New(errcode.ERR2014.Code, errcode.ERR2014.Message)
	}
	claims[defaultIfEmpty(provider.SubjectClaim, "sub")] = sub
	if email := anyToOAuthString(idTokenClaims["email"]); email != "" {
		claims[defaultIfEmpty(provider.EmailClaim, "email")] = email
	}
	if verified, ok := idTokenClaims["email_verified"].(bool); ok {
		claims["email_verified"] = verified
	}
	return nil
}

func generateAppleClientSecretJWT(provider config.OAuthProviderConfig) (string, error) {
	teamID := strings.TrimSpace(provider.TeamID)
	keyID := strings.TrimSpace(provider.KeyID)
	privateKeyPEM := strings.TrimSpace(provider.PrivateKey)
	if teamID == "" || keyID == "" || privateKeyPEM == "" {
		return "", fmt.Errorf("apple: teamId, keyId and privateKey are required")
	}
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", fmt.Errorf("apple: failed to decode PEM private key")
	}
	key, keyErr := x509.ParsePKCS8PrivateKey(block.Bytes)
	if keyErr != nil {
		return "", fmt.Errorf("apple: failed to parse private key: %w", keyErr)
	}
	ecKey, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("apple: private key must be ECDSA (P-256)")
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"iss": teamID,
		"iat": now.Unix(),
		"exp": now.Add(180 * 24 * time.Hour).Unix(),
		"aud": "https://appleid.apple.com",
		"sub": provider.ClientID,
	})
	token.Header["kid"] = keyID
	signed, signErr := token.SignedString(ecKey)
	if signErr != nil {
		return "", fmt.Errorf("apple: failed to sign JWT: %w", signErr)
	}
	return signed, nil
}

func (s *OAuthProviderService) getQQJSON(endpoint string, target interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "rustdesk-api-server-pro")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errcode.Errorf(errcode.ERR2033.Code, errcode.ERR2033.Message, resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(target)
}

func (s *OAuthProviderService) verifyOAuthIDToken(provider config.OAuthProviderConfig, metadata *oauthMetadata, idToken string, claims map[string]interface{}) error {
	expectedIssuer := strings.TrimRight(strings.TrimSpace(provider.Issuer), "/")
	if expectedIssuer == "" {
		return errcode.New(errcode.ERR2019.Code, errcode.ERR2019.Message)
	}
	if metadata == nil || strings.TrimSpace(metadata.JWKSURI) == "" {
		return errcode.New(errcode.ERR2020.Code, errcode.ERR2020.Message)
	}
	oidcVerifier := &OIDCAuthService{httpClient: s.httpClient}
	if err := oidcVerifier.verifyIDTokenSignature(idToken, &oidcMetadata{JWKSURI: metadata.JWKSURI}); err != nil {
		return err
	}
	return fillClaimsByOAuthIDToken(idToken, expectedIssuer, provider.ClientID, claims)
}

func (s *OAuthProviderService) fillClaimsByUserinfo(provider config.OAuthProviderConfig, userinfoEndpoint, accessToken string, claims map[string]interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, userinfoEndpoint, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/json")
	if provider.Type == "github" {
		req.Header.Set("Accept", "application/vnd.github+json")
		req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	}
	req.Header.Set("User-Agent", "rustdesk-api-server-pro")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return errcode.Errorf(errcode.ERR2034.Code, errcode.ERR2034.Message, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, &claims)
}

func (s *OAuthProviderService) fillGithubEmail(provider config.OAuthProviderConfig, accessToken string, claims map[string]interface{}) error {
	userinfoEndpoint := strings.TrimRight(strings.TrimSpace(provider.UserinfoEndpoint), "/")
	apiBase := strings.TrimSuffix(userinfoEndpoint, "/user")
	if apiBase == "" {
		apiBase = "https://api.github.com"
	}
	emailEndpoint := apiBase + "/user/emails"
	req, _ := http.NewRequest(http.MethodGet, emailEndpoint, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
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
		Email    string `json:"email"`
		Primary  bool   `json:"primary"`
		Verified bool   `json:"verified"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&emails); err != nil {
		return err
	}

	for _, item := range emails {
		if item.Primary && item.Verified && strings.TrimSpace(item.Email) != "" {
			claims["email"] = item.Email
			claims["email_verified"] = true
			return nil
		}
	}
	for _, item := range emails {
		if item.Verified && strings.TrimSpace(item.Email) != "" {
			claims["email"] = item.Email
			claims["email_verified"] = true
			return nil
		}
	}
	return nil
}

func fillClaimsByOAuthIDToken(idToken, expectedIssuer, expectedAudience string, claims map[string]interface{}) error {
	parts := strings.Split(idToken, ".")
	if len(parts) != 3 {
		return errcode.New(errcode.ERR2021.Code, errcode.ERR2021.Message)
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return err
	}
	if err = json.Unmarshal(payload, &claims); err != nil {
		return err
	}
	return validateIDTokenClaims(claims, expectedIssuer, expectedAudience)
}

func (s *OAuthProviderService) resolveOAuthUser(provider config.OAuthProviderConfig, claims *OAuthUserClaims) (*model.User, error) {
	isAdmin := provider.AccountRole != "user"
	var account model.OAuthAccount
	has, err := s.db.Where("provider = ? and subject = ? and is_admin = ? and status = 1", provider.Name, claims.Subject, isAdmin).Get(&account)
	if err != nil {
		return nil, err
	}
	if has {
		var user model.User
		ok, err := s.db.Where("id = ? and is_admin = ? and status > 0", account.UserId, isAdmin).Get(&user)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errcode.New(errcode.ERR2022.Code, errcode.ERR2022.Message)
		}
		account.Email = claims.Email
		account.Name = claims.Name
		account.Picture = claims.Picture
		account.LastLoginAt = time.Now()
		_, _ = s.db.Where("id = ?", account.Id).Cols("email", "name", "picture", "last_login_at").Update(&account)
		return &user, nil
	}

	user, err := s.matchOrCreateOAuthUser(provider, claims, isAdmin)
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
		IsAdmin:     isAdmin,
		Status:      1,
		LastLoginAt: time.Now(),
	}
	_, err = s.db.Insert(newAccount)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *OAuthProviderService) matchOrCreateOAuthUser(provider config.OAuthProviderConfig, claims *OAuthUserClaims, isAdmin bool) (*model.User, error) {
	if provider.BindByEmail && claims.Email != "" {
		var user model.User
		has, err := s.db.Where("email = ? and is_admin = ? and status > 0", claims.Email, isAdmin).Get(&user)
		if err != nil {
			return nil, err
		}
		if has {
			return &user, nil
		}
	}

	if (isAdmin && !provider.AutoCreateAdmin) || (!isAdmin && !provider.AutoCreateUser) {
		return nil, errcode.New(errcode.ERR2023.Code, errcode.ERR2023.Message)
	}

	nameSeed := claims.Login
	if nameSeed == "" {
		nameSeed = claims.Email
	}
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
		IsAdmin:             isAdmin,
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
	return "", errcode.New(errcode.ERR2024.Code, errcode.ERR2024.Message)
}

func (s *OAuthProviderService) issueOAuthToken(user *model.User) (string, error) {
	_, _ = s.db.Where("user_id = ? and status = 1 and is_admin = ?", user.Id, user.IsAdmin).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})
	signStr := fmt.Sprintf("%d_%s_%s", user.Id, user.Username, time.Now().String())
	token := util.HmacSha256(signStr, s.cfg.SignKey)
	authToken := &model.AuthToken{
		UserId:    user.Id,
		TokenHash: util.Sha256Hex(token),
		Expired:   time.Now().Add(2 * time.Hour),
		IsAdmin:   user.IsAdmin,
		Status:    1,
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
		return "", errcode.New(errcode.ERR2025.Code, errcode.ERR2025.Message)
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
		target = "/#/login"
	}
	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "//") {
		return "/#/login"
	}
	if !strings.HasPrefix(target, "/") {
		target = "/" + target
	}
	// The admin SPA uses hash routing. A router path such as /system/oauth
	// must therefore be returned as /#/system/oauth; otherwise the browser
	// requests it from the Go server and receives a 404.
	if !strings.HasPrefix(target, "/#/") {
		target = "/#" + target
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
	case "qq":
		return []string{"get_user_info"}
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

func (s *OAuthProviderService) setState(key string, value oauthStateEntry) error {
	if s.db != nil {
		_, _ = s.db.Where("expires_at < ? or status = 0", time.Now()).Delete(&model.OAuthLoginSession{})
		_, err := s.db.Insert(&model.OAuthLoginSession{Kind: "state", KeyHash: util.Sha256Hex(key), Provider: value.ProviderName, RedirectTo: value.RedirectTo, CallbackURL: value.CallbackURL, CodeVerifier: value.CodeVerifier, RustdeskId: value.RustdeskId, Uuid: value.Uuid, DeviceOs: value.DeviceOs, DeviceType: value.DeviceType, DeviceName: value.DeviceName, PollToken: value.PollToken, ExpiresAt: value.ExpiresAt, Status: 1})
		return err
	}
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	for k, v := range globalOAuthRuntimeStore.states {
		if now.After(v.ExpiresAt) {
			delete(globalOAuthRuntimeStore.states, k)
		}
	}
	globalOAuthRuntimeStore.states[key] = value
	return nil
}

func (s *OAuthProviderService) popState(key string) (oauthStateEntry, bool) {
	if s.db != nil {
		var session model.OAuthLoginSession
		has, err := s.db.Where("kind = ? and key_hash = ? and status = 1 and expires_at > ?", "state", util.Sha256Hex(key), time.Now()).Get(&session)
		if err != nil || !has {
			return oauthStateEntry{}, false
		}
		updated, err := s.db.ID(session.Id).Where("status = 1").Cols("status").Update(&model.OAuthLoginSession{Status: 0})
		if err != nil || updated != 1 {
			return oauthStateEntry{}, false
		}
		return oauthStateEntry{ProviderName: session.Provider, RedirectTo: session.RedirectTo, CallbackURL: session.CallbackURL, CodeVerifier: session.CodeVerifier, RustdeskId: session.RustdeskId, Uuid: session.Uuid, DeviceOs: session.DeviceOs, DeviceType: session.DeviceType, DeviceName: session.DeviceName, PollToken: session.PollToken, ExpiresAt: session.ExpiresAt}, true
	}
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

func (s *OAuthProviderService) setTicket(key string, value oauthTicketEntry) error {
	if s.db != nil {
		_, err := s.db.Insert(&model.OAuthLoginSession{Kind: "ticket", KeyHash: util.Sha256Hex(key), UserId: value.UserID, IsAdmin: value.IsAdmin, RustdeskId: value.RustdeskId, Uuid: value.Uuid, DeviceOs: value.DeviceOs, DeviceType: value.DeviceType, DeviceName: value.DeviceName, ExpiresAt: value.ExpiresAt, Status: 1})
		return err
	}
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	for k, v := range globalOAuthRuntimeStore.tickets {
		if now.After(v.ExpiresAt) {
			delete(globalOAuthRuntimeStore.tickets, k)
		}
	}
	globalOAuthRuntimeStore.tickets[key] = value
	return nil
}

func (s *OAuthProviderService) popTicket(key string) (oauthTicketEntry, bool) {
	if s.db != nil {
		var session model.OAuthLoginSession
		has, err := s.db.Where("kind = ? and key_hash = ? and status = 1 and expires_at > ?", "ticket", util.Sha256Hex(key), time.Now()).Get(&session)
		if err != nil || !has {
			return oauthTicketEntry{}, false
		}
		updated, err := s.db.ID(session.Id).Where("status = 1").Cols("status").Update(&model.OAuthLoginSession{Status: 0})
		if err != nil || updated != 1 {
			return oauthTicketEntry{}, false
		}
		return oauthTicketEntry{UserID: session.UserId, IsAdmin: session.IsAdmin, RustdeskId: session.RustdeskId, Uuid: session.Uuid, DeviceOs: session.DeviceOs, DeviceType: session.DeviceType, DeviceName: session.DeviceName, ExpiresAt: session.ExpiresAt}, true
	}
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
		return oauthStateEntry{}, errcode.New(errcode.ERR2026.Code, errcode.ERR2026.Message)
	}

	expectedSignature := util.HmacSha256(parts[0], s.cfg.SignKey)
	rawSignature, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return oauthStateEntry{}, err
	}
	if !util.ConstantTimeStringEqual(string(rawSignature), expectedSignature) {
		return oauthStateEntry{}, errcode.New(errcode.ERR2027.Code, errcode.ERR2027.Message)
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
		return oauthStateEntry{}, errcode.New(errcode.ERR2028.Code, errcode.ERR2028.Message)
	}
	if time.Now().After(entry.ExpiresAt) {
		return oauthStateEntry{}, errcode.New(errcode.ERR2029.Code, errcode.ERR2029.Message)
	}
	return entry, nil
}

func normalizeOAuthProvider(provider config.OAuthProviderConfig) config.OAuthProviderConfig {
	provider.Type = strings.TrimSpace(strings.ToLower(provider.Type))
	if provider.Type == "" {
		provider.Type = "oidc"
	}
	provider.Name = strings.TrimSpace(provider.Name)
	provider.AccountRole = strings.ToLower(strings.TrimSpace(provider.AccountRole))
	if provider.AccountRole != "user" {
		provider.AccountRole = "admin"
	}
	if provider.Name == "" {
		switch provider.Type {
		case "github":
			provider.Name = "github"
		case "google":
			provider.Name = "google"
		case "qq":
			provider.Name = "qq"
		case "microsoft":
			provider.Name = "microsoft"
		case "gitee":
			provider.Name = "gitee"
		case "gitlab":
			provider.Name = "gitlab"
		case "apple":
			provider.Name = "apple"
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
		case "qq":
			provider.DisplayName = "QQ"
		case "microsoft":
			provider.DisplayName = "Microsoft"
		case "gitee":
			provider.DisplayName = "Gitee"
		case "gitlab":
			provider.DisplayName = "GitLab"
		case "apple":
			provider.DisplayName = "Apple"
		default:
			provider.DisplayName = "OIDC"
		}
	}
	if provider.Type == "google" {
		if provider.Issuer == "" {
			provider.Issuer = "https://accounts.google.com"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"openid", "profile", "email"}
		}
	} else if provider.Type == "microsoft" {
		if provider.Issuer == "" {
			provider.Issuer = "https://login.microsoftonline.com/common/v2.0"
		}
		if provider.AuthorizationEndpoint == "" {
			provider.AuthorizationEndpoint = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
		}
		if provider.TokenEndpoint == "" {
			provider.TokenEndpoint = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
		}
		if provider.UserinfoEndpoint == "" {
			provider.UserinfoEndpoint = "https://graph.microsoft.com/oidc/userinfo"
		}
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
	} else if provider.Type == "github" {
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
	} else if provider.Type == "gitee" {
		if provider.AuthorizationEndpoint == "" {
			provider.AuthorizationEndpoint = "https://gitee.com/oauth/authorize"
		}
		if provider.TokenEndpoint == "" {
			provider.TokenEndpoint = "https://gitee.com/oauth/token"
		}
		if provider.UserinfoEndpoint == "" {
			provider.UserinfoEndpoint = "https://gitee.com/api/v5/user"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"user_info"}
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
	} else if provider.Type == "gitlab" {
		if provider.AuthorizationEndpoint == "" {
			provider.AuthorizationEndpoint = "https://gitlab.com/oauth/authorize"
		}
		if provider.TokenEndpoint == "" {
			provider.TokenEndpoint = "https://gitlab.com/oauth/token"
		}
		if provider.UserinfoEndpoint == "" {
			provider.UserinfoEndpoint = "https://gitlab.com/api/v4/user"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"read_user"}
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
	} else if provider.Type == "qq" {
		if provider.AuthorizationEndpoint == "" {
			provider.AuthorizationEndpoint = "https://graph.qq.com/oauth2.0/authorize"
		}
		if provider.TokenEndpoint == "" {
			provider.TokenEndpoint = "https://graph.qq.com/oauth2.0/token"
		}
		if provider.UserinfoEndpoint == "" {
			provider.UserinfoEndpoint = "https://graph.qq.com/user/get_user_info"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"get_user_info"}
		}
		provider.BindByEmail = false
		provider.AllowedEmailDomains = nil
		provider.SubjectClaim = "openid"
		provider.NameClaim = "nickname"
		provider.PictureClaim = "figureurl_qq_2"
	} else if provider.Type == "wechat" {
		if provider.AuthorizationEndpoint == "" {
			provider.AuthorizationEndpoint = "https://open.weixin.qq.com/connect/qrconnect"
		}
		if provider.TokenEndpoint == "" {
			provider.TokenEndpoint = "https://api.weixin.qq.com/sns/oauth2/access_token"
		}
		if provider.UserinfoEndpoint == "" {
			provider.UserinfoEndpoint = "https://api.weixin.qq.com/sns/userinfo"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"snsapi_login"}
		}
		provider.BindByEmail = false
		provider.AllowedEmailDomains = nil
		provider.SubjectClaim = "openid"
		provider.NameClaim = "nickname"
		provider.PictureClaim = "headimgurl"
	} else if provider.Type == "apple" {
		if provider.Issuer == "" {
			provider.Issuer = "https://appleid.apple.com"
		}
		if provider.AuthorizationEndpoint == "" {
			provider.AuthorizationEndpoint = "https://appleid.apple.com/auth/authorize"
		}
		if provider.TokenEndpoint == "" {
			provider.TokenEndpoint = "https://appleid.apple.com/auth/token"
		}
		if provider.JWKSURI == "" {
			provider.JWKSURI = "https://appleid.apple.com/auth/keys"
		}
		if len(provider.Scopes) == 0 {
			provider.Scopes = []string{"email", "name"}
		}
		if provider.SubjectClaim == "" {
			provider.SubjectClaim = "sub"
		}
		if provider.EmailClaim == "" {
			provider.EmailClaim = "email"
		}
	} else {
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
		provider.SuccessRedirect = "/#/login"
	}
	if provider.FailureRedirect == "" {
		provider.FailureRedirect = "/#/login"
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

// ListClientProviders returns enabled OAuth providers available for client login.
// Only providers with accountRole=user are exposed to desktop/mobile clients.
func (s *OAuthProviderService) ListClientProviders() []OAuthProviderMeta {
	metas := make([]OAuthProviderMeta, 0)
	if s == nil || s.cfg == nil {
		return metas
	}
	for _, provider := range s.cfg.OAuthProviders() {
		normalized := normalizeOAuthProvider(provider)
		if !s.isProviderEnabled(normalized) {
			continue
		}
		if normalized.AccountRole != "user" {
			continue
		}
		metas = append(metas, OAuthProviderMeta{
			Name:        normalized.Name,
			DisplayName: normalized.DisplayName,
			Type:        normalized.Type,
			AccountRole: normalized.AccountRole,
		})
	}
	return metas
}

// BuildClientAuthURL builds the OAuth authorization URL for a desktop/mobile client.
// It returns the authorization URL, a one-time poll token the client uses to poll
// for the login ticket, whether the provider is enabled, and any error.
func (s *OAuthProviderService) BuildClientAuthURL(providerName, requestBaseURL, rustdeskId, uuid, deviceOs, deviceType, deviceName string) (string, string, bool, error) {
	provider, ok := s.getProvider(providerName)
	if !ok {
		return "", "", false, nil
	}
	if !s.isProviderEnabled(provider) {
		return "", "", false, nil
	}
	if provider.AccountRole != "user" {
		return "", "", false, nil
	}

	metadata, err := s.getMetadata(provider)
	if err != nil {
		return "", "", true, err
	}

	callbackURL, err := s.resolveClientCallbackURL(provider, requestBaseURL)
	if err != nil {
		return "", "", true, err
	}

	pollToken := randomOAuthToken(24)
	if pollToken == "" {
		return "", "", true, errcode.New(errcode.ERR2207.Code, errcode.ERR2207.Message)
	}

	stateEntry := oauthStateEntry{
		ProviderName: provider.Name,
		CallbackURL:  callbackURL,
		ExpiresAt:    time.Now().Add(s.stateTTL(provider)),
		RustdeskId:   rustdeskId,
		Uuid:         uuid,
		DeviceOs:     deviceOs,
		DeviceType:   deviceType,
		DeviceName:   deviceName,
		PollToken:    pollToken,
	}
	stateEntry.CodeVerifier = randomOAuthToken(32)
	state := randomOAuthToken(24)
	if state == "" {
		return "", "", true, errcode.New(errcode.ERR2005.Code, errcode.ERR2005.Message)
	}

	if err = s.setState(state, stateEntry); err != nil {
		return "", "", true, err
	}

	query := url.Values{}
	if provider.Type == "wechat" {
		query.Set("appid", provider.ClientID)
	} else {
		query.Set("client_id", provider.ClientID)
	}
	query.Set("response_type", "code")
	scopeSeparator := " "
	if provider.Type == "qq" || provider.Type == "wechat" {
		scopeSeparator = ","
	}
	query.Set("scope", strings.Join(s.scopes(provider), scopeSeparator))
	query.Set("redirect_uri", callbackURL)
	query.Set("state", state)
	if provider.Type != "qq" && provider.Type != "wechat" && provider.Type != "apple" {
		challenge := sha256.Sum256([]byte(stateEntry.CodeVerifier))
		query.Set("code_challenge", base64.RawURLEncoding.EncodeToString(challenge[:]))
		query.Set("code_challenge_method", "S256")
	}
	if prompt := strings.TrimSpace(provider.Prompt); prompt != "" {
		query.Set("prompt", prompt)
	}

	return metadata.AuthorizationEndpoint + "?" + query.Encode(), pollToken, true, nil
}

// ConsumeClientCallback handles the OAuth provider redirect for a client login flow.
// It exchanges the authorization code, resolves the local user, issues a one-time
// client ticket, and stores it under the poll token so the client can retrieve it.
// Returns the poll token so the callback page can display it to the user.
func (s *OAuthProviderService) ConsumeClientCallback(providerName, code, state string) (string, error) {
	provider, ok := s.getProvider(providerName)
	if !ok {
		return "", errcode.New(errcode.ERR2001.Code, errcode.ERR2001.Message)
	}
	if !s.isProviderEnabled(provider) {
		return "", errcode.New(errcode.ERR2002.Code, errcode.ERR2002.Message)
	}
	if provider.AccountRole != "user" {
		return "", errcode.New(errcode.ERR2203.Code, errcode.ERR2203.Message)
	}
	if strings.TrimSpace(code) == "" || strings.TrimSpace(state) == "" {
		return "", errcode.New(errcode.ERR2003.Code, errcode.ERR2003.Message)
	}

	stored, ok := s.popState(state)
	if !ok || stored.ProviderName != provider.Name {
		return "", errcode.New(errcode.ERR2004.Code, errcode.ERR2004.Message)
	}
	if stored.PollToken == "" {
		return "", errcode.New(errcode.ERR2204.Code, errcode.ERR2204.Message)
	}

	tokenResp, err := s.exchangeCode(provider, code, stored.CallbackURL, stored.CodeVerifier)
	if err != nil {
		return "", err
	}

	claims, err := s.fetchUserClaims(provider, tokenResp)
	if err != nil {
		return "", err
	}

	user, err := s.resolveOAuthUser(provider, claims)
	if err != nil {
		return "", err
	}

	ticket := randomOAuthToken(24)
	if ticket == "" {
		return "", errcode.New(errcode.ERR2006.Code, errcode.ERR2006.Message)
	}

	ticketTTL := s.ticketTTL(provider)
	if err = s.setTicket(ticket, oauthTicketEntry{
		UserID:     user.Id,
		IsAdmin:    false,
		ExpiresAt:  time.Now().Add(ticketTTL),
		RustdeskId: stored.RustdeskId,
		Uuid:       stored.Uuid,
		DeviceOs:   stored.DeviceOs,
		DeviceType: stored.DeviceType,
		DeviceName: stored.DeviceName,
	}); err != nil {
		return "", err
	}

	if err = s.setPollEntry(stored.PollToken, ticket, time.Now().Add(ticketTTL)); err != nil {
		return "", err
	}

	return stored.PollToken, nil
}

// PollClientTicket peeks the one-time login ticket associated with the poll token.
// The poll entry is not consumed so the client can poll repeatedly until it is ready;
// the ticket itself is consumed atomically during ExchangeClientTicket.
// Returns the ticket and true when ready, or empty string and false when pending.
func (s *OAuthProviderService) PollClientTicket(pollToken string) (string, bool, error) {
	if strings.TrimSpace(pollToken) == "" {
		return "", false, errcode.New(errcode.ERR2205.Code, errcode.ERR2205.Message)
	}
	ticket, ok := s.peekPollEntry(pollToken)
	if !ok {
		return "", false, nil
	}
	return ticket, true, nil
}

// ExchangeClientTicket swaps a one-time client OAuth ticket for a long-lived client
// access token (90 days, is_admin=false, bound to rustdesk_id and uuid) and returns
// the resolved user so the caller can build a login response.
func (s *OAuthProviderService) ExchangeClientTicket(ticket string) (string, *model.User, error) {
	if strings.TrimSpace(ticket) == "" {
		return "", nil, errcode.New(errcode.ERR2007.Code, errcode.ERR2007.Message)
	}
	item, ok := s.popTicket(ticket)
	if !ok {
		return "", nil, errcode.New(errcode.ERR2008.Code, errcode.ERR2008.Message)
	}
	if item.IsAdmin {
		return "", nil, errcode.New(errcode.ERR2206.Code, errcode.ERR2206.Message)
	}
	var user model.User
	has, err := s.db.Where("id = ? and is_admin = 0 and status > 0", item.UserID).Get(&user)
	if err != nil {
		return "", nil, err
	}
	if !has {
		return "", nil, errcode.New(errcode.ERR2009.Code, errcode.ERR2009.Message)
	}
	token, err := s.issueClientOAuthToken(&user, item.RustdeskId, item.Uuid, item.DeviceOs, item.DeviceType, item.DeviceName)
	if err != nil {
		return "", nil, err
	}
	return token, &user, nil
}

func (s *OAuthProviderService) issueClientOAuthToken(user *model.User, rustdeskId, uuid, deviceOs, deviceType, deviceName string) (string, error) {
	_, _ = s.db.Where("user_id = ? and rustdesk_id = ? and status = 1 and is_admin = 0", user.Id, rustdeskId).Cols("status").Update(&model.AuthToken{Status: 0})
	signStr := fmt.Sprintf("%s_%s_%d_%s", rustdeskId, uuid, user.Id, time.Now().String())
	token := util.HmacSha256(signStr, s.cfg.SignKey)
	expired := 90 * 24 * time.Hour
	authToken := &model.AuthToken{
		UserId:     user.Id,
		RustdeskId: rustdeskId,
		Uuid:       uuid,
		DeviceOs:   deviceOs,
		DeviceType: deviceType,
		DeviceName: deviceName,
		TokenHash:  util.Sha256Hex(token),
		Expired:    time.Now().Add(expired),
		IsAdmin:    false,
		Status:     1,
	}
	_, err := s.db.Insert(authToken)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *OAuthProviderService) resolveClientCallbackURL(provider config.OAuthProviderConfig, requestBaseURL string) (string, error) {
	if explicit := strings.TrimSpace(provider.RedirectURL); explicit != "" {
		return explicit, nil
	}
	base := strings.TrimRight(strings.TrimSpace(requestBaseURL), "/")
	if base == "" {
		return "", errcode.New(errcode.ERR2025.Code, errcode.ERR2025.Message)
	}
	return fmt.Sprintf("%s/api/oauth/%s/callback", base, provider.Name), nil
}

func (s *OAuthProviderService) setPollEntry(pollToken, ticket string, expiresAt time.Time) error {
	if s.db != nil {
		_, err := s.db.Insert(&model.OAuthLoginSession{Kind: "poll", KeyHash: util.Sha256Hex(pollToken), Ticket: ticket, ExpiresAt: expiresAt, Status: 1})
		return err
	}
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	for k, v := range globalOAuthRuntimeStore.polls {
		if now.After(v.ExpiresAt) {
			delete(globalOAuthRuntimeStore.polls, k)
		}
	}
	globalOAuthRuntimeStore.polls[pollToken] = oauthPollEntry{Ticket: ticket, ExpiresAt: expiresAt}
	return nil
}

func (s *OAuthProviderService) peekPollEntry(pollToken string) (string, bool) {
	if s.db != nil {
		var session model.OAuthLoginSession
		has, err := s.db.Where("kind = ? and key_hash = ? and status = 1 and expires_at > ?", "poll", util.Sha256Hex(pollToken), time.Now()).Get(&session)
		if err != nil || !has {
			return "", false
		}
		return session.Ticket, true
	}
	now := time.Now()
	globalOAuthRuntimeStore.mu.Lock()
	defer globalOAuthRuntimeStore.mu.Unlock()
	v, ok := globalOAuthRuntimeStore.polls[pollToken]
	if !ok {
		return "", false
	}
	if now.After(v.ExpiresAt) {
		return "", false
	}
	return v.Ticket, true
}

// ConsumePollAndExchange atomically checks the poll entry and, if a ticket is
// ready, exchanges it for a client access token. The result is cached in the
// poll entry so that subsequent calls with the same pollToken return the same
// result without re-consuming the ticket. This makes the /api/oidc/auth-query
// endpoint idempotent and safe for client retries.
func (s *OAuthProviderService) ConsumePollAndExchange(pollToken string) (string, error) {
	if strings.TrimSpace(pollToken) == "" {
		return "", errcode.New(errcode.ERR2205.Code, errcode.ERR2205.Message)
	}
	keyHash := util.Sha256Hex(pollToken)
	if s.db != nil {
		var session model.OAuthLoginSession
		has, err := s.db.Where("kind = ? and key_hash = ? and status = 1 and expires_at > ?", "poll", keyHash, time.Now()).Get(&session)
		if err != nil {
			return "", err
		}
		if !has {
			return "", nil
		}
		if session.Result != "" {
			return session.Result, nil
		}
		ticket := strings.TrimSpace(session.Ticket)
		if ticket == "" {
			return "", nil
		}
		item, ok := s.popTicket(ticket)
		if !ok {
			return "", errcode.New(errcode.ERR2008.Code, errcode.ERR2008.Message)
		}
		if item.IsAdmin {
			return "", errcode.New(errcode.ERR2206.Code, errcode.ERR2206.Message)
		}
		var user model.User
		hasUser, userErr := s.db.Where("id = ? and is_admin = 0 and status > 0", item.UserID).Get(&user)
		if userErr != nil {
			return "", userErr
		}
		if !hasUser {
			return "", errcode.New(errcode.ERR2009.Code, errcode.ERR2009.Message)
		}
		token, tokenErr := s.issueClientOAuthToken(&user, item.RustdeskId, item.Uuid, item.DeviceOs, item.DeviceType, item.DeviceName)
		if tokenErr != nil {
			return "", tokenErr
		}
		var account model.OAuthAccount
		_, _ = s.db.Where("user_id = ? and is_admin = 0 and status = 1", user.Id).Get(&account)
		resultBytes, _ := json.Marshal(iris.Map{
			"access_token": token,
			"type":         "access_token",
			"user": iris.Map{
				"name":            user.Name,
				"display_name":    user.Name,
				"avatar":          account.Picture,
				"email":           user.Email,
				"note":            user.Note,
				"status":          user.Status,
				"info":            "",
				"is_admin":        false,
				"third_auth_type": account.Provider,
			},
		})
		resultStr := string(resultBytes)
		_, _ = s.db.Where("id = ?", session.Id).Cols("result").Update(&model.OAuthLoginSession{Result: resultStr})
		return resultStr, nil
	}
	v, ok := globalOAuthRuntimeStore.polls[pollToken]
	if !ok || time.Now().After(v.ExpiresAt) {
		return "", nil
	}
	if v.Result != "" {
		return v.Result, nil
	}
	ticket := strings.TrimSpace(v.Ticket)
	if ticket == "" {
		return "", nil
	}
	item, ok := s.popTicket(ticket)
	if !ok {
		return "", errcode.New(errcode.ERR2008.Code, errcode.ERR2008.Message)
	}
	if item.IsAdmin {
		return "", errcode.New(errcode.ERR2206.Code, errcode.ERR2206.Message)
	}
	var user model.User
	hasUser, userErr := s.db.Where("id = ? and is_admin = 0 and status > 0", item.UserID).Get(&user)
	if userErr != nil {
		return "", userErr
	}
	if !hasUser {
		return "", errcode.New(errcode.ERR2009.Code, errcode.ERR2009.Message)
	}
	token, tokenErr := s.issueClientOAuthToken(&user, item.RustdeskId, item.Uuid, item.DeviceOs, item.DeviceType, item.DeviceName)
	if tokenErr != nil {
		return "", tokenErr
	}
	var account model.OAuthAccount
	_, _ = s.db.Where("user_id = ? and is_admin = 0 and status = 1", user.Id).Get(&account)
	resultBytes, _ := json.Marshal(iris.Map{
		"access_token": token,
		"type":         "access_token",
		"user": iris.Map{
			"name":            user.Name,
			"display_name":    user.Name,
			"avatar":          account.Picture,
			"email":           user.Email,
			"note":            user.Note,
			"status":          user.Status,
			"info":            "",
			"is_admin":        false,
			"third_auth_type": account.Provider,
		},
	})
	resultStr := string(resultBytes)
	globalOAuthRuntimeStore.mu.Lock()
	globalOAuthRuntimeStore.polls[pollToken] = oauthPollEntry{Ticket: v.Ticket, ExpiresAt: v.ExpiresAt, Result: resultStr}
	globalOAuthRuntimeStore.mu.Unlock()
	return resultStr, nil
}
