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
	b.Handle("POST", "/auth/client-webauth/confirm", "PostClientWebauthConfirm")
	b.Handle("POST", "/auth/webauthn/login/begin", "PostWebauthnLoginBegin")
	b.Handle("POST", "/auth/webauthn/login/finish", "PostWebauthnLoginFinish")
	b.Handle("GET", "/auth/webauthn/enabled", "GetWebauthnEnabled")
	b.Handle("GET", "/auth/webauthn/auth-page", "GetWebauthnAuthPage")
	b.Handle("GET", "/auth/webauthn/register-page", "GetWebauthnRegisterPage")
}

type clientWebauthConfirmForm struct {
	admin.LoginForm
	PollToken string `json:"poll_token"`
}

// PostClientWebauthConfirm 使用现有后台登录页确认 RustDesk 客户端 WebAuth 登录。
func (c *AuthController) PostClientWebauthConfirm() mvc.Result {
	var form clientWebauthConfirmForm
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.dbError(err)
	}
	if strings.TrimSpace(form.PollToken) == "" {
		return c.Error(nil, errcode.New(errcode.ERR2205.Code, errcode.ERR2205.Message).Error())
	}
	if !captcha.VerifyCode(form.CaptchaId, form.Code) {
		c.recordAdminLoginAudit(0, form.Username, false, "client_webauth: CaptchaError")
		return c.Error(nil, errcode.New(errcode.ERR1001.Code, errcode.ERR1001.Message).Error())
	}

	var user model.User
	has, err := c.Db.Where("username = ? and status > 0", form.Username).Get(&user)
	if err != nil {
		return c.dbError(err)
	}
	if !has {
		return c.Error(nil, errcode.New(errcode.ERR1002.Code, errcode.ERR1002.Message).Error())
	}
	if !util.PasswordVerify(form.Password, user.Password) {
		return c.Error(nil, errcode.New(errcode.ERR1003.Code, errcode.ERR1003.Message).Error())
	}

	service := v2service.NewOAuthProviderService(c.Cfg, c.Db)
	if err = service.ConfirmWebauthLogin(form.PollToken, user.Id); err != nil {
		return c.Error(nil, err.Error())
	}
	c.recordAdminLoginAudit(user.Id, form.Username, true, "client_webauth: success")
	return c.Success(iris.Map{"ok": true}, "ok")
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
			body += "<!-- DEBUG: schemeURL=" + schemeURL + " pollToken=" + pollToken + " -->"
		} else {
			body = "<p class=\"ok\">已成功登录，请回到客户端继续。</p>"
			body += "<!-- DEBUG: pollToken is empty -->"
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
</div>
</body>
</html>`
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

func (c *AuthController) GetWebauthnEnabled() mvc.Result {
	service := v2service.NewWebauthnService(config.GetServerConfig(), c.Db)
	enabled := service.IsEnabled()
	if !enabled {
		rpID := strings.TrimSpace(c.Ctx.Host())
		if i := strings.IndexByte(rpID, ':'); i > 0 {
			rpID = rpID[:i]
		}
		if rpID != "" {
			scheme := "https"
			if c.Ctx.Request().TLS == nil {
				scheme = "http"
			}
			if forwardedProto := strings.TrimSpace(c.Ctx.GetHeader("X-Forwarded-Proto")); forwardedProto == "https" {
				scheme = "https"
			}
			origin := scheme + "://" + c.Ctx.Host()
			_ = service.UpdateConfig(rpID, []string{origin})
			enabled = service.IsEnabled()
		}
	}
	return c.Success(iris.Map{"enabled": enabled, "tlsPort": c.Cfg.HttpConfig.TLSPort}, "ok")
}

func (c *AuthController) PostWebauthnLoginBegin() mvc.Result {
	service := v2service.NewWebauthnService(config.GetServerConfig(), c.Db)
	if !service.IsEnabled() {
		rpID := strings.TrimSpace(c.Ctx.Host())
		if i := strings.IndexByte(rpID, ':'); i > 0 {
			rpID = rpID[:i]
		}
		if rpID != "" {
			scheme := "https"
			if c.Ctx.Request().TLS == nil {
				scheme = "http"
			}
			if forwardedProto := strings.TrimSpace(c.Ctx.GetHeader("X-Forwarded-Proto")); forwardedProto == "https" {
				scheme = "https"
			}
			origin := scheme + "://" + c.Ctx.Host()
			_ = service.UpdateConfig(rpID, []string{origin})
		}
	}
	if err := service.EnsureEnabled(); err != nil {
		return c.Error(nil, err.Error())
	}

	var form struct {
		Username string `json:"username"`
	}
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.Error(nil, errcode.New(errcode.ERR4001.Code, errcode.ERR4001.Message).Error())
	}
	form.Username = strings.TrimSpace(form.Username)
	if form.Username == "" {
		return c.Error(nil, errcode.New(errcode.ERR4001.Code, errcode.ERR4001.Message).Error())
	}

	options, _, err := service.BeginLogin(form.Username)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(options, "ok")
}

func (c *AuthController) PostWebauthnLoginFinish() mvc.Result {
	service := v2service.NewWebauthnService(config.GetServerConfig(), c.Db)
	if err := service.EnsureEnabled(); err != nil {
		return c.Error(nil, err.Error())
	}

	parsed, err := v2service.ParseCredentialRequestBody(c.Ctx.Request().Body)
	if err != nil {
		return c.Error(nil, errcode.New(errcode.ERR3106.Code, errcode.ERR3106.Message).Error())
	}

	user, err := service.FinishLoginParsed(parsed)
	if err != nil {
		c.recordAdminSecurityAudit("webauthn_login", false, err.Error())
		return c.Error(nil, err.Error())
	}

	_, _ = c.Db.Where("user_id = ? and status = 1 and is_admin = ?", user.Id, user.IsAdmin).Cols("status").Update(&model.AuthToken{Status: 0})

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
	if _, err = c.Db.Insert(authToken); err != nil {
		return c.dbError(err)
	}

	if user.IsAdmin {
		c.recordAdminLoginAudit(user.Id, user.Username, true, "webauthn")
	} else {
		c.recordUserLoginAudit(user.Id, user.Username, true, "webauthn")
	}
	return c.Success(iris.Map{
		"token":   token,
		"isAdmin": user.IsAdmin,
	}, "ok")
}

// GetWebauthnAuthPage 返回 Passkey 认证 HTML 页面（需通过 HTTPS 访问）
func (c *AuthController) GetWebauthnAuthPage() mvc.Result {
	username := c.Ctx.URLParam("username")
	redirect := c.Ctx.URLParam("redirect")

	html := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Passkey 认证</title>
<style>
* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif; display: flex; justify-content: center; align-items: center; min-height: 100vh; background: #f5f5f5; }
.card { background: #fff; border-radius: 12px; padding: 40px; max-width: 400px; width: 90%; box-shadow: 0 2px 12px rgba(0,0,0,0.08); text-align: center; }
.icon { font-size: 48px; margin-bottom: 16px; }
h2 { color: #333; margin-bottom: 8px; }
p { color: #666; margin-bottom: 24px; line-height: 1.5; }
#status { color: #1890ff; font-size: 14px; margin-top: 16px; min-height: 20px; }
.error { color: #ff4d4f !important; }
.success { color: #52c41a !important; }
</style>
</head>
<body>
<div class="card">
  <div class="icon">🔑</div>
  <h2>Passkey 认证</h2>
  <p id="desc">请在浏览器弹窗中完成 Passkey 认证</p>
  <div id="status">正在准备...</div>
</div>
<script>
const username = "` + username + `";
const redirect = "` + redirect + `";

function setStatus(msg, cls) {
  const el = document.getElementById("status");
  el.textContent = msg;
  el.className = cls || "";
}

function bufferToBase64url(buffer) {
  const bytes = new Uint8Array(buffer);
  let str = "";
  for (let i = 0; i < bytes.length; i++) str += String.fromCharCode(bytes[i]);
  return btoa(str).replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "");
}

function base64urlToBuffer(b64url) {
  const b64 = b64url.replace(/-/g, "+").replace(/_/g, "/");
  const pad = b64.length % 4 === 0 ? "" : "=".repeat(4 - b64.length % 4);
  const binary = atob(b64 + pad);
  const bytes = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) bytes[i] = binary.charCodeAt(i);
  return bytes.buffer;
}

function prepareAssertionOptions(options) {
  const prepared = {
    challenge: base64urlToBuffer(options.challenge),
    rpId: options.rpId,
    timeout: options.timeout,
    userVerification: options.userVerification
  };
  if (options.allowCredentials) {
    prepared.allowCredentials = options.allowCredentials.map(c => ({
      id: base64urlToBuffer(c.id),
      type: c.type,
      transports: c.transports
    }));
  }
  return prepared;
}

function serializeAssertion(cred) {
  return {
    id: cred.id,
    rawId: bufferToBase64url(cred.rawId),
    type: cred.type,
    response: {
      authenticatorData: bufferToBase64url(cred.response.authenticatorData),
      clientDataJSON: bufferToBase64url(cred.response.clientDataJSON),
      signature: bufferToBase64url(cred.response.signature),
      userHandle: cred.response.userHandle ? bufferToBase64url(cred.response.userHandle) : null
    }
  };
}

async function doLogin() {
  try {
    if (!window.PublicKeyCredential) {
      setStatus("此浏览器不支持 Passkey", "error");
      return;
    }

    setStatus("正在请求认证选项...");
    const beginResp = await fetch("/admin/auth/webauthn/login/begin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ username: username })
    });
    const beginData = await beginResp.json();
    if (beginData.code !== 200 || !beginData.data?.publicKey) {
      setStatus(beginData.message || "请求认证选项失败", "error");
      return;
    }

    setStatus("请完成 Passkey 认证...");
    const publicKey = prepareAssertionOptions(beginData.data.publicKey);
    const credential = await navigator.credentials.get({ publicKey });
    if (!credential) {
      setStatus("认证已取消", "error");
      return;
    }

    setStatus("正在验证...");
    const serialized = serializeAssertion(credential);
    const finishResp = await fetch("/admin/auth/webauthn/login/finish", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(serialized)
    });
    const finishData = await finishResp.json();
    if (finishData.code !== 200 || !finishData.data?.token) {
      setStatus(finishData.message || "认证失败", "error");
      return;
    }

    setStatus("认证成功！", "success");

    if (window.opener) {
      window.opener.postMessage({
        type: "webauthn-login",
        token: finishData.data.token,
        isAdmin: finishData.data.isAdmin
      }, "*");
      setTimeout(() => window.close(), 500);
    } else if (redirect) {
      const sep = redirect.includes("?") ? "&" : "?";
      window.location.href = redirect + sep + "token=" + encodeURIComponent(finishData.data.token) + "&isAdmin=" + finishData.data.isAdmin;
    } else {
      document.getElementById("desc").textContent = "认证成功，可以关闭此窗口";
    }
  } catch (e) {
    setStatus("认证过程出错: " + (e.message || e), "error");
  }
}

doLogin();
</script>
</body>
</html>`

	c.Ctx.ContentType("text/html; charset=utf-8")
	c.Ctx.WriteString(html)
	return nil
}

