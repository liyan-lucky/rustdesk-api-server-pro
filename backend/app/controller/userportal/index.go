package userportal

import (
	"rustdesk-api-server-pro/internal/errcode"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type IndexController struct {
	basicController
}

func (c *IndexController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/userinfo", "HandleUserInfo")
}

func (c *IndexController) HandleUserInfo() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errcode.ErrUnauthorized.Error())
	}
	roles := []string{}
	if user.IsAdmin {
		roles = append(roles, "R_SUPER")
	} else {
		roles = append(roles, "R_USER")
	}
	return c.Success(iris.Map{
		"userId":          user.Id,
		"userName":        user.Name,
		"email":           user.Email,
		"username":        user.Username,
		"note":            user.Note,
		"licensedDevices": user.LicensedDevices,
		"isAdmin":         user.IsAdmin,
		"roles":           roles,
		"buttons":         []string{},
	}, "ok")
}
