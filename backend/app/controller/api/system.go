package api

import (
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/transport/httpdto"

	"github.com/kataras/iris/v12/mvc"
)

type SystemController struct {
	basicController
}

func (c *SystemController) PostHeartbeat() mvc.Result {
	// {"conns":[762],"id":"182921366","modified_at":1725698100,"uuid":"xxx","ver":1002070}
	var form api.HeartbeatForm
	err := c.readJSONBody(&form)
	if err != nil {
		return c.fail(err)
	}

	result, err := c.systemService().HandleHeartbeat(core.HeartbeatCommand{
		RustdeskID: form.RustdeskId,
		UUID:       form.Uuid,
		ConnCount:  len(form.Conns),
	})
	if err != nil {
		return c.fail(err)
	}

	return mvc.Response{
		Object: httpdto.NewHeartbeatResponse(result),
	}
}

func (c *SystemController) PostSysinfo() mvc.Result {
	var form api.DeviceForm
	err := c.readJSONBody(&form)
	if err != nil {
		return c.fail(err)
	}

	result, err := c.systemService().UpdateSysinfo(core.SysinfoUpdateCommand{
		RustdeskID: form.RustdeskId,
		CPU:        form.Cpu,
		Hostname:   form.Hostname,
		Memory:     form.Memory,
		OS:         form.Os,
		Username:   form.Username,
		UUID:       form.Uuid,
		Version:    form.Version,
	})
	if err != nil {
		return c.fail(err)
	}
	if !result.Updated {
		return mvc.Response{
			Text: "ID_NOT_FOUND",
		}
	}

	return mvc.Response{
		Text: "SYSINFO_UPDATED",
	}
}
