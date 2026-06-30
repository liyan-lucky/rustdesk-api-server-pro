package middleware

import (
	"rustdesk-api-server-pro/config"

	"github.com/kataras/iris/v12"
)

func RequestLogger() iris.Handler {
	return func(context iris.Context) {
		if config.GetServerConfig().DebugMode && context.Request().RequestURI != "/api/heartbeat" {
			context.Application().Logger().Infof("▶ %s:%s", context.Method(), context.Request().RequestURI)
		}

		context.Next()
	}
}
