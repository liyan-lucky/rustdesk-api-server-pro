package api

import (
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/transport/httpdto"
)

type PeerController struct {
	basicController
}

func (c *PeerController) GetPeers() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	status := c.Ctx.URLParamIntDefault("status", 1)

	user := c.GetUser()
	result, err := c.peerService().ListPeers(core.PeerListQuery{
		UserID:   user.Id,
		Username: user.Username,
		Current:  current,
		PageSize: pageSize,
		Status:   status,
	})
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{
		Object: httpdto.NewPeerListResponse(result, status, user.Username),
	}
}
