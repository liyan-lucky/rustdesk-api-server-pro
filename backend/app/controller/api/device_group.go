package api

import (
	"rustdesk-api-server-pro/app/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

// DeviceGroupController provides a compatibility endpoint for newer RustDesk clients.
// This project does not yet model device groups, so it returns an empty list instead of 404.
type DeviceGroupController struct {
	basicController
}

func (c *DeviceGroupController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "device-group/accessible", "HandleAccessible")
}

func (c *DeviceGroupController) HandleAccessible() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 100)
	if current < 1 {
		current = 1
	}
	if pageSize <= 0 {
		pageSize = 100
	}

	total, err := c.Db.Count(&model.DeviceGroup{})
	if err != nil {
		return c.fail(err)
	}
	var groups []model.DeviceGroup
	if err := c.Db.Asc("name").Limit(pageSize, (current-1)*pageSize).Find(&groups); err != nil {
		return c.fail(err)
	}
	data := make([]iris.Map, 0, len(groups))
	for _, g := range groups {
		data = append(data, iris.Map{"guid": g.Guid, "name": g.Name})
	}
	return mvc.Response{
		Object: iris.Map{
			"total":    total,
			"current":  current,
			"pageSize": pageSize,
			"data":     data,
		},
	}
}
