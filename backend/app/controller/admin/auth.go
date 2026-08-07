package admin

import (
	"net/url"
	"rustdesk-api-server-pro/app/form/admin"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/helper/captcha"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/errcode"
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
		return c.dbError(err)
	}

	if !captcha.VerifyCode(loginForm.CaptchaId, loginForm.Code) {
		c.recordAdminLoginAudit(0, loginForm.Username, false, "CaptchaError")
		return c.Error(nil, errcode.New(errcode.ERR1001.Code, errcode.ERR1001.Message).Error())
	}

	var user model.User
	get, err := c.Db.Where("username = ? and status > 0", loginForm.Username).Get(&user)
	if err != nil {
		c.recordAdminLoginAudit(0, loginForm.Username, false, err.Error())
		return c.dbError(err)
	}

	if !get {
		c.recordAdminLoginAudit(0, loginForm.Username, false, "UserNotExists")
		return c.Error(nil, errcode.New(errcode.ERR1002.Code, errcode.ERR1002.Message).Error())
	}

	if !util.PasswordVerify(loginForm.Password, user.Password) {
		c.recordAdminLoginAudit(user.Id, loginForm.Username, false, "UsernameOrPasswordError")
		return c.Error(nil, errcode.New(errcode.ERR1003.Code, errcode.ERR1003.Message).Error())
	}

	_, _ = c.Db.Where("user_id = ? and status = 1 and is_admin = ?", user.Id, user.IsAdmin).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})

	signStr := strconv.Itoa(user.Id) + user.Username + time.Now().String()
	token := util.HmacSha256(signStr, c.Cfg.SignKey)
	expired := 2 * time.Hour

	authToken := &model.AuthToken{
		UserId:    user.Id,
		TokenHash: util.Sha256Hex(token),
		Expired:   time.Now().Add(expired),
		IsAdmin:   user.IsAdmin,
		Status:    1,
	}

	_, err = c.Db.Insert(authToken)
	if err != nil {
		c.recordAdminLoginAudit(user.Id, loginForm.Username, false, err.Error())
		return c.dbError(err)
	}

	if user.IsAdmin {
		c.recordAdminLoginAudit(user.Id, loginForm.Username, true, "token")
	} else {
		c.recordUserLoginAudit(user.Id, loginForm.Username, true, "token")
	}
	return c.Success(iris.Map{
		"token":   token,
		"isAdmin": user.IsAdmin,
	}, "ok")
}

func (c *AuthController) recordUserLoginAudit(userID int, username string, success bool, reason string) {
	_ = c.auditService().CreateSecurityAudit(core.SecurityAuditCreateCommand{
		UserID:    userID,
		Username:  username,
		Event:     "web_user_login",
		IP:        c.Ctx.RemoteAddr(),
		UserAgent: c.Ctx.GetHeader("User-Agent"),
		Success:   success,
		Reason:    reason,
	})
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
		return c.dbError(err)
	}
	return c.Success(iris.Map{
		"enabled": enabled,
		"url":     authURL,
	}, "ok")
}

func (c *AuthController) GetAuthOidcToken() mvc.Result {
	service := v2service.NewOIDCAuthService(c.Cfg, c.Db)
	ticket := c.Ctx.URLParamDefault("ticket", "")
	token, isAdmin, err := service.ExchangeAdminTicket(ticket)
	if err != nil {
		c.recordAdminSecurityAudit("admin_oidc_token_exchange", false, err.Error())
		return c.dbError(err)
	}
	c.recordAdminSecurityAudit("admin_oidc_token_exchange", true, "token")
	return c.Success(iris.Map{
		"token":   token,
		"isAdmin": isAdmin,
	}, "ok")
}

