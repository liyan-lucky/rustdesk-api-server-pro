package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type ContainerLogController struct {
	basicController
}

func (c *ContainerLogController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/container-logs/list", "HandleList")
	b.Handle("DELETE", "/container-logs/clear", "HandleClear")
}

func (c *ContainerLogController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	level := c.Ctx.URLParamDefault("level", "")
	source := c.Ctx.URLParamDefault("source", "")
	path := c.Ctx.URLParamDefault("path", "")
	method := c.Ctx.URLParamDefault("method", "")
	status_code := c.Ctx.URLParamDefault("status_code", "")
	user_name := c.Ctx.URLParamDefault("user_name", "")
	client_ip := c.Ctx.URLParamDefault("client_ip", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.ContainerLog{})
		if level != "" {
			q.Where("container_log.level = ?", level)
		}
		if source != "" {
			q.Where("container_log.source = ?", source)
		}
		if path != "" {
			q.Where("container_log.path LIKE ?", "%"+path+"%")
		}
		if method != "" {
			q.Where("container_log.method = ?", method)
		}
		if status_code != "" {
			q.Where("container_log.status_code = ?", status_code)
		}
		if user_name != "" {
			q.Where("container_log.user_name LIKE ?", "%"+user_name+"%")
		}
		if client_ip != "" {
			q.Where("container_log.client_ip = ?", client_ip)
		}
		if created_at_0 != "" && created_at_1 != "" {
			q.Where("container_log.created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		q.Desc("id")
		return q
	}

	pagination := db.NewPagination(currentPage, pageSize)
	list := make([]model.ContainerLog, 0)
	err := pagination.Paginate(query, &model.ContainerLog{}, &list)
	if err != nil {
		return c.dbError(err)
	}

	records := make([]iris.Map, 0, len(list))
	for _, item := range list {
		records = append(records, iris.Map{
			"id":          item.Id,
			"timestamp":   item.Timestamp.Format(config.TimeFormat),
			"level":       item.Level,
			"source":      item.Source,
			"message":     item.Message,
			"method":      item.Method,
			"path":        item.Path,
			"status_code": item.StatusCode,
			"duration_ms": item.DurationMs,
			"client_ip":   item.ClientIP,
			"user_id":     item.UserId,
			"user_name":   item.UserName,
			"user_agent":  item.UserAgent,
			"request_id":  item.RequestId,
			"error_msg":   item.ErrorMsg,
			"created_at":  item.CreatedAt.Format(config.TimeFormat),
		})
	}

	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": records,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *ContainerLogController) HandleClear() mvc.Result {
	_, err := c.Db.Exec("DELETE FROM container_log")
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}
