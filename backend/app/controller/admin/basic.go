package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/internal/repository"
	v2service "rustdesk-api-server-pro/internal/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type basicController struct {
	Ctx iris.Context
	Db  *xorm.Engine
}

func (c *basicController) GetUser() *model.User {
	v := c.Ctx.Values().Get(config.AdminUserKey)
	if v == nil {
		return nil
	}
	return v.(*model.User)
}

func (c *basicController) GetToken() string {
	v := c.Ctx.Values().Get(config.AdminAuthTokenString)
	if v == nil {
		return ""
	}
	return v.(string)
}

func (c *basicController) GetAuthToken() *model.AuthToken {
	v := c.Ctx.Values().Get(config.AdminAuthToken)
	if v == nil {
		return nil
	}
	return v.(*model.AuthToken)
}

func (c *basicController) auditService() *v2service.AuditService {
	return v2service.NewAuditService(repository.NewXormAuditRepository(c.Db))
}

func (c *basicController) Success(data interface{}, message string) mvc.Result {
	return c.response(200, data, message)
}

func (c *basicController) Error(data interface{}, message string) mvc.Result {
	return c.response(500, data, message)
}

func (c *basicController) dbError(err error) mvc.Result {
	return c.Error(nil, errcode.Errorf(errcode.ERRB010.Code, errcode.ERRB010.Message+": "+err.Error()).Error())
}

func (c *basicController) response(code int, data interface{}, message string) mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"code":    code,
			"message": message,
			"data":    data,
		},
	}
}
