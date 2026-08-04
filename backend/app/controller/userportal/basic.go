package userportal

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type basicController struct {
	Ctx iris.Context
	Db  *xorm.Engine
}

func (c *basicController) GetUser() *model.User {
	v := c.Ctx.Values().Get(config.WebUserKey)
	if v == nil {
		return nil
	}
	return v.(*model.User)
}

func (c *basicController) Success(data interface{}, message string) mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"code":    200,
			"message": message,
			"data":    data,
		},
	}
}

func (c *basicController) Error(data interface{}, message string) mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"code":    500,
			"message": message,
			"data":    data,
		},
	}
}

func (c *basicController) dbError(err error) mvc.Result {
	return c.Error(nil, errcode.Errorf(errcode.ERRB010.Code, errcode.ERRB010.Message+": "+err.Error()).Error())
}
