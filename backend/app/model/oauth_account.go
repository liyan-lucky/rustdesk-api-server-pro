package model

import "time"

type OAuthAccount struct {
	Id          int       `xorm:"'id' int notnull pk autoincr"`
	UserId      int       `xorm:"'user_id' int index notnull"`
	Provider    string    `xorm:"'provider' varchar(50) index notnull"`
	Subject     string    `xorm:"'subject' varchar(255) index notnull"`
	Email       string    `xorm:"'email' varchar(255)"`
	Name        string    `xorm:"'name' varchar(255)"`
	Picture     string    `xorm:"'picture' varchar(1024)"`
	IsAdmin     bool      `xorm:"'is_admin' tinyint index"`
	Status      int       `xorm:"'status' tinyint"`
	LastLoginAt time.Time `xorm:"'last_login_at' datetime"`
	CreatedAt   time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt   time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *OAuthAccount) TableName() string {
	return "oauth_account"
}
