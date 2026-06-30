package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/util"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type SessionsController struct {
	basicController
}

func (c *SessionsController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/sessions/list", "HandleList")
	b.Handle("POST", "/sessions/kill", "HandleKill")
}

func (c *SessionsController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	username := c.Ctx.URLParamDefault("username", "")
	created_at_0 := c.Ctx.URLParamDefault("created_at[0]", "")
	created_at_1 := c.Ctx.URLParamDefault("created_at[1]", "")
	query := func() *xorm.Session {
		q := c.Db.Table(&model.AuthToken{})
		q.Join("INNER", &model.User{}, "auth_token.user_id = user.id")
		q.Where("auth_token.status = 1 and auth_token.is_admin = 0")
		if username != "" {
			q.Where("user.username = ?", username)
		}
		if created_at_0 != "" && created_at_1 != "" {
			q.Where("created_at BETWEEN ? AND ?", created_at_0, created_at_1)
		}
		q.Desc("auth_token.id")
		return q
	}

	type Session struct {
		model.AuthToken `xorm:"extends"`
		model.User      `xorm:"extends"`
	}

	pagination := db.NewPagination(currentPage, pageSize)
	sessionList := make([]Session, 0)
	err := pagination.Paginate(query, &Session{}, &sessionList)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0)
	for _, s := range sessionList {
		list = append(list, iris.Map{
			"id":          s.AuthToken.Id,
			"username":    s.User.Username,
			"rustdesk_id": s.AuthToken.RustdeskId,
			"expired":     s.AuthToken.Expired.Format(config.TimeFormat),
			"created_at":  s.AuthToken.CreatedAt.Format(config.TimeFormat),
		})
	}
	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *SessionsController) HandleKill() mvc.Result {
	type killParams struct {
		Ids []int `json:"ids"`
	}
	var params killParams
	err := c.Ctx.ReadJSON(&params)
	if err != nil {
		c.recordSessionOperationAudit("admin_session_kill", "", nil, iris.Map{"ids": params.Ids}, "failure", err.Error())
		return c.Error(nil, err.Error())
	}
	ids := util.RemoveElement(params.Ids, 1)
	if len(ids) == 0 {
		c.recordSessionOperationAudit("admin_session_kill", "", nil, iris.Map{"ids": params.Ids}, "failure", "NoSessionIds")
		return c.Error(nil, "NoSessionIds")
	}

	beforeSessions := make([]model.AuthToken, 0)
	_ = c.Db.In("id", ids).Find(&beforeSessions)
	beforeAudit := make([]iris.Map, 0)
	for _, s := range beforeSessions {
		session := s
		beforeAudit = append(beforeAudit, sanitizeSessionForAudit(&session))
	}

	_, err = c.Db.In("id", ids).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})
	if err != nil {
		c.recordSessionOperationAudit("admin_session_kill", auditIDsResource(ids), beforeAudit, iris.Map{"ids": ids, "status": 0}, "failure", err.Error())
		return c.Error(nil, err.Error())
	}

	c.recordSessionOperationAudit("admin_session_kill", auditIDsResource(ids), beforeAudit, iris.Map{"ids": ids, "status": 0}, "success", "")
	return c.Success(nil, "SessionKillSuccess")
}

func (c *SessionsController) recordSessionOperationAudit(action string, resourceID string, beforeData interface{}, afterData interface{}, result string, errorMessage string) {
	actor := c.GetUser()
	cmd := core.OperationAuditCreateCommand{
		Action:       action,
		ResourceType: "session",
		ResourceID:   resourceID,
		BeforeData:   auditJSON(beforeData),
		AfterData:    auditJSON(afterData),
		IP:           c.Ctx.RemoteAddr(),
		UserAgent:    c.Ctx.GetHeader("User-Agent"),
		Result:       result,
		ErrorMessage: errorMessage,
	}
	if actor != nil {
		cmd.ActorUserID = actor.Id
		cmd.ActorUsername = actor.Username
	}
	_ = c.auditService().CreateOperationAudit(cmd)
}

func sanitizeSessionForAudit(session *model.AuthToken) iris.Map {
	if session == nil {
		return nil
	}
	return iris.Map{
		"id":          session.Id,
		"user_id":     session.UserId,
		"rustdesk_id": session.RustdeskId,
		"is_admin":    session.IsAdmin,
		"status":      session.Status,
		"expired":     session.Expired.Format(config.TimeFormat),
	}
}
