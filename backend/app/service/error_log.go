package service

import (
	"rustdesk-api-server-pro/app/model"

	"xorm.io/xorm"
)

func RecordErrorLog(db *xorm.Engine, code string, message string, module string, path string, method string, userId int, userName string, clientIP string, userAgent string) {
	_, _ = db.InsertOne(&model.ErrorLog{
		Code:      code,
		Message:   message,
		Module:    module,
		Path:      path,
		Method:    method,
		UserId:    userId,
		UserName:  userName,
		ClientIP:  clientIP,
		UserAgent: userAgent,
	})
}
