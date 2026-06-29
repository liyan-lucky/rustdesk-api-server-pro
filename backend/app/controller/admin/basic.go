package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
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
	return c.Ctx.Values().Get(config.AdminUserKey).(*model.User)
}

func (c *basicController) GetToken() string {
	return c.Ctx.Values().Get(config.AdminAuthTokenString).(string)
}

func (c *basicController) GetAuthToken() *model.AuthToken {
	return c.Ctx.Values().Get(config.AdminAuthToken).(*model.AuthToken)
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

func (c *basicController) response(code int, data interface{}, message string) mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"code":    code,
			"message": message,
			"data":    data,
		},
	}
}
