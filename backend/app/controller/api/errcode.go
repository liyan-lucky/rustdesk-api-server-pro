package api

import (
	"rustdesk-api-server-pro/internal/errcode"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ErrCodeController struct {
	Ctx iris.Context
}

func (c *ErrCodeController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/errcode", "Get")
}

func (c *ErrCodeController) Get() mvc.Response {
	entries := errcode.List()
	return mvc.Response{
		Object: map[string]interface{}{
			"errorCodes": entries,
		},
	}
}