func (c *AuthController) GetAuthOidcCallback() mvc.Result {
	service := v2service.NewOIDCAuthService(c.Cfg, c.Db)
	code := c.Ctx.URLParamDefault("code", "")
	state := c.Ctx.URLParamDefault("state", "")

	ticket, redirectTo, err := service.ConsumeAdminCallback(code, state)
	if err != nil {
		c.recordAdminSecurityAudit("admin_oidc_callback", false, err.Error())
		c.Ctx.Redirect(withQuery(redirectTo, "oidc_error", "auth_failed"), iris.StatusFound)
		return mvc.Response{}
	}

	c.recordAdminSecurityAudit("admin_oidc_callback", true, "ticket")
	target := withQuery(redirectTo, "oidc_ticket", ticket)
	c.Ctx.Redirect(target, iris.StatusFound)
	return mvc.Response{}
}

func (c *AuthController) GetAuthOauthProviders() mvc.Result {
	service := v2service.NewOAuthProviderService(config.GetServerConfig(), c.Db)
	return c.Success(service.ListEnabledProviders(), "ok")
}

func (c *AuthController) GetAuthOauthUrl() mvc.Result {
	service := v2service.NewOAuthProviderService(config.GetServerConfig(), c.Db)
	provider := c.Ctx.URLParamDefault("provider", "")
	redirect := c.Ctx.URLParamDefault("redirect", "")
	authURL, enabled, err := service.BuildAdminAuthURL(provider, c.currentBaseURL(), redirect)
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(iris.Map{
		"enabled": enabled,
		"url":     authURL,
	}, "ok")
}

func (c *AuthController) GetAuthOauthToken() mvc.Result {
	service := v2service.NewOAuthProviderService(config.GetServerConfig(), c.Db)
	ticket := c.Ctx.URLParamDefault("ticket", "")
	token, isAdmin, err := service.ExchangeAdminTicket(ticket)
	if err != nil {
		c.recordAdminSecurityAudit("admin_oauth_token_exchange", false, err.Error())
		return c.dbError(err)
	}
	c.recordAdminSecurityAudit("admin_oauth_token_exchange", true, "token")
	return c.Success(iris.Map{
		"token":   token,
		"isAdmin": isAdmin,
	}, "ok")
}

func (c *AuthController) HandleOauthCallback() mvc.Result {
	service := v2service.NewOAuthProviderService(config.GetServerConfig(), c.Db)
	provider := c.Ctx.Params().Get("provider")
	code := c.Ctx.URLParamDefault("code", "")
	state := c.Ctx.URLParamDefault("state", "")

	pollToken, ticket, redirectTo, err := service.ConsumeUnifiedCallback(provider, code, state)
	if err != nil {
		if pollToken != "" {
			c.recordAdminSecurityAudit("client_oauth_callback", false, provider+": "+err.Error())
			return c.renderOAuthCallbackPage(false, oauthCallbackErrorCode(err), pollToken)
		}
		c.recordAdminSecurityAudit("admin_oauth_callback", false, provider+": "+err.Error())
		c.Ctx.Redirect(withQuery(redirectTo, "oauth_error", oauthCallbackErrorCode(err)), iris.StatusFound)
		return mvc.Response{}
	}

	if pollToken != "" {
		c.recordAdminSecurityAudit("client_oauth_callback", true, provider+": poll_token")
		return c.renderOAuthCallbackPage(true, "", pollToken)
	}

	c.recordAdminSecurityAudit("admin_oauth_callback", true, provider+": ticket")
	target := withQuery(withQuery(redirectTo, "oauth_provider", provider), "oauth_ticket", ticket)
	c.Ctx.Redirect(target, iris.StatusFound)
	return mvc.Response{}
}

