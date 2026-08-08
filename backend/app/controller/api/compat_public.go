package api

import (
	"encoding/json"
	"strconv"
	"strings"

	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/internal/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/tidwall/gjson"
)

// CompatPublicController provides compatibility endpoints used by newer RustDesk clients.
// Unsupported features return stable payloads instead of 404.
type CompatPublicController struct {
	basicController
}

func (c *CompatPublicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "status", "HandleStatus")
	b.Handle("POST", "status", "HandleStatus")
	b.Handle("GET", "health", "HandleHealth")
	b.Handle("POST", "health", "HandleHealth")
	b.Handle("GET", "ping", "HandleHealth")
	b.Handle("POST", "ping", "HandleHealth")
	b.Handle("GET", "info", "HandleVersion")
	b.Handle("POST", "info", "HandleVersion")
	b.Handle("GET", "version", "HandleVersion")
	b.Handle("POST", "version", "HandleVersion")
	b.Handle("GET", "features", "HandleFeatures")
	b.Handle("POST", "features", "HandleFeatures")
	b.Handle("GET", "capabilities", "HandleFeatures")
	b.Handle("POST", "capabilities", "HandleFeatures")
	b.Handle("GET", "compat/features", "HandleFeatures")
	b.Handle("POST", "compat/features", "HandleFeatures")
	b.Handle("GET", "config", "HandleClientConfig")
	b.Handle("POST", "config", "HandleClientConfig")
	b.Handle("GET", "client-config", "HandleClientConfig")
	b.Handle("POST", "client-config", "HandleClientConfig")
	b.Handle("GET", "client_config", "HandleClientConfig")
	b.Handle("POST", "client_config", "HandleClientConfig")
	b.Handle("GET", "server-config", "HandleClientConfig")
	b.Handle("POST", "server-config", "HandleClientConfig")
	b.Handle("GET", "server_config", "HandleClientConfig")
	b.Handle("POST", "server_config", "HandleClientConfig")
	b.Handle("GET", "server/info", "HandleVersion")
	b.Handle("POST", "server/info", "HandleVersion")
	b.Handle("GET", "compat-target", "HandleCompatTarget")
	b.Handle("POST", "compat-target", "HandleCompatTarget")
	b.Handle("GET", "compat/target", "HandleCompatTarget")
	b.Handle("POST", "compat/target", "HandleCompatTarget")
	b.Handle("GET", "compat/version", "HandleCompatTarget")
	b.Handle("POST", "compat/version", "HandleCompatTarget")
	b.Handle("GET", "sysinfo_ver", "HandleSysinfoVer")
	b.Handle("POST", "sysinfo_ver", "HandleSysinfoVer")
	b.Handle("POST", "oidc/auth", "HandleOidcAuth")
	b.Handle("GET", "oidc/auth-query", "HandleOidcAuthQuery")
	b.Handle("POST", "oidc/auth-query", "HandleOidcAuthQuery")
	b.Handle("POST", "record", "HandleRecord")
	b.Handle("GET", "devices/deploy", "HandleDevicesDeploy")
	b.Handle("POST", "devices/deploy", "HandleDevicesDeploy")
}

func (c *CompatPublicController) HandleHealth() mvc.Result {
	c.recordCompatAPIAudit(false, 200, "ok", "", nil)
	return mvc.Response{Object: iris.Map{
		"ok":      true,
		"status":  "ok",
		"service": "rustdesk-api-server-pro",
	}}
}

func (c *CompatPublicController) HandleStatus() mvc.Result {
	c.recordCompatAPIAudit(false, 200, "ok", "", nil)
	return mvc.Response{Object: iris.Map{
		"ok":            true,
		"status":        "ok",
		"service":       "rustdesk-api-server-pro",
		"version":       service.CompatSysinfoVersion(),
		"compat_target": c.compatService().Target(),
	}}
}

func (c *CompatPublicController) HandleVersion() mvc.Result {
	c.recordCompatAPIAudit(false, 200, "ok", "", nil)
	return mvc.Response{Object: iris.Map{
		"version":       service.CompatSysinfoVersion(),
		"server":        "rustdesk-api-server-pro",
		"compat_target": c.compatService().Target(),
	}}
}

