package admin

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/db"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type OAuthController struct {
	basicController
}

func (c *OAuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/oauth/accounts", "HandleListAccounts")
	b.Handle("DELETE", "/oauth/account/{id:int}", "HandleDeleteAccount")
	b.Handle("GET", "/oauth/providers", "HandleListProviders")
}

func (c *OAuthController) HandleListAccounts() mvc.Result {
	currentPage := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)

	query := func() *xorm.Session {
		return c.Db.Table(&model.OAuthAccount{}).Desc("id")
	}

	pagination := db.NewPagination(currentPage, pageSize)
	accountList := make([]model.OAuthAccount, 0)
	if err := pagination.Paginate(query, &model.OAuthAccount{}, &accountList); err != nil {
		return c.Error(nil, err.Error())
	}

	list := make([]iris.Map, 0, len(accountList))
	for _, a := range accountList {
		list = append(list, iris.Map{
			"id":           a.Id,
			"user_id":      a.UserId,
			"provider":     a.Provider,
			"subject":      a.Subject,
			"email":        a.Email,
			"name":         a.Name,
			"is_admin":     a.IsAdmin,
			"status":       a.Status,
			"last_login_at": a.LastLoginAt.Format(config.TimeFormat),
			"created_at":   a.CreatedAt.Format(config.TimeFormat),
		})
	}

	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": list,
		"current": currentPage,
		"size":    pageSize,
	}, "ok")
}

func (c *OAuthController) HandleDeleteAccount() mvc.Result {
	id := c.Ctx.Params().GetIntDefault("id", 0)
	if id == 0 {
		return c.Error(nil, "InvalidAccountId")
	}

	_, err := c.Db.ID(id).Delete(&model.OAuthAccount{})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(nil, "ok")
}

func (c *OAuthController) HandleListProviders() mvc.Result {
	cfg := config.GetServerConfig()
	providers := cfg.OAuthProviders()

	list := make([]iris.Map, 0, len(providers))
	for _, p := range providers {
		list = append(list, iris.Map{
			"type":        p.Type,
			"name":        p.Name,
			"displayName": p.DisplayName,
			"enabled":     p.Enabled,
		})
	}

	return c.Success(list, "ok")
}
