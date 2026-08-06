package middleware

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/helper"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"xorm.io/xorm"
)

var skipPaths = map[string]bool{
	"/api/heartbeat": true,
}

var skipPrefixes = []string{
	"/admin/container-logs/",
}

func ContainerLogRecorder(app *iris.Application) iris.Handler {
	return func(context iris.Context) {
		start := time.Now()
		path := context.Request().URL.Path
		method := context.Method()
		clientIP := context.RemoteAddr()
		userAgent := context.GetHeader("User-Agent")
		requestID := context.GetHeader("X-Request-ID")

		context.Next()

		if skipPaths[path] {
			return
		}
		for _, prefix := range skipPrefixes {
			if len(path) >= len(prefix) && path[:len(prefix)] == prefix {
				return
			}
		}

		duration := time.Since(start)
		statusCode := context.GetStatusCode()

		db, ok := helper.GetAppDependency(app, "*xorm.Engine").(*xorm.Engine)
		if !ok || db == nil {
			return
		}

		level := "INFO"
		if statusCode >= 500 {
			level = "ERROR"
		} else if statusCode >= 400 {
			level = "WARN"
		}

		source := "http"
		if len(path) >= 6 && path[:6] == "/admin" {
			source = "admin"
		} else if len(path) >= 4 && path[:4] == "/api" {
			source = "client"
		} else if len(path) >= 12 && path[:12] == "/user-portal" {
			source = "portal"
		}

		var userId int
		var userName string
		if v := context.Values().Get(config.AdminUserKey); v != nil {
			if u, ok := v.(*model.User); ok {
				userId = u.Id
				userName = u.Username
			}
		} else if v := context.Values().Get(config.CurrentUserKey); v != nil {
			if u, ok := v.(*model.User); ok {
				userId = u.Id
				userName = u.Username
			}
		} else if v := context.Values().Get(config.WebUserKey); v != nil {
			if u, ok := v.(*model.User); ok {
				userId = u.Id
				userName = u.Username
			}
		}

		msg := method + " " + path + " " + strconv.Itoa(statusCode)
		sc := statusCode
		dur := duration.Milliseconds()
		uid := userId
		uname := userName

		go func() {
			_, _ = db.InsertOne(&model.ContainerLog{
				Timestamp:  start,
				Level:      level,
				Source:     source,
				Message:    msg,
				Method:     method,
				Path:       path,
				StatusCode: sc,
				DurationMs: dur,
				ClientIP:   clientIP,
				UserId:     uid,
				UserName:   uname,
				UserAgent:  userAgent,
				RequestId:  requestID,
			})
		}()
	}
}

func RecordContainerEvent(db *xorm.Engine, level string, source string, message string) {
	if db == nil {
		return
	}
	_, _ = db.InsertOne(&model.ContainerLog{
		Timestamp: time.Now(),
		Level:     level,
		Source:    source,
		Message:   message,
	})
}