func (c *CompatPublicController) HandleCompatTarget() mvc.Result {
	c.recordCompatAPIAudit(false, 200, "ok", "", nil)
	return mvc.Response{Object: c.compatService().Target()}
}

func (c *CompatPublicController) HandleFeatures() mvc.Result {
	c.recordCompatAPIAudit(false, 200, "ok", "", nil)
	return mvc.Response{Object: iris.Map{
		"address_book":           true,
		"audit":                  true,
		"file_transfer_audit":     true,
		"alarm_audit":            true,
		"device_group":           true,
		"user_group":             true,
		"strategy":               true,
		"record":                 true,
		"plugin_sign_passthrough": true,
		"compat_api_audit":       true,
		"compat_target":          c.compatService().Target(),
	}}
}

func (c *CompatPublicController) HandleClientConfig() mvc.Result {
	c.recordCompatAPIAudit(false, 200, "ok", "", nil)
	return mvc.Response{Object: iris.Map{
		"server": iris.Map{
			"name":    "rustdesk-api-server-pro",
		"version":       service.CompatSysinfoVersion(),
		},
		"compat_target": c.compatService().Target(),
		"features": iris.Map{
			"address_book":     true,
			"audit":            true,
			"record":           true,
			"compat_api_audit": true,
		},
		"login_options": c.compatService().LoginOptions().Options,
	}}
}

func (c *CompatPublicController) HandleSysinfoVer() mvc.Result {
	c.recordCompatAPIAudit(false, 200, "ok", "", nil)
	return mvc.Response{
		Text: service.CompatSysinfoVersion(),
	}
}

func (c *CompatPublicController) HandleOidcAuth() mvc.Result {
	body, _ := c.readBodyBytes()
	op := strings.TrimSpace(gjson.GetBytes(body, "op").String())
	op = strings.TrimPrefix(op, "oidc/")
	id := gjson.GetBytes(body, "id").String()
	uuid := gjson.GetBytes(body, "uuid").String()
	deviceOs := gjson.GetBytes(body, "deviceInfo.os").String()
	deviceType := gjson.GetBytes(body, "deviceInfo.type").String()
	deviceName := gjson.GetBytes(body, "deviceInfo.name").String()

	if op == "" {
		c.recordCompatAPIAudit(true, 400, "missing op", "", body)
		errJSON, _ := json.Marshal(iris.Map{"error": errcode.New(errcode.ERR2213.Code, errcode.ERR2213.Message).Error()})
		return mvc.Response{Object: iris.Map{"error": string(errJSON)}}
	}

	if op == "webauth" {
		cfg := config.GetServerConfig()
		if cfg.WebAuthn == nil || !cfg.WebAuthn.Enabled {
			c.recordCompatAPIAudit(true, 200, "webauth disabled", "", body)
			errJSON, _ := json.Marshal(iris.Map{"error": errcode.New(errcode.ERR2216.Code, errcode.ERR2216.Message).Error()})
			return mvc.Response{Object: iris.Map{"error": string(errJSON)}}
		}
		oauthService := service.NewOAuthProviderService(config.GetServerConfig(), c.Db)
		loginURL, pollToken, err := oauthService.StartWebauthLogin(c.currentBaseURL(), id, uuid, deviceOs, deviceType, deviceName)
		if err != nil {
			c.recordCompatAPIAudit(true, 500, err.Error(), "", body)
			errJSON, _ := json.Marshal(iris.Map{"error": err.Error()})
			return mvc.Response{Object: iris.Map{"error": string(errJSON)}}
		}
		c.recordCompatAPIAudit(false, 200, "ok", "", body)
		return mvc.Response{Object: iris.Map{
			"code": pollToken,
			"url":  loginURL,
		}}
	}

	oauthService := service.NewOAuthProviderService(config.GetServerConfig(), c.Db)
	authURL, pollToken, enabled, err := oauthService.BuildClientAuthURL(op, c.currentBaseURL(), id, uuid, deviceOs, deviceType, deviceName)
	if err != nil {
		c.recordCompatAPIAudit(true, 500, err.Error(), "", body)
		errJSON, _ := json.Marshal(iris.Map{"error": err.Error()})
		return mvc.Response{Object: iris.Map{"error": string(errJSON)}}
	}
	if !enabled {
		c.recordCompatAPIAudit(true, 200, "provider disabled", "", body)
		errJSON, _ := json.Marshal(iris.Map{"error": errcode.New(errcode.ERR2216.Code, errcode.ERR2216.Message).Error()})
		return mvc.Response{Object: iris.Map{"error": string(errJSON)}}
	}
	c.recordCompatAPIAudit(false, 200, "ok", "", body)
	return mvc.Response{Object: iris.Map{
		"code": pollToken,
		"url":  authURL,
	}}
}

