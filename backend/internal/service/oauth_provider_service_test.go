package service

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"testing"
)

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

	ticket, redirectTo, err := svc.ConsumeAdminCallback("github", "github-code", state)
	if err != nil {
		t.Fatalf("consume callback: %v", err)
	}
	if ticket == "" {
		t.Fatalf("ticket should not be empty")
	}
	if redirectTo == "" {
		t.Fatalf("redirect should not be empty")
	}
	if _, failureRedirect, replayErr := svc.ConsumeAdminCallback("github", "github-code", state); replayErr == nil {
		t.Fatalf("oauth state replay must be rejected")
	} else if failureRedirect != "/#/login" {
		t.Fatalf("failed OAuth callback must return to login, got %q", failureRedirect)
	}

	token, err := svc.ExchangeAdminTicket(ticket)
	if err != nil {
		t.Fatalf("exchange ticket: %v", err)
	}
	if token == "" {
		t.Fatalf("token should not be empty")
	}
	if _, replayErr := svc.ExchangeAdminTicket(ticket); replayErr == nil {
		t.Fatalf("oauth ticket replay must be rejected")
	}

	var users []model.User
	if err = engine.Where("is_admin = 1").Find(&users); err != nil {
		t.Fatalf("query users: %v", err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 admin user, got %d", len(users))
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
	ticket, _, err := svc.ConsumeAdminCallback("qq", "qq-code", u.Query().Get("state"))
	if err != nil || ticket == "" {
		t.Fatalf("consume QQ callback: ticket=%q err=%v", ticket, err)
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

	ticket, _, err := svc.ConsumeAdminCallback("github", "github-code", state)
	if err != nil {
		t.Fatalf("consume callback without in-memory state: %v", err)
	}
	if ticket == "" {
		t.Fatalf("ticket should not be empty")
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
