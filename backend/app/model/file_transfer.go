package model

import "time"

type FileTransfer struct {
	Id           int       `xorm:"'id' int notnull pk autoincr"`
	RustdeskId   string    `xorm:"'rustdesk_id' varchar(100) index"`
	Info         string    `xorm:"'info' text"`
	IsFile       bool      `xorm:"'is_file' tinyint"`
	Path         string    `xorm:"'path' text"`
	FileName     string    `xorm:"'file_name' varchar(512)"`
	PeerId       string    `xorm:"'peer_id' varchar(100) index"`
	SessionId    string    `xorm:"'session_id' varchar(100) index"`
	Type         int       `xorm:"'type' tinyint"`
	Uuid         string    `xorm:"'uuid' varchar(255) index"`
	Direction    string    `xorm:"'direction' varchar(32) index"`
	SizeBytes    int64     `xorm:"'size_bytes' bigint"`
	Result       string    `xorm:"'result' varchar(32) index"`
	ErrorMessage string    `xorm:"'error_message' varchar(1024)"`
	Raw          string    `xorm:"'raw' text"`
	CreatedAt    time.Time `xorm:"'created_at' datetime created index"`
	UpdatedAt    time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *FileTransfer) TableName() string {
	return "file_transfer"
}
