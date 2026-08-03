package admin

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	v2service "rustdesk-api-server-pro/internal/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type OAuthController struct {
	basicController
}

func (c *OAuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/oauth/accounts", "HandleListAccounts")
	b.Handle("DELETE", "/oauth/account/{id:int}", "HandleDeleteAccount")
	b.Handle("GET", "/oauth/providers", "HandleListProviders")
	b.Handle("GET", "/oauth/providers/config", "HandleProviderConfigs")
	b.Handle("POST", "/oauth/provider", "HandleSaveProvider")
	b.Handle("DELETE", "/oauth/provider/{name:string}", "HandleDeleteProvider")
	b.Handle("POST", "/oauth/provider/{name:string}/test", "HandleTestProvider")
}

var oauthProviderNamePattern = regexp.MustCompile(`^[a-z0-9][a-z0-9_-]{0,49}$`)

type oauthProviderForm struct {
	OriginalName        string   `json:"originalName"`
	Type                string   `json:"type"`
	Name                string   `json:"name"`
	DisplayName         string   `json:"displayName"`
	Enabled             bool     `json:"enabled"`
	ClientID            string   `json:"clientId"`
	ClientSecret        string   `json:"clientSecret"`
	RedirectURL         string   `json:"redirectUrl"`
	Scopes              []string `json:"scopes"`
	AccountRole         string   `json:"accountRole"`
	BindByEmail         bool     `json:"bindByEmail"`
	AutoCreateAdmin     bool     `json:"autoCreateAdmin"`
	AutoCreateUser      bool     `json:"autoCreateUser"`
	AllowedEmailDomains []string `json:"allowedEmailDomains"`
}

func (c *OAuthController) HandleListAccounts() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)

	query := func() *xorm.Session {
		return c.Db.Table(&model.OAuthAccount{}).Desc("id")
	}

	pagination := db.NewPagination(currentPage, pageSize)
	accountList := make([]model.OAuthAccount, 0)
	if err := pagination.Paginate(query, &model.OAuthAccount{}, &accountList); err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0, len(accountList))
	for _, a := range accountList {
		list = append(list, iris.Map{
			"id":            a.Id,
			"user_id":       a.UserId,
			"provider":      a.Provider,
			"subject":       a.Subject,
			"email":         a.Email,
			"name":          a.Name,
			"is_admin":      a.IsAdmin,
			"status":        a.Status,
			"last_login_at": a.LastLoginAt.Format(config.TimeFormat),
			"created_at":    a.CreatedAt.Format(config.TimeFormat),
		})
	}

	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *OAuthController) HandleDeleteAccount() mvc.Result {
	id := c.Ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		return c.Error(nil, "InvalidAccountId")
	}

	_, err := c.Db.ID(id).Delete(&model.OAuthAccount{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(nil, "ok")
}

func (c *OAuthController) HandleListProviders() mvc.Result {
	cfg := config.GetServerConfig()
	providers := cfg.OAuthProviders()

	list := make([]iris.Map, 0, len(providers))
	for _, p := range providers {
		list = append(list, iris.Map{
			"type":        p.Type,
			"name":        p.Name,
			"displayName": p.DisplayName,
			"enabled":     p.Enabled,
			"accountRole": p.AccountRole,
		})
	}

	return c.Success(list, "ok")
}

func (c *OAuthController) HandleProviderConfigs() mvc.Result {
	cfg := config.GetServerConfig()
	list := make([]iris.Map, 0)
	if cfg.OAuth == nil {
		return c.Success(list, "ok")
	}
	for _, p := range cfg.OAuth.Providers {
		list = append(list, iris.Map{
			"type": p.Type, "name": p.Name, "displayName": p.DisplayName, "enabled": p.Enabled,
			"clientId": p.ClientID, "secretConfigured": strings.TrimSpace(p.ClientSecret) != "",
			"redirectUrl": p.RedirectURL, "scopes": p.Scopes, "accountRole": p.AccountRole,
			"bindByEmail": p.BindByEmail, "autoCreateAdmin": p.AutoCreateAdmin,
			"autoCreateUser": p.AutoCreateUser, "allowedEmailDomains": p.AllowedEmailDomains,
		})
	}
	return c.Success(list, "ok")
}

func (c *OAuthController) HandleSaveProvider() mvc.Result {
	var form oauthProviderForm
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.Error(nil, err.Error())
	}
	form.Type = strings.ToLower(strings.TrimSpace(form.Type))
	form.Name = strings.ToLower(strings.TrimSpace(form.Name))
	form.OriginalName = strings.ToLower(strings.TrimSpace(form.OriginalName))
	form.DisplayName = strings.TrimSpace(form.DisplayName)
	form.AccountRole = strings.ToLower(strings.TrimSpace(form.AccountRole))
	if form.Type != "github" {
		return c.Error(nil, "OnlyGitHubProviderIsCurrentlyEditable")
	}
	if !oauthProviderNamePattern.MatchString(form.Name) {
		return c.Error(nil, "InvalidProviderName")
	}
	if form.AccountRole != "user" {
		form.AccountRole = "admin"
	}
	if err := validateOAuthRedirectURL(form.RedirectURL); err != nil {
		return c.Error(nil, err.Error())
	}
	if form.Enabled && strings.TrimSpace(form.ClientID) == "" {
		return c.Error(nil, "ClientIdRequired")
	}

	cfg := config.GetServerConfig()
	if cfg.OAuth == nil {
		cfg.OAuth = &config.OAuthConfig{}
	}
	target := form.OriginalName
	if target == "" {
		target = form.Name
	}
	index := -1
	for i, provider := range cfg.OAuth.Providers {
		if strings.EqualFold(provider.Name, target) {
			index = i
		}
		if strings.EqualFold(provider.Name, form.Name) && !strings.EqualFold(provider.Name, target) {
			return c.Error(nil, "ProviderNameExists")
		}
	}
	secret := strings.TrimSpace(form.ClientSecret)
	if secret == "********" {
		secret = ""
	}
	if index >= 0 && secret == "" {
		secret = cfg.OAuth.Providers[index].ClientSecret
	}
	if form.Enabled && secret == "" {
		return c.Error(nil, "ClientSecretRequired")
	}
	provider := config.OAuthProviderConfig{}
	if index >= 0 {
		// Preserve lifecycle and redirect settings that are not exposed by this
		// first GitHub editor instead of resetting them on every save.
		provider = cfg.OAuth.Providers[index]
	}
	provider.Type = form.Type
	provider.Name = form.Name
	provider.DisplayName = form.DisplayName
	provider.Enabled = form.Enabled
	provider.ClientID = strings.TrimSpace(form.ClientID)
	provider.ClientSecret = secret
	provider.RedirectURL = strings.TrimSpace(form.RedirectURL)
	provider.Scopes = cleanOAuthValues(form.Scopes)
	provider.AccountRole = form.AccountRole
	provider.BindByEmail = form.BindByEmail
	provider.AutoCreateAdmin = form.AutoCreateAdmin
	provider.AutoCreateUser = form.AutoCreateUser
	provider.AllowedEmailDomains = cleanOAuthValues(form.AllowedEmailDomains)
	if provider.DisplayName == "" {
		provider.DisplayName = "GitHub"
	}
	if index >= 0 {
		cfg.OAuth.Providers[index] = provider
	} else {
		cfg.OAuth.Providers = append(cfg.OAuth.Providers, provider)
	}
	if err := config.SaveServerConfig(cfg); err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(iris.Map{"name": provider.Name, "secretConfigured": secret != ""}, "ok")
}

