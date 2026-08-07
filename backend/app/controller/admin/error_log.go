package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/errcode"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type ErrorLogController struct {
	basicController
}

func (c *ErrorLogController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/error-logs/list", "HandleList")
	b.Handle("DELETE", "/error-logs/clear", "HandleClear")
}

func (c *ErrorLogController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	code := c.Ctx.URLParamDefault("code", "")
	module := c.Ctx.URLParamDefault("module", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.ErrorLog{})
		if code != "" {
			q.Where("error_log.code = ?", code)
		}
		if module != "" {
			q.Where("error_log.module = ?", module)
		}
		if created_at_0 != "" && created_at_1 != "" {
			q.Where("error_log.created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		q.Desc("id")
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	list := make([]model.ErrorLog, 0)
	err := pagination.Paginate(query, &model.ErrorLog{}, &list)
	if err != nil {
		return c.dbError(err)
	}

	records := make([]iris.Map, 0, len(list))
	for _, item := range list {
		records = append(records, iris.Map{
			"id":         item.Id,
			"code":       item.Code,
			"message":    item.Message,
			"module":     item.Module,
			"path":       item.Path,
			"method":     item.Method,
			"user_id":    item.UserId,
			"user_name":  item.UserName,
			"client_ip":  item.ClientIP,
			"user_agent": item.UserAgent,
			"created_at": item.CreatedAt.Format(config.TimeFormat),
		})
	}

	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": records,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *ErrorLogController) HandleClear() mvc.Result {
	user := c.GetUser()
	if user == nil || !user.IsAdmin {
		return c.Error(nil, errcode.New(errcode.ERR1010.Code, errcode.ERR1010.Message).Error())
	}
	_, err := c.Db.Exec("DELETE FROM error_log")
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}
