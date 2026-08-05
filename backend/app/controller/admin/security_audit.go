package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type SecurityAuditController struct {
	basicController
}

func (c *SecurityAuditController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/security-audit/list", "HandleList")
	b.Handle("DELETE", "/security-audit/clear", "HandleClear")
}

func (c *SecurityAuditController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	username := c.Ctx.URLParamDefault("username", "")
	event := c.Ctx.URLParamDefault("event", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.SecurityAudit{})
		user := c.GetUser()
		if user == nil {
			return q.Where("1 = 0")
		}
		if !user.IsAdmin {
			q = q.Where("user_id = ?", user.Id)
		} else if username != "" {
			q.Where("username LIKE ?", "%"+username+"%")
		}
		if event != "" {
			q.Where("event = ?", event)
		}
		q.Desc("id")
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	list := make([]model.SecurityAudit, 0)
	if err := pagination.Paginate(query, &model.SecurityAudit{}, &list); err != nil {
		return c.dbError(err)
	}

	records := make([]iris.Map, 0, len(list))
	for _, item := range list {
		records = append(records, iris.Map{
			"id":         item.Id,
			"user_id":    item.UserId,
			"username":   item.Username,
			"event":      item.Event,
			"ip":         item.IP,
			"user_agent": item.UserAgent,
			"success":    item.Success,
			"reason":     item.Reason,
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

func (c *SecurityAuditController) HandleClear() mvc.Result {
	_, err := c.Db.Exec("DELETE FROM security_audit")
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}