// GetWebauthnRegisterPage 返回 Passkey 注册 HTML 页面（需通过 HTTPS 访问）
// 通过 postMessage 从 opener 获取认证 token，完成注册后通知 opener
func (c *AuthController) GetWebauthnRegisterPage() mvc.Result {
	html := `<!DOCTYPE html>
<html lang="zh-CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Passkey 注册</title>
<style>
* { margin: 0; padding: 0; box-sizing: border-box; }
body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif; display: flex; justify-content: center; align-items: center; min-height: 100vh; background: #f5f5f5; }
.card { background: #fff; border-radius: 12px; padding: 40px; max-width: 400px; width: 90%; box-shadow: 0 2px 12px rgba(0,0,0,0.08); text-align: center; }
.icon { font-size: 48px; margin-bottom: 16px; }
h2 { color: #333; margin-bottom: 8px; }
p { color: #666; margin-bottom: 24px; line-height: 1.5; }
#status { color: #1890ff; font-size: 14px; margin-top: 16px; min-height: 20px; }
.error { color: #ff4d4f !important; }
.success { color: #52c41a !important; }
</style>
</head>
<body>
<div class="card">
  <div class="icon">🔑</div>
  <h2>Passkey 注册</h2>
  <p id="desc">请在浏览器弹窗中完成 Passkey 注册</p>
  <div id="status">等待认证信息...</div>
</div>
<script>
let authToken = "";
let credName = "";

function setStatus(msg, cls) {
  const el = document.getElementById("status");
  el.textContent = msg;
  el.className = cls || "";
}

function bufferToBase64url(buffer) {
  const bytes = new Uint8Array(buffer);
  let str = "";
  for (let i = 0; i < bytes.length; i++) str += String.fromCharCode(bytes[i]);
  return btoa(str).replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "");
}

function base64urlToBuffer(b64url) {
  const b64 = b64url.replace(/-/g, "+").replace(/_/g, "/");
  const pad = b64.length % 4 === 0 ? "" : "=".repeat(4 - b64.length % 4);
  const binary = atob(b64 + pad);
  const bytes = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) bytes[i] = binary.charCodeAt(i);
  return bytes.buffer;
}

function prepareCreationOptions(options) {
  const prepared = {
    challenge: base64urlToBuffer(options.challenge),
    rp: { name: options.rp.name, id: options.rp.id },
    user: {
      id: base64urlToBuffer(options.user.id),
      name: options.user.name,
      displayName: options.user.displayName
    },
    pubKeyCredParams: options.pubKeyCredParams,
    timeout: options.timeout,
    excludeCredentials: (options.excludeCredentials || []).map(c => ({
      id: base64urlToBuffer(c.id),
      type: c.type,
      transports: c.transports
    })),
    authenticatorSelection: options.authenticatorSelection,
    attestation: options.attestation || "none"
  };
  return prepared;
}

function serializeCreation(cred) {
  return {
    id: cred.id,
    rawId: bufferToBase64url(cred.rawId),
    type: cred.type,
    response: {
      attestationObject: bufferToBase64url(cred.response.attestationObject),
      clientDataJSON: bufferToBase64url(cred.response.clientDataJSON)
    },
    getClientExtensionResults: {}
  };
}

window.addEventListener("message", async function(event) {
  if (event.data?.type === "webauthn-register-data") {
    authToken = event.data.token || "";
    credName = event.data.name || "";
    window.removeEventListener("message", arguments.callee);
    await doRegister();
  }
});

if (window.opener) {
  window.opener.postMessage({ type: "webauthn-register-ready" }, "*");
} else {
  setStatus("请在管理后台中打开此页面", "error");
}

async function doRegister() {
  try {
    if (!window.PublicKeyCredential) {
      setStatus("此浏览器不支持 Passkey", "error");
      notifyOpener(false, "此浏览器不支持 Passkey");
      return;
    }
    if (!authToken) {
      setStatus("未收到认证信息", "error");
      notifyOpener(false, "未收到认证信息");
      return;
    }

    setStatus("正在请求注册选项...");
    const beginResp = await fetch("/admin/webauthn/register/begin", {
      method: "POST",
      headers: { "Content-Type": "application/json", "Authorization": "Bearer " + authToken }
    });
    const beginData = await beginResp.json();
    if (beginData.code !== 200 || !beginData.data?.publicKey) {
      setStatus(beginData.message || "请求注册选项失败", "error");
      notifyOpener(false, beginData.message || "请求注册选项失败");
      return;
    }

    setStatus("请完成 Passkey 注册...");
    const publicKey = prepareCreationOptions(beginData.data.publicKey);
    const credential = await navigator.credentials.create({ publicKey });
    if (!credential) {
      setStatus("注册已取消", "error");
      notifyOpener(false, "注册已取消");
      return;
    }

    setStatus("正在保存...");
    const serialized = serializeCreation(credential);
    const name = credName || new Date().toLocaleString();
    const finishResp = await fetch("/admin/webauthn/register/finish", {
      method: "POST",
      headers: { "Content-Type": "application/json", "Authorization": "Bearer " + authToken },
      body: JSON.stringify({ ...serialized, name: name })
    });
    const finishData = await finishResp.json();
    if (finishData.code !== 200) {
      setStatus(finishData.message || "注册失败", "error");
      notifyOpener(false, finishData.message || "注册失败");
      return;
    }

    setStatus("注册成功！", "success");
    notifyOpener(true, "");
    setTimeout(() => window.close(), 1000);
  } catch (e) {
    const msg = "注册过程出错: " + (e.message || e);
    setStatus(msg, "error");
    notifyOpener(false, msg);
  }
}

function notifyOpener(success, message) {
  if (window.opener) {
    window.opener.postMessage({
      type: "webauthn-register-result",
      success: success,
      message: message
    }, "*");
  }
}
</script>
</body>
</html>`

	c.Ctx.ContentType("text/html; charset=utf-8")
	c.Ctx.WriteString(html)
	return nil
}
