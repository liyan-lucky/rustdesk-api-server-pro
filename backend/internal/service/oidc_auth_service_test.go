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

func TestOIDCAuthService_AutoCreateAdminAndTicketFlow(t *testing.T) {
	provider := newMockOIDCProvider(t)
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
		OIDC: &config.OIDCConfig{
			Enabled:         true,
			ProviderName:    "mock",
			Issuer:          provider.URL,
			ClientID:        "cid",
			ClientSecret:    "csecret",
			RedirectURL:     "http://localhost/admin/auth/oidc/callback",
			Scopes:          []string{"openid", "profile", "email"},
			BindByEmail:     true,
			AutoCreateAdmin: true,
			SuccessRedirect: "/login",
			FailureRedirect: "/login",
			SubjectClaim:    "sub",
			EmailClaim:      "email",
			NameClaim:       "name",
			PictureClaim:    "picture",
		},
	}
	svc := NewOIDCAuthService(cfg, engine)

	authURL, enabled, err := svc.BuildAdminAuthURL("http://localhost:12345", "/login?redirect=%2F")
	if err != nil {
		t.Fatalf("build auth url: %v", err)
	}
	if !enabled {
		t.Fatalf("expected oidc enabled")
	}

	u, err := url.Parse(authURL)
	if err != nil {
		t.Fatalf("parse auth url: %v", err)
	}
	state := u.Query().Get("state")
	if state == "" {
		t.Fatalf("state should not be empty")
	}

	ticket, redirectTo, err := svc.ConsumeAdminCallback("mock-code", state)
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
	if err = engine.Where("provider = ? and subject = ?", "mock", "sub-001").Find(&accounts); err != nil {
		t.Fatalf("query oauth account: %v", err)
	}
	if len(accounts) != 1 {
		t.Fatalf("expected 1 oauth account, got %d", len(accounts))
	}

	var tokens []model.AuthToken
	if err = engine.Where("is_admin = 1 and status = 1").Find(&tokens); err != nil {
		t.Fatalf("query auth token: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 active admin token, got %d", len(tokens))
	}
}

func TestOIDCAuthService_BindByEmail(t *testing.T) {
	provider := newMockOIDCProvider(t)
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
	admin := &model.User{
		Username: "admin",
		Password: "x",
		Name:     "Admin",
		Email:    "admin@example.com",
		Status:   1,
		IsAdmin:  true,
	}
	if _, err = engine.Insert(admin); err != nil {
		t.Fatalf("insert admin: %v", err)
	}

	cfg := &config.ServerConfig{
		SignKey: "test-sign-key",
		OIDC: &config.OIDCConfig{
			Enabled:         true,
			ProviderName:    "mock",
			Issuer:          provider.URL,
			ClientID:        "cid",
			ClientSecret:    "csecret",
			RedirectURL:     "http://localhost/admin/auth/oidc/callback",
			BindByEmail:     true,
			AutoCreateAdmin: false,
			Scopes:          []string{"openid", "profile", "email"},
		},
	}
	svc := NewOIDCAuthService(cfg, engine)

	authURL, _, err := svc.BuildAdminAuthURL("http://localhost:12345", "/login")
	if err != nil {
		t.Fatalf("build auth url: %v", err)
	}
	u, _ := url.Parse(authURL)
	state := u.Query().Get("state")

	ticket, _, err := svc.ConsumeAdminCallback("mock-code", state)
	if err != nil {
		t.Fatalf("consume callback: %v", err)
	}
	if ticket == "" {
		t.Fatalf("expected non-empty ticket")
	}
}

func newMockOIDCProvider(t *testing.T) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	mux.HandleFunc("/.well-known/openid-configuration", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"authorization_endpoint": server.URL + "/authorize",
			"token_endpoint":         server.URL + "/token",
			"userinfo_endpoint":      server.URL + "/userinfo",
		})
	})

	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.Form.Get("code") == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": "mock-access-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	})

	mux.HandleFunc("/userinfo", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{
			"sub":     "sub-001",
			"email":   "admin@example.com",
			"name":    "OIDC Admin",
			"picture": "https://example.com/avatar.png",
		})
	})

	mux.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	return server
}
