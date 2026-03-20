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
	if err = engine.Sync(new(model.User), new(model.AuthToken), new(model.OAuthAccount)); err != nil {
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

	token, err := svc.ExchangeAdminTicket(ticket)
	if err != nil {
		t.Fatalf("exchange ticket: %v", err)
	}
	if token == "" {
		t.Fatalf("token should not be empty")
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
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "github-access-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"id":         10001,
			"email":      "github-admin@example.com",
			"name":       "GitHub Admin",
			"avatar_url": "https://example.com/github-admin.png",
		})
	})

	return server
}