func (c *CompatPublicController) HandleOidcAuthQuery() mvc.Result {
	code := strings.TrimSpace(c.Ctx.URLParamDefault("code", ""))

	if code == "" {
		c.recordCompatAPIAudit(true, 400, "missing code", "", nil)
		errJSON, _ := json.Marshal(iris.Map{"error": errcode.New(errcode.ERR2214.Code, errcode.ERR2214.Message).Error()})
		return mvc.Response{Object: iris.Map{"body": string(errJSON)}}
	}

	oauthService := service.NewOAuthProviderService(config.GetServerConfig(), c.Db)
	result, err := oauthService.ConsumePollAndExchange(code)
	if err != nil {
		c.recordCompatAPIAudit(true, 500, err.Error(), "", nil)
		errJSON, _ := json.Marshal(iris.Map{"error": err.Error()})
		return mvc.Response{Object: iris.Map{"body": string(errJSON)}}
	}
	if result == "" {
		c.recordCompatAPIAudit(false, 200, "pending", "", nil)
		pendingJSON, _ := json.Marshal(iris.Map{"error": "No authed oidc is found"})
		return mvc.Response{Object: iris.Map{"body": string(pendingJSON)}}
	}
	resultSummary := result
	if len(resultSummary) > 200 {
		resultSummary = resultSummary[:200]
	}
	c.recordCompatAPIAudit(false, 200, "ok", resultSummary, nil)
	return mvc.Response{Object: iris.Map{"body": result}}
}

func (c *CompatPublicController) HandleRecord() mvc.Result {
	offset, _ := strconv.ParseInt(c.Ctx.URLParamDefault("offset", "0"), 10, 64)
	body, err := c.readBodyBytes()
	if err != nil {
		c.recordCompatAPIAudit(true, 400, "error", err.Error(), nil)
		return c.fail(err)
	}
	err = c.compatService().HandleRecord(core.CompatRecordCommand{
		Op:       strings.ToLower(strings.TrimSpace(c.Ctx.URLParamDefault("type", ""))),
		FileName: c.Ctx.URLParamDefault("file", ""),
		Offset:   offset,
		Body:     body,
	})
	if err != nil {
		c.recordCompatAPIAudit(true, 400, "error", err.Error(), body)
		return c.fail(err)
	}
	c.recordCompatAPIAudit(false, 200, "ok", "", body)
	return c.ok()
}

func (c *CompatPublicController) HandleDevicesDeploy() mvc.Result {
	if c.Ctx.Method() == iris.MethodGet {
		c.recordCompatAPIAudit(true, 200, "NOT_ENABLED", "", nil)
		return mvc.Response{Object: iris.Map{"result": "NOT_ENABLED"}}
	}

	body, _ := c.readBodyBytes()
	result := c.compatService().HandleDeviceDeploy(core.CompatDeviceDeployCommand{
		RustdeskID: gjson.GetBytes(body, "id").String(),
		UUID:      gjson.GetBytes(body, "uuid").String(),
		PublicKey: gjson.GetBytes(body, "pk").String(),
	})
	statusCode := 200
	if result.Result == "INVALID_INPUT" {
		statusCode = 400
	}
	c.recordCompatAPIAudit(true, statusCode, result.Result, "", body)
	return mvc.Response{
		Object: iris.Map{
			"result": result.Result,
		},
	}
}

func (c *CompatPublicController) currentBaseURL() string {
	scheme := "http"
	if c.Ctx.Request().TLS != nil {
		scheme = "https"
	}
	if forwardedProto := strings.TrimSpace(c.Ctx.GetHeader("X-Forwarded-Proto")); forwardedProto == "https" {
		scheme = "https"
	}
	host := strings.TrimSpace(c.Ctx.Host())
	if host == "" {
		host = strings.TrimSpace(c.Ctx.Request().Host)
	}
	return scheme + "://" + host
}
