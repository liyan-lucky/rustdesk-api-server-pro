package admin

import (
	"log"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/app/service"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/internal/repository"
	v2service "rustdesk-api-server-pro/internal/service"
	"strings"

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
	user := c.GetUser()
	userId := 0
	userName := ""
	if user != nil {
		userId = user.Id
		userName = user.Username
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic in RecordErrorLog (admin/Error): %v", r)
			}
		}()
		service.RecordErrorLog(c.Db, extractErrCode(message), message, "admin", c.Ctx.Path(), c.Ctx.Method(), userId, userName, c.Ctx.RemoteAddr(), c.Ctx.GetHeader("User-Agent"))
	}()
	return c.response(500, data, message)
}

func extractErrCode(message string) string {
	message = strings.TrimSpace(message)
	if message == "" || !strings.HasPrefix(message, "ERR-") {
		return ""
	}
	idx := strings.Index(message, ":")
	if idx <= 0 {
		return message
	}
	code := message[:idx]
	if len(code) < 5 {
		return ""
	}
	return code
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
