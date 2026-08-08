package model

import (
	"time"
)

// WebauthnCredential 存储用户绑定的 WebAuthn 凭据（Passkey）
type WebauthnCredential struct {
	Id              int       `xorm:"'id' int notnull pk autoincr"`
	UserId          int       `xorm:"'user_id' int notnull index"`
	CredentialId    string    `xorm:"'credential_id' varchar(512) notnull unique"`
	PublicKey       []byte    `xorm:"'public_key' blob"`
	AttestationType string    `xorm:"'attestation_type' varchar(64)"`
	Transport       string    `xorm:"'transport' varchar(256)"`
	SignCount       uint32    `xorm:"'sign_count' int"`
	AAGUID          string    `xorm:"'aaguid' varchar(64)"`
	Name            string    `xorm:"'name' varchar(100)"`
	CreatedAt       time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt       time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *WebauthnCredential) TableName() string {
	return "webauthn_credential"
}

// WebauthnSession 存储 WebAuthn 注册/登录过程中的临时 session（challenge）
type WebauthnSession struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	KeyHash   string    `xorm:"'key_hash' varchar(64) notnull unique"`
	Kind      string    `xorm:"'kind' varchar(20) notnull"`
	UserId    int       `xorm:"'user_id' int"`
	Challenge []byte    `xorm:"'challenge' blob"`
	ExpiresAt time.Time `xorm:"'expires_at' datetime"`
	Status    int       `xorm:"'status' tinyint"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
}

func (m *WebauthnSession) TableName() string {
	return "webauthn_session"
}
