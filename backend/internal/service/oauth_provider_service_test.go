package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/util"
	"strings"
	"testing"
	"time"
)

func TestOAuthProviderService_ConfirmOAuthBindingRequiresTargetPassword(t *testing.T) {
	engine, err := db.NewEngine(&config.DbConfig{Driver: "sqlite", Dsn: ":memory:", TimeZone: "Asia/Shanghai", ShowSql: false})
	if err != nil {
		t.Fatalf("new engine: %v", err)
	}
	if err = engine.Sync(new(model.User), new(model.OAuthAccount), new(model.OAuthLoginSession)); err != nil {
		t.Fatalf("sync: %v", err)
	}
	passwordHash, err := util.Password("correct-password")
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}
	user := &model.User{Username: "target-user", Password: passwordHash, Status: 1, IsAdmin: false}
	if _, err = engine.Insert(user); err != nil {
		t.Fatalf("insert user: %v", err)
	}
	cfg := &config.ServerConfig{SignKey: "test-sign-key", OAuth: &config.OAuthConfig{Providers: []config.OAuthProviderConfig{{
		Type: "github", Name: "github", Enabled: true, AccountRole: "user", ClientID: "client-id", ClientSecret: "client-secret",
	}}}}
	svc := NewOAuthProviderService(cfg, engine)
	bindingTicket := "binding-ticket"
	if err = svc.setBindingTicket(bindingTicket, oauthBindingEntry{
		Provider: "github", Claims: OAuthUserClaims{Subject: "subject-1", Email: "third@example.com", Name: "Third User"},
	}, time.Now().Add(time.Minute)); err != nil {
		t.Fatalf("set binding ticket: %v", err)
	}
	if _, _, _, err = svc.ConfirmOAuthBinding(bindingTicket, user.Username, "wrong-password"); err == nil {
		t.Fatal("wrong target password must reject binding")
	}
	loginTicket, clientFlow, _, err := svc.ConfirmOAuthBinding(bindingTicket, user.Username, "correct-password")
	if err != nil || loginTicket == "" || clientFlow {
		t.Fatalf("confirm binding: ticket=%q client=%v err=%v", loginTicket, clientFlow, err)
	}
	var account model.OAuthAccount
	has, err := engine.Where("provider = ? and subject = ?", "github", "subject-1").Get(&account)
	if err != nil || !has || account.UserId != user.Id {
		t.Fatalf("binding account not persisted: has=%v account=%+v err=%v", has, account, err)
	}
	if _, _, _, replayErr := svc.ConfirmOAuthBinding(bindingTicket, user.Username, "correct-password"); replayErr == nil {
		t.Fatal("binding ticket replay must be rejected")
	}
}