func (c *OAuthController) HandleDeleteProvider() mvc.Result {
	name := strings.ToLower(strings.TrimSpace(c.Ctx.Params().Get("name")))
	cfg := config.GetServerConfig()
	if cfg.OAuth == nil {
		return c.Success(nil, "ok")
	}
	providers := make([]config.OAuthProviderConfig, 0, len(cfg.OAuth.Providers))
	for _, provider := range cfg.OAuth.Providers {
		if !strings.EqualFold(provider.Name, name) {
			providers = append(providers, provider)
		}
	}
	cfg.OAuth.Providers = providers
	if err := config.SaveServerConfig(cfg); err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(nil, "ok")
}

func (c *OAuthController) HandleTestProvider() mvc.Result {
	name := strings.TrimSpace(c.Ctx.Params().Get("name"))
	cfg := config.GetServerConfig()
	var provider *config.OAuthProviderConfig
	if cfg.OAuth != nil {
		for i := range cfg.OAuth.Providers {
			if strings.EqualFold(cfg.OAuth.Providers[i].Name, name) {
				provider = &cfg.OAuth.Providers[i]
				break
			}
		}
	}
	if provider == nil {
		return c.Error(nil, "ProviderNotFound")
	}
	if !provider.Enabled {
		return c.Error(nil, "ProviderNotEnabled")
	}
	if strings.TrimSpace(provider.ClientID) == "" {
		return c.Error(nil, "ClientIdRequired")
	}
	if strings.TrimSpace(provider.ClientSecret) == "" {
		return c.Error(nil, "ClientSecretRequired")
	}
	if err := validateOAuthRedirectURL(provider.RedirectURL); err != nil {
		return c.Error(nil, err.Error())
	}

	svc := v2service.NewOAuthProviderService(cfg, c.Db)
	authURL, enabled, err := svc.BuildAdminAuthURL(name, c.Ctx.URLParamDefault("baseUrl", ""), "/#/login")
	if err != nil {
		return c.Error(nil, err.Error())
	}
	if !enabled || authURL == "" {
		return c.Error(nil, "ProviderNotEnabledOrIncomplete")
	}
	return c.Success(iris.Map{
		"authorizationUrl": authURL,
		"checks": iris.Map{
			"enabled": true, "clientIdConfigured": true, "clientSecretConfigured": true,
			"authorizationUrlGenerated": true, "credentialValidityVerified": false,
		},
	}, "ConfigurationCompleteCredentialValidityRequiresOAuthCallback")
}

func validateOAuthRedirectURL(raw string) error {
	value := strings.TrimSpace(raw)
	if value == "" {
		return nil
	}
	u, err := url.Parse(value)
	if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") {
		return errors.New("InvalidRedirectUrl")
	}
	return nil
}

func cleanOAuthValues(values []string) []string {
	cleaned := make([]string, 0, len(values))
	seen := map[string]struct{}{}
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		cleaned = append(cleaned, value)
	}
	return cleaned
}
