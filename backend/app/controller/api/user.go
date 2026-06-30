package api

import (
	"time"

	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"
	v2service "rustdesk-api-server-pro/internal/service"
	"rustdesk-api-server-pro/internal/transport/httpdto"
	"rustdesk-api-server-pro/util"

	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	basicController
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/currentUser", "HandleCurrentUser")
	b.Handle("POST", "/currentUser", "HandleCurrentUser")
	b.Handle("GET", "/logout", "HandleLogout")
	b.Handle("POST", "/logout", "HandleLogout")
	b.Handle("DELETE", "/logout", "HandleLogout")
}

func (c *UserController) HandleCurrentUser() mvc.Result {
	user := c.GetUser()
	return mvc.Response{
		Object: httpdto.NewUserResponse(c.userService().CurrentUserView(user.Name, user.Email, user.Note, user.Status, user.IsAdmin)),
	}
}

func (c *UserController) GetUsers() mvc.Result {
	user := c.GetUser()
	hasAccessibleParam := c.Ctx.Request().URL.Query().Has("accessible")
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	status := c.Ctx.URLParamIntDefault("status", 1)

	result, err := c.userService().ListUsers(core.UserListQuery{
		RequestUserID:      user.Id,
		RequestUserIsAdmin: user.IsAdmin,
		HasAccessibleParam: hasAccessibleParam,
		Current:            current,
		PageSize:           pageSize,
		Status:             status,
	})
	if err != nil {
		if err == v2service.ErrAdminRequired {
			return c.failMsg("Admin required!")
		}
		return c.fail(err)
	}
	return mvc.Response{
		Object: httpdto.NewUserListResponse(result),
	}
}

func (c *UserController) HandleLogout() mvc.Result {
	user := c.GetUser()
	rustdeskID := c.Ctx.URLParamDefault("id", "")
	if rustdeskID == "" {
		rustdeskID = c.Ctx.URLParamDefault("rustdesk_id", "")
	}

	if c.Ctx.Request().ContentLength > 0 {
		var f api.LoginForm
		if err := c.readJSONBody(&f); err == nil && f.RustdeskId != "" {
			rustdeskID = f.RustdeskId
		}
	}

	logoutMode := "token"
	if rustdeskID != "" {
		logoutMode = "rustdesk_id"
		if err := c.userService().LogoutByRustdeskID(rustdeskID); err != nil {
			c.recordUserSecurityAudit(user, "api_logout", false, "logout_by_rustdesk_id: "+err.Error())
			return c.fail(err)
		}
	} else {
		token := c.GetToken()
		if token != "" {
			_, err := c.Db.Where("token_hash = ?", util.Sha256Hex(token)).Cols("expired", "status").Update(&model.AuthToken{
				Expired: time.Now().Add(-time.Second),
				Status:  0,
			})
			if err != nil {
				c.recordUserSecurityAudit(user, "api_logout", false, "logout_by_token_hash: "+err.Error())
				return c.fail(err)
			}
			// Backward compatibility for sessions issued before token_hash existed.
			_, err = c.Db.Where("token = ?", token).Cols("expired", "status").Update(&model.AuthToken{
				Expired: time.Now().Add(-time.Second),
				Status:  0,
			})
			if err != nil {
				c.recordUserSecurityAudit(user, "api_logout", false, "logout_by_legacy_token: "+err.Error())
				return c.fail(err)
			}
		}
	}

	c.recordUserSecurityAudit(user, "api_logout", true, logoutMode)
	return c.okText("ok")
}

func (c *UserController) recordUserSecurityAudit(user *model.User, event string, success bool, reason string) {
	cmd := core.SecurityAuditCreateCommand{
		Event:     event,
		IP:        c.Ctx.RemoteAddr(),
		UserAgent: c.Ctx.GetHeader("User-Agent"),
		Success:   success,
		Reason:    reason,
	}
	if user != nil {
		cmd.UserID = user.Id
		cmd.Username = user.Username
	}
	_ = c.auditService().CreateSecurityAudit(cmd)
}
