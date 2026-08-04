package service

import (
	"net/url"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"testing"
)

func TestOAuthProviderService_ListClientProviders(t *testing.T) {
	cfg := &config.ServerConfig{
		OAuth: &config.OAuthConfig{
			Providers: []config.OAuthProviderConfig{
				{
					Type: "github", Name: "github-admin", DisplayName: "GitHub Admin",
					Enabled: true, ClientID: "c1", ClientSecret: "s1",
					AuthorizationEndpoint: "https://github.com/login/oauth/authorize",
					TokenEndpoint:         "https://github.com/login/oauth/access_token",
					AccountRole:           "admin",
				},
				{
					Type: "github", Name: "github-client", DisplayName: "GitHub Client",
					Enabled: true, ClientID: "c2", ClientSecret: "s2",
					AuthorizationEndpoint: "https://github.com/login/oauth/authorize",
					TokenEndpoint:         "https://github.com/login/oauth/access_token",
					AccountRole:           "user",
				},
				{
					Type: "github", Name: "github-disabled", DisplayName: "GitHub Disabled",
					Enabled: false, ClientID: "c3", ClientSecret: "s3",
					AuthorizationEndpoint: "https://github.com/login/oauth/authorize",
					TokenEndpoint:         "https://github.com/login/oauth/access_token",
					AccountRole:           "user",
				},
			},
		},
	}
	svc := NewOAuthProviderService(cfg, nil)
	providers := svc.ListClientProviders()
	if len(providers) != 1 {
		t.Fatalf("expected 1 client provider, got %d", len(providers))
	}
	if providers[0].Name != "github-client" {
		t.Fatalf("expected github-client, got %s", providers[0].Name)
	}
	if providers[0].AccountRole != "user" {
		t.Fatalf("expected accountRole=user, got %s", providers[0].AccountRole)
	}
}

func TestOAuthProviderService_ClientGitHubFlow(t *testing.T) {
	provider := newMockGitHubOAuthProvider(t)
	defer provider.Close()

	engine, err := db.NewEngine(&config.DbConfig{Driver: "sqlite", Dsn: ":memory:", TimeZone: "Asia/Shanghai", ShowSql: false})
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
					AccountRole:           "user",
					AutoCreateUser:        true,
					BindByEmail:           true,
				},
			},
		},
	}

	svc := NewOAuthProviderService(cfg, engine)

	authURL, pollToken, enabled, err := svc.BuildClientAuthURL("github", "http://localhost:12345", "rustdesk-id-1", "uuid-1", "Windows", "desktop", "MyPC")
	if err != nil {
		t.Fatalf("build client auth url: %v", err)
	}
	if !enabled {
		t.Fatalf("expected github provider enabled for client")
	}
	if pollToken == "" {
		t.Fatalf("poll token should not be empty")
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
		t.Fatalf("client github authorization must use PKCE S256")
	}
	if u.Query().Get("redirect_uri") != "http://localhost:12345/api/oauth/github/callback" {
		t.Fatalf("unexpected client callback url: %s", u.Query().Get("redirect_uri"))
	}

	consumedPollToken, err := svc.ConsumeClientCallback("github", "github-code", state)
	if err != nil {
		t.Fatalf("consume client callback: %v", err)
	}
	if consumedPollToken != pollToken {
		t.Fatalf("poll token mismatch: got %q want %q", consumedPollToken, pollToken)
	}

	if _, replayErr := svc.ConsumeClientCallback("github", "github-code", state); replayErr == nil {
		t.Fatalf("client oauth state replay must be rejected")
	}

	ticket, ready, err := svc.PollClientTicket(pollToken)
	if err != nil {
		t.Fatalf("poll client ticket: %v", err)
	}
	if !ready {
		t.Fatalf("ticket should be ready after callback")
	}
	if ticket == "" {
		t.Fatalf("ticket should not be empty")
	}

	ticket2, ready2, err := svc.PollClientTicket(pollToken)
	if err != nil || !ready2 || ticket2 != ticket {
		t.Fatalf("poll should be idempotent before exchange: ready=%v ticket=%q err=%v", ready2, ticket2, err)
	}

	token, user, err := svc.ExchangeClientTicket(ticket)
	if err != nil {
		t.Fatalf("exchange client ticket: %v", err)
	}
	if token == "" {
		t.Fatalf("client token should not be empty")
	}
	if user == nil {
		t.Fatalf("user should not be nil")
	}
	if user.IsAdmin {
		t.Fatalf("client oauth user must not be admin")
	}

	if _, _, replayErr := svc.ExchangeClientTicket(ticket); replayErr == nil {
		t.Fatalf("client ticket replay must be rejected")
	}

	var users []model.User
	if err = engine.Where("is_admin = 0").Find(&users); err != nil {
		t.Fatalf("query users: %v", err)
	}
	if len(users) != 1 {
		t.Fatalf("expected 1 normal user, got %d", len(users))
	}

	var tokens []model.AuthToken
	if err = engine.Where("user_id = ? and is_admin = 0 and rustdesk_id = ?", users[0].Id, "rustdesk-id-1").Find(&tokens); err != nil {
		t.Fatalf("query tokens: %v", err)
	}
	if len(tokens) != 1 {
		t.Fatalf("expected 1 client token bound to rustdesk-id-1, got %d", len(tokens))
	}
	if tokens[0].Uuid != "uuid-1" {
		t.Fatalf("token uuid mismatch: got %q want uuid-1", tokens[0].Uuid)
	}
}

