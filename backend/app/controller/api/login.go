package api

import (
	apiform "rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/config"

	"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	basicController
	Cfg *config.ServerConfig
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/login-options", "HandleLoginOptions")
}

func (c *LoginController) PostLogin() mvc.Result {
	var loginForm apiform.LoginForm
	if err := c.readJSONBody(&loginForm); err != nil {
		return c.fail(err)
	}

	result := c.loginService().HandleLogin(loginForm)
	return mvc.Response{Object: result}
}

func (c *LoginController) HandleLoginOptions() mvc.Result {
	result := c.compatService().LoginOptions()
	return mvc.Response{
		Object: result.Options,
	}
}
