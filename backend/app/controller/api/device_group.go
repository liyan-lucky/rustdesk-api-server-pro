package api

import (
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
	return mvc.Response{
		Object: iris.Map{
			"total": 0,
			"data":  []iris.Map{},
		},
	}
}

