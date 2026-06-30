package admin

import (
	"net/url"
	"rustdesk-api-server-pro/app/form/admin"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/helper/captcha"
	"rustdesk-api-server-pro/internal/core"
	v2service "rustdesk-api-server-pro/internal/service"
	"rustdesk-api-server-pro/util"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AuthController struct {
	basicController
	Cfg *config.ServerConfig
}

func (c *AuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/auth/oauth/providers", "GetAuthOauthProviders")
	b.Handle("GET", "/auth/oauth/url", "GetAuthOauthUrl")
	b.Handle("GET", "/auth/oauth/token", "GetAuthOauthToken")
	b.Handle("GET", "/auth/oauth/{provider:string}/callback", "HandleOauthCallback")
}

func (c *AuthController) PostAuthLogin() mvc.Result {
	var loginForm admin.LoginForm
	err := c.Ctx.ReadJSON(&loginForm)
	if err != nil {
		c.recordAdminLoginAudit(0, "", false, "decode_error: "+err.Error())
		return c.Error(nil, err.Error())
	}

	if !captcha.VerifyCode(loginForm.CaptchaId, loginForm.Code) {
		c.recordAdminLoginAudit(0, loginForm.Username, false, "CaptchaError")
		return c.Error(nil, "CaptchaError")
	}

	var user model.User
	get, err := c.Db.Where("username = ? and is_admin = 1", loginForm.Username).Get(&user)
	if err != nil {
		c.recordAdminLoginAudit(0, loginForm.Username, false, err.Error())
		return c.Error(nil, err.Error())
	}

	if !get {
		c.recordAdminLoginAudit(0, loginForm.Username, false, "UserNotExists")
		return c.Error(nil, "UserNotExists")
	}

	if !util.PasswordVerify(loginForm.Password, user.Password) {
		c.recordAdminLoginAudit(user.Id, loginForm.Username, false, "UsernameOrPasswordError")
		return c.Error(nil, "UsernameOrPasswordError")
	}

	// make other tokens expired
	_, _ = c.Db.Where("user_id = ? and status = 1 and is_admin = 1", user.Id).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})

	signStr := strconv.Itoa(user.Id) + user.Username + time.Now().String()
	token := util.HmacSha256(signStr, c.Cfg.SignKey)
	expired := 2 * time.Hour // 2 hours

	authToken := &model.AuthToken{
		UserId:    user.Id,
		TokenHash: util.Sha256Hex(token),
		Expired:   time.Now().Add(expired),
		IsAdmin:   true,
		Status:    1,
	}

	_, err = c.Db.Insert(authToken)
	if err != nil {
		c.recordAdminLoginAudit(user.Id, loginForm.Username, false, err.Error())
		return c.Error(nil, err.Error())
	}

	c.recordAdminLoginAudit(user.Id, loginForm.Username, true, "token")
	return c.Success(iris.Map{
		"token": token,
	}, "ok")
}

func (c *AuthController) GetAuthCaptcha() mvc.Result {
	id, img := captcha.CreateCaptcha()
	return c.Success(iris.Map{
		"id":  id,
		"img": img,
	}, "ok")
}

func (c *AuthController) GetAuthOidcUrl() mvc.Result {
	service := v2service.NewOIDCAuthService(c.Cfg, c.Db)
	redirect := c.Ctx.URLParamDefault("redirect", "")
	authURL, enabled, err := service.BuildAdminAuthURL(c.currentBaseURL(), redirect)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(iris.Map{
		"enabled": enabled,
		"url":     authURL,
	}, "ok")
}

func (c *AuthController) GetAuthOidcToken() mvc.Result {
	service := v2service.NewOIDCAuthService(c.Cfg, c.Db)
	ticket := c.Ctx.URLParamDefault("ticket", "")
	token, err := service.ExchangeAdminTicket(ticket)
	if err != nil {
		c.recordAdminSecurityAudit("admin_oidc_token_exchange", false, err.Error())
		return c.Error(nil, err.Error())
	}
	c.recordAdminSecurityAudit("admin_oidc_token_exchange", true, "token")
	return c.Success(iris.Map{
		"token": token,
	}, "ok")
}

func (c *AuthController) GetAuthOidcCallback() mvc.Result {
	service := v2service.NewOIDCAuthService(c.Cfg, c.Db)
	code := c.Ctx.URLParamDefault("code", "")
	state := c.Ctx.URLParamDefault("state", "")

	ticket, redirectTo, err := service.ConsumeAdminCallback(code, state)
	if err != nil {
		c.recordAdminSecurityAudit("admin_oidc_callback", false, err.Error())
		c.Ctx.Redirect(withQuery(redirectTo, "oidc_error", err.Error()), iris.StatusFound)
		return mvc.Response{}
	}

	c.recordAdminSecurityAudit("admin_oidc_callback", true, "ticket")
	target := withQuery(redirectTo, "oidc_ticket", ticket)
	c.Ctx.Redirect(target, iris.StatusFound)
	return mvc.Response{}
}

