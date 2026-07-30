package userportal

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type DevicesController struct {
	basicController
}

func (c *DevicesController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/devices/my", "HandleMyDevices")
}

func (c *DevicesController) HandleMyDevices() mvc.Result {
	user := c.GetUser()
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)

	type idRow struct {
		RustdeskId string `xorm:"'rustdesk_id'"`
	}
	idRows := make([]idRow, 0)
	if err := c.Db.Table(&model.AuthToken{}).
		Where("user_id = ? and is_admin = 0 and status = 1 and expired > ?", user.Id, time.Now().Format(config.TimeFormat)).
		Cols("rustdesk_id").
		Find(&idRows); err != nil {
		return c.Error(nil, err.Error())
	}

	myRustdeskIds := make([]string, 0, len(idRows))
	seen := make(map[string]bool)
	for _, r := range idRows {
		if r.RustdeskId != "" && !seen[r.RustdeskId] {
			seen[r.RustdeskId] = true
			myRustdeskIds = append(myRustdeskIds, r.RustdeskId)
		}
	}

	if len(myRustdeskIds) == 0 {
		return c.Success(iris.Map{
			"total":   0,
			"records": []iris.Map{},
			"current": currentPage,
			"size":    pageSize,
		}, "ok")
	}

	query := func() *xorm.Session {
		return c.Db.Table(&model.Device{}).In("rustdesk_id", myRustdeskIds).Desc("updated_at")
	}

	pagination := db.NewPagination(currentPage, pageSize)
	deviceList := make([]model.Device, 0)

	err := pagination.Paginate(query, &model.Device{}, &deviceList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0, len(deviceList))
	for _, d := range deviceList {
		list = append(list, iris.Map{
			"id":          d.Id,
			"rustdesk_id": d.RustdeskId,
			"hostname":    d.Hostname,
			"username":    d.Username,
			"uuid":        d.Uuid,
			"version":     d.Version,
			"os":          d.Os,
			"cpu":         d.Cpu,
			"memory":      d.Memory,
			"disabled":    d.Disabled,
			"is_online":   d.IsOnline,
			"conns":       d.Conns,
			"updated_at":  d.UpdatedAt.Format(config.TimeFormat),
		})
	}

	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}
