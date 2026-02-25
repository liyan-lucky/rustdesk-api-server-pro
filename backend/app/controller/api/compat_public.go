package api

import (
	"strconv"
	"strings"

	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// CompatPublicController provides compatibility endpoints used by newer RustDesk clients.
// Unsupported features return stable "error" payloads instead of 404.
type CompatPublicController struct {
	basicController
}

func (c *CompatPublicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "sysinfo_ver", "HandleSysinfoVer")
	b.Handle("POST", "oidc/auth", "HandleOidcAuth")
	b.Handle("GET", "oidc/auth-query", "HandleOidcAuthQuery")
	b.Handle("POST", "record", "HandleRecord")
}

func (c *CompatPublicController) HandleSysinfoVer() mvc.Result {
	return mvc.Response{
		Text: service.CompatSysinfoVersion,
	}
}

func (c *CompatPublicController) HandleOidcAuth() mvc.Result {
	result := c.compatService().OidcAuth()
	return mvc.Response{
		Object: iris.Map{
			"error":   result.Error,
			"enabled": result.Enabled,
			"url":     result.URL,
		},
	}
}

func (c *CompatPublicController) HandleOidcAuthQuery() mvc.Result {
	result := c.compatService().OidcAuthQuery()
	return mvc.Response{
		Object: iris.Map{
			"error":   result.Error,
			"enabled": result.Enabled,
			"user":    result.User,
		},
	}
}

func (c *CompatPublicController) HandleRecord() mvc.Result {
	offset, _ := strconv.ParseInt(c.Ctx.URLParamDefault("offset", "0"), 10, 64)
	body, err := c.readBodyBytes()
	if err != nil {
		return c.fail(err)
	}
	err = c.compatService().HandleRecord(core.CompatRecordCommand{
		Op:       strings.ToLower(strings.TrimSpace(c.Ctx.URLParamDefault("type", ""))),
		FileName: c.Ctx.URLParamDefault("file", ""),
		Offset:   offset,
		Body:     body,
	})
	if err != nil {
		return c.fail(err)
	}
	return c.ok()
}
