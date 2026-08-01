package userportal

import (
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
		return c.Error(nil, "unauthorized")
	}
	roles := []string{}
	if user.IsAdmin {
		roles = append(roles, "R_SUPER")
	}
	return c.Success(iris.Map{
		"userId":   user.Id,
		"userName": user.Name,
		"email":    user.Email,
		"isAdmin":  user.IsAdmin,
		"roles":    roles,
		"buttons":  []string{},
	}, "ok")
}
