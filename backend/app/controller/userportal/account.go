package userportal

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/errcode"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type AccountController struct{ basicController }

func (c *AccountController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/overview", "HandleOverview")
	b.Handle("GET", "/sessions", "HandleSessions")
	b.Handle("DELETE", "/sessions", "HandleSessionsDelete")
	b.Handle("GET", "/security-events", "HandleSecurityEvents")
}

func (c *AccountController) HandleOverview() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errcode.ErrUnauthorized.Error())
	}
	now := time.Now().Format(config.TimeFormat)
	sessions, err := c.Db.Where("user_id = ? and status = 1 and expired > ?", user.Id, now).Count(new(model.AuthToken))
	if err != nil {
		return c.dbError(err)
	}
	books, err := c.Db.Where("user_id = ?", user.Id).Count(new(model.AddressBook))
	if err != nil {
		return c.dbError(err)
	}
	events, err := c.Db.Where("user_id = ?", user.Id).Count(new(model.SecurityAudit))
	if err != nil {
		return c.dbError(err)
	}
	deviceIDs := make([]string, 0)
	if err = c.Db.Table(new(model.AuthToken)).Where("user_id = ? and rustdesk_id <> ''", user.Id).Distinct("rustdesk_id").Find(&deviceIDs); err != nil {
		return c.dbError(err)
	}
	return c.Success(iris.Map{"devices": len(deviceIDs), "sessions": sessions, "address_books": books, "security_events": events, "licensed_devices": user.LicensedDevices}, "ok")
}

func (c *AccountController) HandleSessions() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errcode.ErrUnauthorized.Error())
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	size := c.Ctx.URLParamIntDefault("size", 10)
	query := func() *xorm.Session {
		return c.Db.Table(new(model.AuthToken)).Where("user_id = ? and status = 1", user.Id).Desc("updated_at")
	}
	pagination := db.NewPagination(current, size)
	var sessions []model.AuthToken
	if err := pagination.Paginate(query, new(model.AuthToken), &sessions); err != nil {
		return c.dbError(err)
	}
	currentToken, _ := c.Ctx.Values().Get(config.WebAuthToken).(*model.AuthToken)
	records := make([]iris.Map, 0, len(sessions))
	for _, session := range sessions {
		records = append(records, iris.Map{"id": session.Id, "rustdesk_id": session.RustdeskId, "device_name": session.DeviceName, "device_os": session.DeviceOs, "device_type": session.DeviceType, "expired": session.Expired.Format(config.TimeFormat), "created_at": session.CreatedAt.Format(config.TimeFormat), "current": currentToken != nil && currentToken.Id == session.Id})
	}
	return c.Success(iris.Map{"total": pagination.TotalCount, "records": records, "current": current, "size": size}, "ok")
}

func (c *AccountController) HandleSessionsDelete() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errcode.ErrUnauthorized.Error())
	}
	var ids []int
	if err := c.Ctx.ReadJSON(&ids); err != nil {
		return c.dbError(err)
	}
	currentToken, _ := c.Ctx.Values().Get(config.WebAuthToken).(*model.AuthToken)
	if currentToken != nil {
		filtered := ids[:0]
		for _, id := range ids {
			if id != currentToken.Id {
				filtered = append(filtered, id)
			}
		}
		ids = filtered
	}
	if len(ids) > 0 {
		if _, err := c.Db.Where("user_id = ?", user.Id).In("id", ids).Cols("status").Update(&model.AuthToken{Status: 0}); err != nil {
			return c.dbError(err)
		}
	}
	return c.Success(nil, "ok")
}

func (c *AccountController) HandleSecurityEvents() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errcode.ErrUnauthorized.Error())
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	size := c.Ctx.URLParamIntDefault("size", 10)
	query := func() *xorm.Session {
		return c.Db.Table(new(model.SecurityAudit)).Where("user_id = ?", user.Id).Desc("created_at")
	}
	pagination := db.NewPagination(current, size)
	var events []model.SecurityAudit
	if err := pagination.Paginate(query, new(model.SecurityAudit), &events); err != nil {
		return c.dbError(err)
	}
	records := make([]iris.Map, 0, len(events))
	for _, event := range events {
		records = append(records, iris.Map{"id": event.Id, "event": event.Event, "ip": event.IP, "user_agent": event.UserAgent, "success": event.Success, "reason": event.Reason, "created_at": event.CreatedAt.Format(config.TimeFormat)})
	}
	return c.Success(iris.Map{"total": pagination.TotalCount, "records": records, "current": current, "size": size}, "ok")
}
