package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/internal/core"
	v2service "rustdesk-api-server-pro/internal/service"
	"rustdesk-api-server-pro/internal/transport/httpdto"

	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	basicController
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/currentUser", "HandleCurrentUser")
}

func (c *UserController) HandleCurrentUser() mvc.Result {
	user := c.GetUser()
	return mvc.Response{
		Object: httpdto.NewUserResponse(c.userService().CurrentUserView(user.Name, user.Email, user.Note, user.Status, user.IsAdmin)),
	}
}

func (c *UserController) GetUsers() mvc.Result {
	user := c.GetUser()
	hasAccessibleParam := c.Ctx.Request().URL.Query().Has("accessible")
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	status := c.Ctx.URLParamIntDefault("status", 1)

	result, err := c.userService().ListUsers(core.UserListQuery{
		RequestUserID:      user.Id,
		RequestUserIsAdmin: user.IsAdmin,
		HasAccessibleParam: hasAccessibleParam,
		Current:            current,
		PageSize:           pageSize,
		Status:             status,
	})
	if err != nil {
		if err == v2service.ErrAdminRequired {
			return c.failMsg("Admin required!")
		}
		return c.fail(err)
	}
	return mvc.Response{
		Object: httpdto.NewUserListResponse(result),
	}
}

func (c *UserController) PostLogout() mvc.Result {
	var f api.LoginForm
	err := c.readJSONBody(&f)
	if err != nil {
		return c.fail(err)
	}
	err = c.userService().LogoutByRustdeskID(f.RustdeskId)
	if err != nil {
		return c.fail(err)
	}
	return c.okText("ok")
}
