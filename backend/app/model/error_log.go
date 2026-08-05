package model

import "time"

type ErrorLog struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Code      string    `xorm:"'code' varchar(20) index"`
	Message   string    `xorm:"'message' varchar(500)"`
	Module    string    `xorm:"'module' varchar(50) index"`
	Path      string    `xorm:"'path' varchar(255)"`
	Method    string    `xorm:"'method' varchar(10)"`
	UserId    int       `xorm:"'user_id' int index"`
	UserName  string    `xorm:"'user_name' varchar(50)"`
	ClientIP  string    `xorm:"'client_ip' varchar(50)"`
	UserAgent string    `xorm:"'user_agent' varchar(500)"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
}

func (m *ErrorLog) TableName() string {
	return "error_log"
}
