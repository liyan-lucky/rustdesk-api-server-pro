package model

import (
	"rustdesk-api-server-pro/util"
	"time"
)

type AuthToken struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	UserId     int       `xorm:"'user_id' int"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(255)"`
	Uuid       string    `xorm:"'uuid' varchar(255)"`
	DeviceOs   string    `xorm:"'device_os' varchar(255)"`
	DeviceType string    `xorm:"'device_type' varchar(50)"`
	DeviceName string    `xorm:"'device_name' varchar(255)"`
	Token      string    `xorm:"'token' varchar(255)"`
	TokenHash  string    `xorm:"'token_hash' varchar(64) index"`
	Expired    time.Time `xorm:"'expired' datetime"`
	IsAdmin    bool      `xorm:"'is_admin' tinyint"`
	Status     int       `xorm:"'status' tinyint"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt  time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *AuthToken) TableName() string {
	return "auth_token"
}

func (m *AuthToken) BeforeInsert() {
	m.normalizeTokenHash()
}

func (m *AuthToken) BeforeUpdate() {
	m.normalizeTokenHash()
}

func (m *AuthToken) normalizeTokenHash() {
	if m == nil || m.Token == "" {
		return
	}
	if m.TokenHash == "" {
		m.TokenHash = util.Sha256Hex(m.Token)
	}
	m.Token = ""
}