func (c *AuthController) GetAuthOauthProviders() mvc.Result {
	service := v2service.NewOAuthProviderService(c.Cfg, c.Db)
	return c.Success(service.ListEnabledProviders(), "ok")
}

func (c *AuthController) GetAuthOauthUrl() mvc.Result {
	service := v2service.NewOAuthProviderService(c.Cfg, c.Db)
	provider := c.Ctx.URLParamDefault("provider", "")
	redirect := c.Ctx.URLParamDefault("redirect", "")
	authURL, enabled, err := service.BuildAdminAuthURL(provider, c.currentBaseURL(), redirect)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(iris.Map{
		"enabled": enabled,
		"url":     authURL,
	}, "ok")
}

func (c *AuthController) GetAuthOauthToken() mvc.Result {
	service := v2service.NewOAuthProviderService(c.Cfg, c.Db)
	ticket := c.Ctx.URLParamDefault("ticket", "")
	token, err := service.ExchangeAdminTicket(ticket)
	if err != nil {
		c.recordAdminSecurityAudit("admin_oauth_token_exchange", false, err.Error())
		return c.Error(nil, err.Error())
	}
	c.recordAdminSecurityAudit("admin_oauth_token_exchange", true, "token")
	return c.Success(iris.Map{
		"token": token,
	}, "ok")
}

func (c *AuthController) HandleOauthCallback() mvc.Result {
	service := v2service.NewOAuthProviderService(c.Cfg, c.Db)
	provider := c.Ctx.Params().Get("provider")
	code := c.Ctx.URLParamDefault("code", "")
	state := c.Ctx.URLParamDefault("state", "")

	ticket, redirectTo, err := service.ConsumeAdminCallback(provider, code, state)
	if err != nil {
		c.recordAdminSecurityAudit("admin_oauth_callback", false, provider+": "+err.Error())
		c.Ctx.Redirect(withQuery(redirectTo, "oauth_error", err.Error()), iris.StatusFound)
		return mvc.Response{}
	}

	c.recordAdminSecurityAudit("admin_oauth_callback", true, provider+": ticket")
	target := withQuery(withQuery(redirectTo, "oauth_provider", provider), "oauth_ticket", ticket)
	c.Ctx.Redirect(target, iris.StatusFound)
	return mvc.Response{}
}

func (c *AuthController) currentBaseURL() string {
	scheme := strings.TrimSpace(c.Ctx.GetHeader("X-Forwarded-Proto"))
	if scheme == "" {
		if c.Ctx.Request().TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}
	host := strings.TrimSpace(c.Ctx.GetHeader("X-Forwarded-Host"))
	if host == "" {
		host = c.Ctx.Host()
	}
	return scheme + "://" + host
}

func (c *AuthController) recordAdminLoginAudit(userID int, username string, success bool, reason string) {
	_ = c.auditService().CreateSecurityAudit(core.SecurityAuditCreateCommand{
		UserID:    userID,
		Username:  username,
		Event:     "admin_login",
		IP:        c.Ctx.RemoteAddr(),
		UserAgent: c.Ctx.GetHeader("User-Agent"),
		Success:   success,
		Reason:    reason,
	})
}

func (c *AuthController) recordAdminSecurityAudit(event string, success bool, reason string) {
	_ = c.auditService().CreateSecurityAudit(core.SecurityAuditCreateCommand{
		Event:     event,
		IP:        c.Ctx.RemoteAddr(),
		UserAgent: c.Ctx.GetHeader("User-Agent"),
		Success:   success,
		Reason:    reason,
	})
}

func withQuery(target, key, value string) string {
	if target == "" {
		target = "/#/login"
	}
	u, err := url.Parse(target)
	if err != nil {
		return "/#/login"
	}
	if strings.HasPrefix(u.Fragment, "/") {
		fragmentURL, fragmentErr := url.Parse(u.Fragment)
		if fragmentErr == nil {
			q := fragmentURL.Query()
			q.Set(key, value)
			fragmentURL.RawQuery = q.Encode()
			u.Fragment = fragmentURL.String()
			return u.String()
		}
	}
	q := u.Query()
	q.Set(key, value)
	u.RawQuery = q.Encode()
	return u.String()
}