func (c *AuthController) renderOAuthCallbackPage(success bool, errorCode, pollToken string) mvc.Result {
	var title, body string
	var schemeURL string
	if success && pollToken != "" {
		schemeURL = "rustdesk://oauth/callback?poll_token=" + pollToken
	}
	if success {
		title = "第三方登录成功"
		if schemeURL != "" {
			body = "<p class=\"ok\">已成功登录！</p>"
			body += "<a href=\"" + schemeURL + "\" class=\"launch-btn\" id=\"launch-btn\">返回 RustDesk 客户端</a>"
			body += "<p class=\"tip\">请点击上方按钮返回客户端<br>如按钮无效，请手动切回客户端</p>"
		} else {
			body = "<p class=\"ok\">已成功登录，请回到客户端继续。</p>"
		}
	} else {
		title = "第三方登录失败"
		body = "<p class=\"err\">登录失败，请回到客户端重试。</p><p class=\"code\">错误码：" + errorCode + "</p>"
	}
	html := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>` + title + `</title>
<style>
body{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,"PingFang SC","Microsoft YaHei",sans-serif;display:flex;align-items:center;justify-content:center;min-height:100vh;margin:0;background:#f5f5f5}
.card{background:#fff;border-radius:12px;box-shadow:0 2px 12px rgba(0,0,0,.08);padding:40px 48px;max-width:420px;text-align:center}
h1{font-size:20px;margin:0 0 16px}
.ok{color:#16a34a;font-size:16px;margin-bottom:8px}
.err{color:#dc2626;font-size:16px}
.code{color:#6b7280;font-size:13px;margin-top:8px}
.launch-btn{display:inline-block;margin-top:16px;padding:14px 36px;background:#2563eb;color:#fff;border-radius:8px;text-decoration:none;font-size:18px;font-weight:600;transition:background .2s;box-shadow:0 2px 8px rgba(37,99,235,.3)}
.launch-btn:hover{background:#1d4ed8;transform:translateY(-1px)}
.tip{color:#9ca3af;font-size:13px;margin-top:16px;line-height:1.6}
</style>
</head>
<body>
<div class="card">
<h1>` + title + `</h1>
` + body + `
</div>`
	if schemeURL != "" {
		html += "\n<script>setTimeout(function(){var btn=document.getElementById('launch-btn');if(btn)btn.click();},500);</script>"
	}
	html += "\n</body>\n</html>"
	c.Ctx.ContentType("text/html; charset=utf-8")
	_, _ = c.Ctx.WriteString(html)
	return mvc.Response{}
}

func oauthCallbackErrorCode(err error) string {
	if err == nil {
		return errcode.ERR2212.Message
	}
	msg := err.Error()
	if strings.HasPrefix(msg, "ERR-") {
		code := strings.SplitN(msg, ":", 2)[0]
		switch code {
		case errcode.ERR2023.Code, errcode.ERR2022.Code:
			return errcode.ERR2208.Message
		case errcode.ERR2004.Code, errcode.ERR2029.Code:
			return errcode.ERR2210.Message
		case errcode.ERR2030.Code, errcode.ERR2031.Code, errcode.ERR2034.Code:
			return errcode.ERR2209.Message
		default:
			return errcode.ERR2212.Message
		}
	}
	switch {
	case strings.Contains(msg, "NoBindableOauthAccount"), strings.Contains(msg, "BoundAdminUserNotAvailable"):
		return errcode.ERR2208.Message
	case strings.Contains(msg, "timeout"), strings.Contains(msg, "deadline exceeded"), strings.Contains(msg, "connection refused"):
		return errcode.ERR2209.Message
	case strings.Contains(msg, "StateInvalidOrExpired"), strings.Contains(msg, "StateExpired"):
		return errcode.ERR2210.Message
	default:
		return errcode.ERR2212.Message
	}
}

func (c *AuthController) currentBaseURL() string {
	scheme := "http"
	if c.Ctx.Request().TLS != nil {
		scheme = "https"
	}
	if forwardedProto := strings.TrimSpace(c.Ctx.GetHeader("X-Forwarded-Proto")); forwardedProto == "https" {
		scheme = "https"
	}

	// Do not trust X-Forwarded-Host here. OAuth/OIDC callback URLs must not be
	// derived from attacker-controlled forwarding headers. Operators that need a
	// public reverse-proxy URL should set oidc.redirectUrl or oauth.providers[].redirectUrl.
	host := strings.TrimSpace(c.Ctx.Host())
	if host == "" {
		host = strings.TrimSpace(c.Ctx.Request().Host)
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
