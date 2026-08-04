package model

import "time"

// OAuthLoginSession stores short-lived, one-time OAuth state and login tickets.
// Only hashes of browser-visible state/ticket values are persisted.
type OAuthLoginSession struct {
	Id           int       `xorm:"'id' int notnull pk autoincr"`
	Kind         string    `xorm:"'kind' varchar(20) index notnull"`
	KeyHash      string    `xorm:"'key_hash' varchar(64) unique notnull"`
	Provider     string    `xorm:"'provider' varchar(50) index"`
	RedirectTo   string    `xorm:"'redirect_to' varchar(1024)"`
	CallbackURL  string    `xorm:"'callback_url' varchar(1024)"`
	CodeVerifier string    `xorm:"'code_verifier' varchar(255)"`
	UserId       int       `xorm:"'user_id' int index"`
	IsAdmin      bool      `xorm:"'is_admin' tinyint"`
	RustdeskId   string    `xorm:"'rustdesk_id' varchar(100)"`
	Uuid         string    `xorm:"'uuid' varchar(100)"`
	DeviceOs     string    `xorm:"'device_os' varchar(100)"`
	DeviceType   string    `xorm:"'device_type' varchar(100)"`
	DeviceName   string    `xorm:"'device_name' varchar(255)"`
	PollToken    string    `xorm:"'poll_token' varchar(64) index"`
	Ticket       string    `xorm:"'ticket' varchar(64)"`
	ExpiresAt    time.Time `xorm:"'expires_at' datetime index"`
	Status       int       `xorm:"'status' tinyint index"`
	CreatedAt    time.Time `xorm:"'created_at' datetime created"`
}

func (m *OAuthLoginSession) TableName() string { return "oauth_login_session" }
