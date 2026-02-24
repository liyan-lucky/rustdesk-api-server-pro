package api

import "github.com/kataras/iris/v12/mvc"

// CompatLicController provides compatibility endpoints under /lic/web/api used by newer clients/plugins.
// This project does not implement the official plugin signing service, but returns a stable JSON shape
// to avoid 404/decode failures in clients.
type CompatLicController struct {
	basicController
}

type pluginSignReq struct {
	PluginID string `json:"plugin_id"`
	Version  string `json:"version"`
	Msg      []byte `json:"msg"`
}

type pluginSignResp struct {
	SignedMsg []byte `json:"signed_msg"`
}

func (c *CompatLicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "plugin-sign", "HandlePluginSign")
}

func (c *CompatLicController) HandlePluginSign() mvc.Result {
	var req pluginSignReq
	// Always return the expected JSON shape even when payload is invalid, so client plugins do not fail on decode.
	if err := c.Ctx.ReadJSON(&req); err != nil {
		return mvc.Response{
			Object: pluginSignResp{SignedMsg: []byte{}},
		}
	}
	// Placeholder implementation: echo bytes back as a deterministic non-empty response.
	// Real signature generation requires the official licensing/signing service.
	return mvc.Response{
		Object: pluginSignResp{SignedMsg: req.Msg},
	}
}
