package model

import "time"

type ContainerLog struct {
	Id          int       `xorm:"'id' int notnull pk autoincr"`
	Timestamp   time.Time `xorm:"'timestamp' datetime index"`
	Level       string    `xorm:"'level' varchar(20) index"`
	Source      string    `xorm:"'source' varchar(50) index"`
	Message     string    `xorm:"'message' text"`
	Method      string    `xorm:"'method' varchar(10)"`
	Path        string    `xorm:"'path' varchar(255) index"`
	StatusCode  int       `xorm:"'status_code' int"`
	DurationMs  int64     `xorm:"'duration_ms' bigint"`
	ClientIP    string    `xorm:"'client_ip' varchar(50)"`
	UserId      int       `xorm:"'user_id' int index"`
	UserName    string    `xorm:"'user_name' varchar(50)"`
	UserAgent   string    `xorm:"'user_agent' varchar(500)"`
	RequestId   string    `xorm:"'request_id' varchar(64) index"`
	ErrorMsg    string    `xorm:"'error_msg' varchar(1024)"`
	CreatedAt   time.Time `xorm:"'created_at' datetime created"`
}

func (m *ContainerLog) TableName() string {
	return "container_log"
}