func TestOAuthProviderServiceTokenResponsePreservesProviderError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"error":"bad_verification_code","error_description":"The code passed is incorrect or expired.\nretry"}`))
	}))
	defer server.Close()

	service := NewOAuthProviderService(&config.ServerConfig{}, nil)
	req, err := http.NewRequest(http.MethodPost, server.URL, nil)
	if err != nil {
		t.Fatalf("create token request: %v", err)
	}
	_, err = service.doTokenRequest(req)
	if err == nil {
		t.Fatal("expected provider token error")
	}
	message := err.Error()
	if !strings.HasPrefix(message, "ERR-2012:") || !strings.Contains(message, "provider_error=bad_verification_code") {
		t.Fatalf("provider error was not preserved: %s", message)
	}
	if strings.ContainsAny(message, "\r\n") {
		t.Fatalf("provider error contains a log-breaking newline: %q", message)
	}
}

func TestOAuthProviderService_ListEnabledProviders(t *testing.T) {
	cfg := &config.ServerConfig{
		OIDC: &config.OIDCConfig{
			Enabled:      true,
			ProviderName: "oidc",
			Issuer:       "https://sso.example.com",
			ClientID:     "legacy-client",
			ClientSecret: "legacy-secret",
		},
		OAuth: &config.OAuthConfig{
			Providers: []config.OAuthProviderConfig{
				{
					Type:         "google",
					Name:         "google",
					DisplayName:  "Google",
					Enabled:      true,
					ClientID:     "google-client",
					ClientSecret: "google-secret",
					Issuer:       "https://accounts.google.com",
				},
				{
					Type:         "github",
					Name:         "github",
					DisplayName:  "GitHub",
					Enabled:      false,
					ClientID:     "github-client",
					ClientSecret: "github-secret",
				},
			},
		},
	}

	svc := NewOAuthProviderService(cfg, nil)
	providers := svc.ListEnabledProviders()

	if len(providers) != 2 {
		t.Fatalf("expected 2 enabled providers, got %d", len(providers))
	}
	if providers[0].Name != "oidc" {
		t.Fatalf("expected first provider oidc, got %s", providers[0].Name)
	}
	if providers[1].Name != "google" {
		t.Fatalf("expected second provider google, got %s", providers[1].Name)
	}
}

func TestOAuthProviderService_GithubTicketFlow(t *testing.T) {
	provider := newMockGitHubOAuthProvider(t)
	defer provider.Close()

	engine, err := db.NewEngine(&config.DbConfig{
		Driver:   "sqlite",
		Dsn:      ":memory:",
		TimeZone: "Asia/Shanghai",
		ShowSql:  false,
	})
	if err != nil {
		t.Fatalf("new engine: %v", err)
	}
	if err = engine.Sync(new(model.User), new(model.AuthToken), new(model.OAuthAccount), new(model.OAuthLoginSession)); err != nil {
		t.Fatalf("sync: %v", err)
	}

	cfg := &config.ServerConfig{
		SignKey: "test-sign-key",
		OAuth: &config.OAuthConfig{
			Providers: []config.OAuthProviderConfig{
				{
					Type:                  "github",
					Name:                  "github",
					DisplayName:           "GitHub",
					Enabled:               true,
					ClientID:              "github-client-id",
					ClientSecret:          "github-client-secret",
					AuthorizationEndpoint: provider.URL + "/login/oauth/authorize",
					TokenEndpoint:         provider.URL + "/login/oauth/access_token",
					UserinfoEndpoint:      provider.URL + "/user",
					BindByEmail:           true,
					AutoCreateAdmin:       true,
					AutoCreateUser:        true,
					SuccessRedirect:       "/login",
					FailureRedirect:       "/login",
				},
			},
		},
	}

	svc := NewOAuthProviderService(cfg, engine)
	authURL, enabled, err := svc.BuildAdminAuthURL("github", "http://localhost:12345", "/login?redirect=%2F")
	if err != nil {
		t.Fatalf("build auth url: %v", err)
	}
	if !enabled {
		t.Fatalf("expected github provider enabled")
	}

	u, err := url.Parse(authURL)
	if err != nil {
		t.Fatalf("parse auth url: %v", err)
	}
	state := u.Query().Get("state")
	if state == "" {
		t.Fatalf("state should not be empty")
	}
	if u.Query().Get("code_challenge_method") != "S256" || u.Query().Get("code_challenge") == "" {
		t.Fatalf("github authorization must use PKCE S256")
	}

	_, bindingTicket, redirectTo, err := svc.ConsumeUnifiedCallback("github", "github-code", state)
	if err == nil || !strings.HasPrefix(err.Error(), "ERR-2023") || bindingTicket == "" {
		t.Fatalf("first callback must require explicit binding: ticket=%q err=%v", bindingTicket, err)
	}
	if redirectTo == "" {
		t.Fatalf("redirect should not be empty")
	}
	if _, _, _, replayErr := svc.ConsumeUnifiedCallback("github", "github-code", state); replayErr == nil {
		t.Fatalf("oauth state replay must be rejected")
	}
	ticket, clientFlow, _, err := svc.ConfirmOAuthBinding(bindingTicket, "", "")
	if err != nil || ticket == "" || clientFlow {
		t.Fatalf("explicit ordinary-user creation failed: ticket=%q client=%v err=%v", ticket, clientFlow, err)
	}

	token, _, err := svc.ExchangeAdminTicket(ticket)
	if err != nil {
		t.Fatalf("exchange ticket: %v", err)
	}
	if token == "" {
		t.Fatalf("token should not be empty")
	}
	if _, _, replayErr := svc.ExchangeAdminTicket(ticket); replayErr == nil {
		t.Fatalf("oauth ticket replay must be rejected")
	}

	var users []model.User
	if err = engine.Where("is_admin = 0").Find(&users); err != nil {
		t.Fatalf("query users: %v", err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 ordinary user, got %d", len(users))
	}

	var accounts []model.OAuthAccount
	if err = engine.Where("provider = ? and subject = ?", "github", "10001").Find(&accounts); err != nil {
		t.Fatalf("query oauth account: %v", err)
	}
	if len(accounts) != 1 {
		t.Fatalf("expected 1 oauth account, got %d", len(accounts))
	}
}

func TestOAuthProviderService_QQTicketFlow(t *testing.T) {
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/oauth2.0/token":
			if r.Method != http.MethodGet || r.URL.Query().Get("fmt") != "json" || r.URL.Query().Get("need_openid") != "1" {
				t.Fatalf("unexpected QQ token request: %s %s", r.Method, r.URL.RawQuery)
			}
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"access_token": "qq-token", "openid": "qq-openid", "expires_in": 3600})
		case "/user/get_user_info":
			if r.URL.Query().Get("access_token") != "qq-token" || r.URL.Query().Get("openid") != "qq-openid" || r.URL.Query().Get("oauth_consumer_key") != "qq-app-id" {
				t.Fatalf("unexpected QQ userinfo request: %s", r.URL.RawQuery)
			}
			_ = json.NewEncoder(w).Encode(map[string]interface{}{"ret": 0, "nickname": "QQ User", "figureurl_qq_2": "https://qlogo.example/avatar"})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	engine, err := db.NewEngine(&config.DbConfig{Driver: "sqlite", Dsn: ":memory:", TimeZone: "Asia/Shanghai", ShowSql: false})
	if err != nil {
		t.Fatalf("new engine: %v", err)
	}
	if err = engine.Sync(new(model.User), new(model.AuthToken), new(model.OAuthAccount), new(model.OAuthLoginSession)); err != nil {
		t.Fatalf("sync: %v", err)
	}
	cfg := &config.ServerConfig{SignKey: "test-sign-key", OAuth: &config.OAuthConfig{Providers: []config.OAuthProviderConfig{{
		Type: "qq", Name: "qq", DisplayName: "QQ", Enabled: true, ClientID: "qq-app-id", ClientSecret: "qq-app-key",
		AuthorizationEndpoint: server.URL + "/oauth2.0/authorize", TokenEndpoint: server.URL + "/oauth2.0/token", UserinfoEndpoint: server.URL + "/user/get_user_info",
		AccountRole: "user", AutoCreateUser: true, BindByEmail: true,
	}}}}
	svc := NewOAuthProviderService(cfg, engine)
	authURL, enabled, err := svc.BuildAdminAuthURL("qq", "http://localhost:12345", "/login")
	if err != nil || !enabled {
		t.Fatalf("build QQ auth url: enabled=%v err=%v", enabled, err)
	}
	u, _ := url.Parse(authURL)
	if u.Query().Get("scope") != "get_user_info" || u.Query().Get("code_challenge") != "" {
		t.Fatalf("unexpected QQ authorization query: %s", u.RawQuery)
	}
	_, bindingTicket, _, err := svc.ConsumeUnifiedCallback("qq", "qq-code", u.Query().Get("state"))
	if err == nil || bindingTicket == "" {
		t.Fatalf("QQ first callback must require binding: ticket=%q err=%v", bindingTicket, err)
	}
	if _, _, _, err = svc.ConfirmOAuthBinding(bindingTicket, "", ""); err != nil {
		t.Fatalf("explicit QQ user creation failed: %v", err)
	}
	var account model.OAuthAccount
	has, err := engine.Where("provider = ? and subject = ?", "qq", "qq-openid").Get(&account)
	if err != nil || !has || account.Name != "QQ User" {
		t.Fatalf("QQ account was not persisted correctly: has=%v account=%+v err=%v", has, account, err)
	}
}

func TestOAuthProviderService_GithubTicketFlowWithoutInMemoryState(t *testing.T) {
	provider := newMockGitHubOAuthProvider(t)
	defer provider.Close()

	engine, err := db.NewEngine(&config.DbConfig{
		Driver:   "sqlite",
		Dsn:      ":memory:",
		TimeZone: "Asia/Shanghai",
		ShowSql:  false,
	})
	if err != nil {
		t.Fatalf("new engine: %v", err)
	}
	if err = engine.Sync(new(model.User), new(model.AuthToken), new(model.OAuthAccount), new(model.OAuthLoginSession)); err != nil {
		t.Fatalf("sync: %v", err)
	}

	cfg := &config.ServerConfig{
		SignKey: "test-sign-key",
		OAuth: &config.OAuthConfig{
			Providers: []config.OAuthProviderConfig{
				{
					Type:                  "github",
					Name:                  "github",
					DisplayName:           "GitHub",
					Enabled:               true,
					ClientID:              "github-client-id",
					ClientSecret:          "github-client-secret",
					AuthorizationEndpoint: provider.URL + "/login/oauth/authorize",
					TokenEndpoint:         provider.URL + "/login/oauth/access_token",
					UserinfoEndpoint:      provider.URL + "/user",
					BindByEmail:           true,
					AccountRole:           "user",
					AutoCreateUser:        true,
					SuccessRedirect:       "/login",
					FailureRedirect:       "/login",
				},
			},
		},
	}

	svc := NewOAuthProviderService(cfg, engine)
	authURL, enabled, err := svc.BuildAdminAuthURL("github", "http://localhost:12345", "/login")
	if err != nil {
		t.Fatalf("build auth url: %v", err)
	}
	if !enabled {
		t.Fatalf("expected github provider enabled")
	}

	u, err := url.Parse(authURL)
	if err != nil {
		t.Fatalf("parse auth url: %v", err)
	}
	state := u.Query().Get("state")
	if state == "" {
		t.Fatalf("state should not be empty")
	}

	globalOAuthRuntimeStore.mu.Lock()
	globalOAuthRuntimeStore.states = map[string]oauthStateEntry{}
	globalOAuthRuntimeStore.mu.Unlock()

	_, bindingTicket, _, err := svc.ConsumeUnifiedCallback("github", "github-code", state)
	if err == nil || bindingTicket == "" {
		t.Fatalf("callback without in-memory state must require binding: ticket=%q err=%v", bindingTicket, err)
	}
	if _, _, _, err = svc.ConfirmOAuthBinding(bindingTicket, "", ""); err != nil {
		t.Fatalf("explicit user creation failed: %v", err)
	}
	var normalUsers []model.User
	if err = engine.Where("is_admin = 0").Find(&normalUsers); err != nil || len(normalUsers) != 1 {
		t.Fatalf("expected one normal oauth user, users=%d err=%v", len(normalUsers), err)
	}
}

func newMockGitHubOAuthProvider(t *testing.T) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/login/oauth/authorize", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("/login/oauth/access_token", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Form.Get("code") == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Form.Get("code_verifier") == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "github-access-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":         10001,
			"email":      nil,
			"login":      "github-admin",
			"name":       "GitHub Admin",
			"avatar_url": "https://example.com/github-admin.png",
		})
	})
	mux.HandleFunc("/user/emails", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode([]map[string]interface{}{{
			"email": "github-admin@example.com", "primary": true, "verified": true,
		}})
	})

	return server
}

func TestNormalizeOAuthRedirectTargetUsesHashRouter(t *testing.T) {
	tests := map[string]string{
		"":                    "/#/login",
		"/login":              "/#/login",
		"/system/oauth?tab=1": "/#/system/oauth?tab=1",
		"/#/system/oauth":     "/#/system/oauth",
		"https://evil.test/x": "/#/login",
	}
	for input, expected := range tests {
		if actual := normalizeOAuthRedirectTarget(input); actual != expected {
			t.Fatalf("normalizeOAuthRedirectTarget(%q) = %q, want %q", input, actual, expected)
		}
	}
}
