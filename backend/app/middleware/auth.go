package middleware

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/helper"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/util"
	"time"

	"github.com/golang-module/carbon/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"xorm.io/xorm"
)

func ApiAuth(app *iris.Application) iris.Handler {
	return func(context iris.Context) {
		db := helper.GetAppDependency(app, "*xorm.Engine").(*xorm.Engine)
		token := extractToken(context)

		authToken, get, err := getActiveAuthToken(db, token, false)
		if !get || err != nil {
			recordAuthSecurityAudit(db, context, "api_token_invalid", 0, "", false, authFailureReason(err, "Unauthorized"))
			context.StopWithJSON(iris.StatusUnauthorized, iris.Map{"code": 401, "data": nil, "message": errcode.ErrUnauthorized.Error()})
			return
		}

		var user model.User
		get, err = db.Where("id = ? and status > 0", authToken.UserId).Get(&user)
		if !get || err != nil {
			recordAuthSecurityAudit(db, context, "api_token_user_invalid", authToken.UserId, "", false, authFailureReason(err, "NotAcceptable"))
			context.StopWithJSON(iris.StatusNotAcceptable, iris.Map{"code": 406, "data": nil, "message": errcode.ErrUnauthorized.Error()})
			return
		}

		context.Values().Set(config.CurrentUserKey, &user)
		context.Values().Set(config.CurrentAuthTokenString, token)
		context.Values().Set(config.CurrentAuthToken, &authToken)
		context.Next()
	}
}

func extractToken(context iris.Context) string {
	token := jwt.FromHeader(context)
	if token == "" {
		token = context.GetHeader("Authorization")
	}
	return token
}

func AdminAuth(app *iris.Application) iris.Handler {
	return func(context iris.Context) {
		db := helper.GetAppDependency(app, "*xorm.Engine").(*xorm.Engine)
		token := extractToken(context)

		authToken, get, err := getActiveAuthToken(db, token, true)
		if !get || err != nil {
			recordAuthSecurityAudit(db, context, "admin_token_invalid", 0, "", false, authFailureReason(err, "Unauthorized"))
			context.StopWithJSON(iris.StatusUnauthorized, iris.Map{"code": 401, "data": nil, "message": errcode.ErrUnauthorized.Error()})
			return
		}

		var user model.User
		get, err = db.Where("id = ? and status > 0", authToken.UserId).Get(&user)
		if !get || err != nil {
			recordAuthSecurityAudit(db, context, "admin_token_user_invalid", authToken.UserId, "", false, authFailureReason(err, "NotAcceptable"))
			context.StopWithJSON(iris.StatusNotAcceptable, iris.Map{"code": 406, "data": nil, "message": errcode.ErrUnauthorized.Error()})
			return
		}

		s := carbon.Now().DiffInSeconds(carbon.Parse(authToken.Expired.Format(config.TimeFormat)))
		if s > 0 && s <= 60*5 {
			authToken.Expired = carbon.Parse(authToken.Expired.Format(config.TimeFormat)).AddHours(2).ToStdTime()
			_, _ = db.Where("id = ?", authToken.Id).Cols("expired").Update(&authToken)
		}

		context.Values().Set(config.AdminUserKey, &user)
		context.Values().Set(config.AdminAuthTokenString, token)
		context.Values().Set(config.AdminAuthToken, &authToken)
		context.Next()
	}
}

func getActiveAuthToken(db *xorm.Engine, token string, isAdmin bool) (model.AuthToken, bool, error) {
	var authToken model.AuthToken
	if token == "" {
		return authToken, false, nil
	}

	now := time.Now().Format(config.TimeFormat)
	tokenHash := util.Sha256Hex(token)
	get, err := db.Where("token_hash = ? and expired > ? and status = 1 and is_admin = ?", tokenHash, now, isAdmin).Get(&authToken)
	if get || err != nil {
		return authToken, get, err
	}

	// Backward compatibility for tokens issued before token_hash existed.
	get, err = db.Where("token = ? and expired > ? and status = 1 and is_admin = ?", token, now, isAdmin).Get(&authToken)
	if get && authToken.TokenHash == "" {
		authToken.TokenHash = tokenHash
		authToken.Token = ""
		_, _ = db.Where("id = ?", authToken.Id).Cols("token_hash", "token").Update(&authToken)
	}
	return authToken, get, err
}

func UserAuth(app *iris.Application) iris.Handler {
	return func(context iris.Context) {
		db := helper.GetAppDependency(app, "*xorm.Engine").(*xorm.Engine)
		token := extractToken(context)

		authToken, get, err := getActiveAuthToken(db, token, false)
		if !get || err != nil {
			authToken, get, err = getActiveAuthToken(db, token, true)
		}
		if !get || err != nil {
			recordAuthSecurityAudit(db, context, "web_user_token_invalid", 0, "", false, authFailureReason(err, "Unauthorized"))
			context.StopWithJSON(iris.StatusUnauthorized, iris.Map{"code": 401, "data": nil, "message": errcode.ErrUnauthorized.Error()})
			return
		}

		var user model.User
		get, err = db.Where("id = ? and status > 0", authToken.UserId).Get(&user)
		if !get || err != nil {
			recordAuthSecurityAudit(db, context, "web_user_invalid", authToken.UserId, "", false, authFailureReason(err, "NotAcceptable"))
			context.StopWithJSON(iris.StatusNotAcceptable, iris.Map{"code": 406, "data": nil, "message": errcode.ErrUnauthorized.Error()})
			return
		}

		s := carbon.Now().DiffInSeconds(carbon.Parse(authToken.Expired.Format(config.TimeFormat)))
		if s > 0 && s <= 60*5 {
			authToken.Expired = carbon.Parse(authToken.Expired.Format(config.TimeFormat)).AddHours(2).ToStdTime()
			_, _ = db.Where("id = ?", authToken.Id).Cols("expired").Update(&authToken)
		}

		context.Values().Set(config.WebUserKey, &user)
		context.Values().Set(config.WebAuthTokenString, token)
		context.Values().Set(config.WebAuthToken, &authToken)
		context.Next()
	}
}

func recordAuthSecurityAudit(db *xorm.Engine, context iris.Context, event string, userID int, username string, success bool, reason string) {
	if db == nil || context == nil {
		return
	}
	_, _ = db.Insert(&model.SecurityAudit{
		UserId:    userID,
		Username:  username,
		Event:     event,
		IP:        context.RemoteAddr(),
		UserAgent: context.GetHeader("User-Agent"),
		Success:   success,
		Reason:    reason,
	})
}

func authFailureReason(err error, fallback string) string {
	if err != nil {
		return err.Error()
	}
	return fallback
}
