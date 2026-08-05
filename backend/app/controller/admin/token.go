package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/errcode"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type TokenController struct {
	basicController
}

func (c *TokenController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tokens/list", "HandleList")
	b.Handle("POST", "/tokens/kill", "HandleKill")
	b.Handle("POST", "/tokens/clear", "HandleClear")
}

func (c *TokenController) HandleList() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	username := c.Ctx.URLParamDefault("username", "")

	query := func() *xorm.Session {
		q := c.Db.Table(&model.AuthToken{})
		q.Join("INNER", &model.User{}, "auth_token.user_id = user.id")
		if username != "" {
			q.Where("user.username LIKE ?", "%"+username+"%")
		}
		q.Desc("auth_token.id")
		return q
	}

	type TokenRow struct {
		model.AuthToken `xorm:"extends"`
		Username        string `xorm:"user.username"`
	}

	pagination := db.NewPagination(currentPage, pageSize)
	tokenList := make([]TokenRow, 0)
	if err := pagination.Paginate(query, &TokenRow{}, &tokenList); err != nil {
		return c.dbError(err)
	}

	list := make([]iris.Map, 0, len(tokenList))
	for _, t := range tokenList {
		list = append(list, iris.Map{
			"id":          t.AuthToken.Id,
			"user_id":     t.AuthToken.UserId,
			"username":    t.Username,
			"rustdesk_id": t.AuthToken.RustdeskId,
			"uuid":        t.AuthToken.Uuid,
			"device_os":   t.AuthToken.DeviceOs,
			"device_type": t.AuthToken.DeviceType,
			"device_name": t.AuthToken.DeviceName,
			"token_hash":  t.AuthToken.TokenHash,
			"is_admin":    t.AuthToken.IsAdmin,
			"status":      t.AuthToken.Status,
			"expired":     t.AuthToken.Expired.Format(config.TimeFormat),
			"created_at":  t.AuthToken.CreatedAt.Format(config.TimeFormat),
		})
	}

	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *TokenController) HandleKill() mvc.Result {
	var params struct {
		Ids []int `json:"ids"`
	}
	if err := c.Ctx.ReadJSON(&params); err != nil {
		return c.dbError(err)
	}
	if len(params.Ids) == 0 {
		return c.Error(nil, errcode.New(errcode.ERR9002.Code, errcode.ERR9002.Message).Error())
	}

	_, err := c.Db.In("id", params.Ids).Cols("status").Update(&model.AuthToken{Status: 0})
	if err != nil {
		return c.dbError(err)
	}

	return c.Success(nil, "ok")
}

func (c *TokenController) HandleClear() mvc.Result {
	_, err := c.Db.Cols("status").Update(&model.AuthToken{Status: 0})
	if err != nil {
		return c.dbError(err)
	}

	return c.Success(nil, "ok")
}