func TestOAuthProviderService_ClientPollPending(t *testing.T) {
	engine, err := db.NewEngine(&config.DbConfig{Driver: "sqlite", Dsn: ":memory:", TimeZone: "Asia/Shanghai", ShowSql: false})
	if err != nil {
		t.Fatalf("new engine: %v", err)
	}
	if err = engine.Sync(new(model.OAuthLoginSession)); err != nil {
		t.Fatalf("sync: %v", err)
	}
	svc := NewOAuthProviderService(&config.ServerConfig{SignKey: "k"}, engine)

	ticket, ready, err := svc.PollClientTicket("non-existent-poll-token")
	if err != nil {
		t.Fatalf("poll should not error when pending: %v", err)
	}
	if ready {
		t.Fatalf("poll should not be ready for non-existent token")
	}
	if ticket != "" {
		t.Fatalf("ticket should be empty when pending")
	}
}

func TestOAuthProviderService_ClientRejectsAdminProvider(t *testing.T) {
	provider := newMockGitHubOAuthProvider(t)
	defer provider.Close()

	engine, err := db.NewEngine(&config.DbConfig{Driver: "sqlite", Dsn: ":memory:", TimeZone: "Asia/Shanghai", ShowSql: false})
	if err != nil {
		t.Fatalf("new engine: %v", err)
	}
	if err = engine.Sync(new(model.OAuthLoginSession)); err != nil {
		t.Fatalf("sync: %v", err)
	}

	cfg := &config.ServerConfig{
		SignKey: "test-sign-key",
		OAuth: &config.OAuthConfig{
			Providers: []config.OAuthProviderConfig{
				{
					Type:                  "github",
					Name:                  "github-admin-only",
					Enabled:               true,
					ClientID:              "c",
					ClientSecret:          "s",
					AuthorizationEndpoint: provider.URL + "/login/oauth/authorize",
					TokenEndpoint:         provider.URL + "/login/oauth/access_token",
					AccountRole:           "admin",
				},
			},
		},
	}
	svc := NewOAuthProviderService(cfg, engine)

	_, _, enabled, err := svc.BuildClientAuthURL("github-admin-only", "http://localhost:12345", "id", "uuid", "os", "type", "name")
	if err != nil {
		t.Fatalf("build should not error for admin provider: %v", err)
	}
	if enabled {
		t.Fatalf("admin provider should not be enabled for client login")
	}

	if _, err := svc.ConsumeClientCallback("github-admin-only", "code", "state"); err == nil {
		t.Fatalf("consume should reject admin provider for client login")
	}
}
