package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/util"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"xorm.io/xorm"
)

type WebauthnService struct {
	cfg    *config.ServerConfig
	db     *xorm.Engine
	wauthn *webauthn.WebAuthn
}

func NewWebauthnService(cfg *config.ServerConfig, db *xorm.Engine) *WebauthnService {
	s := &WebauthnService{cfg: cfg, db: db}
	s.initFromConfig()
	return s
}

func (s *WebauthnService) initFromConfig() {
	if s.cfg == nil || s.cfg.WebAuthn == nil || !s.cfg.WebAuthn.Enabled {
		return
	}
	rpID := strings.TrimSpace(s.cfg.WebAuthn.RPID)
	rpName := strings.TrimSpace(s.cfg.WebAuthn.RPName)
	if rpName == "" {
		rpName = "RustDesk API Server Pro"
	}
	origins := s.cfg.WebAuthn.RPOrigins
	if rpID != "" && len(origins) > 0 {
		wconfig := &webauthn.Config{
			RPDisplayName: rpName,
			RPID:          rpID,
			RPOrigins:     origins,
		}
		if w, err := webauthn.New(wconfig); err == nil {
			s.wauthn = w
		}
	}
}

func (s *WebauthnService) IsEnabled() bool {
	return s != nil && s.wauthn != nil
}

func (s *WebauthnService) EnsureEnabled() error {
	if !s.IsEnabled() {
		if s.cfg == nil || s.cfg.WebAuthn == nil || !s.cfg.WebAuthn.Enabled {
			return errcode.New(errcode.ERR3101.Code, errcode.ERR3101.Message)
		}
		return errcode.New(errcode.ERR3102.Code, errcode.ERR3102.Message)
	}
	return nil
}

type webauthnUserAdapter struct {
	user        *model.User
	credentials []webauthn.Credential
}

func (u *webauthnUserAdapter) WebAuthnID() []byte {
	return []byte(strconv.Itoa(u.user.Id))
}

func (u *webauthnUserAdapter) WebAuthnName() string {
	return u.user.Username
}

func (u *webauthnUserAdapter) WebAuthnDisplayName() string {
	if u.user.Name != "" {
		return u.user.Name
	}
	return u.user.Username
}

func (u *webauthnUserAdapter) WebAuthnCredentials() []webauthn.Credential {
	return u.credentials
}

func (u *webauthnUserAdapter) WebAuthnIcon() string {
	return ""
}

func (s *WebauthnService) loadCredentials(userID int) ([]webauthn.Credential, error) {
	var creds []model.WebauthnCredential
	if err := s.db.Where("user_id = ?", userID).Find(&creds); err != nil {
		return nil, err
	}
	result := make([]webauthn.Credential, 0, len(creds))
	for _, c := range creds {
		credID, err := base64.RawURLEncoding.DecodeString(c.CredentialId)
		if err != nil {
			continue
		}
		var transports []protocol.AuthenticatorTransport
		if c.Transport != "" {
			for _, t := range strings.Split(c.Transport, ",") {
				t = strings.TrimSpace(t)
				if t != "" {
					transports = append(transports, protocol.AuthenticatorTransport(t))
				}
			}
		}
		var aaguid []byte
		if c.AAGUID != "" {
			aaguid, _ = base64.RawURLEncoding.DecodeString(c.AAGUID)
		}
		result = append(result, webauthn.Credential{
			ID:              credID,
			PublicKey:       c.PublicKey,
			AttestationType: c.AttestationType,
			Transport:       transports,
			Authenticator: webauthn.Authenticator{
				SignCount: c.SignCount,
				AAGUID:    aaguid,
			},
		})
	}
	return result, nil
}

func (s *WebauthnService) buildUserAdapter(user *model.User) (*webauthnUserAdapter, error) {
	creds, err := s.loadCredentials(user.Id)
	if err != nil {
		return nil, err
	}
	return &webauthnUserAdapter{user: user, credentials: creds}, nil
}

// BeginRegistration 生成注册 challenge
func (s *WebauthnService) BeginRegistration(user *model.User) (*protocol.CredentialCreation, error) {
	if err := s.EnsureEnabled(); err != nil {
		return nil, err
	}
	adapter, err := s.buildUserAdapter(user)
	if err != nil {
		return nil, err
	}
	options, session, err := s.wauthn.BeginRegistration(adapter)
	if err != nil {
		return nil, err
	}
	if err = s.saveSession("register", user.Id, session); err != nil {
		return nil, err
	}
	return options, nil
}

// FinishRegistrationParsed 验证注册响应并存储凭据
func (s *WebauthnService) FinishRegistrationParsed(user *model.User, parsed *protocol.ParsedCredentialCreationData, credentialName string) (*model.WebauthnCredential, error) {
	if err := s.EnsureEnabled(); err != nil {
		return nil, err
	}
	session, err := s.popSession("register", user.Id)
	if err != nil {
		return nil, err
	}
	adapter, err := s.buildUserAdapter(user)
	if err != nil {
		return nil, err
	}
	credential, err := s.wauthn.CreateCredential(adapter, *session, parsed)
	if err != nil {
		return nil, errcode.New(errcode.ERR3105.Code, errcode.ERR3105.Message)
	}

	credID := base64.RawURLEncoding.EncodeToString(credential.ID)
	var existing model.WebauthnCredential
	has, _ := s.db.Where("credential_id = ?", credID).Get(&existing)
	if has {
		return nil, errcode.New(errcode.ERR3109.Code, errcode.ERR3109.Message)
	}

	transports := make([]string, 0, len(credential.Transport))
	for _, t := range credential.Transport {
		transports = append(transports, string(t))
	}
	aaguidStr := ""
	if len(credential.Authenticator.AAGUID) > 0 {
		aaguidStr = base64.RawURLEncoding.EncodeToString(credential.Authenticator.AAGUID)
	}
	record := &model.WebauthnCredential{
		UserId:          user.Id,
		CredentialId:    credID,
		PublicKey:       credential.PublicKey,
		AttestationType: credential.AttestationType,
		Transport:       strings.Join(transports, ","),
		SignCount:       credential.Authenticator.SignCount,
		AAGUID:          aaguidStr,
		Name:            credentialName,
	}
	if _, err = s.db.Insert(record); err != nil {
		return nil, err
	}
	return record, nil
}

// BeginLogin 生成登录 challenge
func (s *WebauthnService) BeginLogin(username string) (*protocol.CredentialAssertion, int, error) {
	if err := s.EnsureEnabled(); err != nil {
		return nil, 0, err
	}
	var user model.User
	has, err := s.db.Where("username = ? and status > 0", username).Get(&user)
	if err != nil {
		return nil, 0, err
	}
	if !has {
		return nil, 0, errcode.New(errcode.ERR3107.Code, errcode.ERR3107.Message)
	}
	adapter, err := s.buildUserAdapter(&user)
	if err != nil {
		return nil, 0, err
	}
	if len(adapter.credentials) == 0 {
		return nil, 0, errcode.New(errcode.ERR3108.Code, errcode.ERR3108.Message)
	}
	options, session, err := s.wauthn.BeginLogin(adapter)
	if err != nil {
		return nil, 0, err
	}
	if err = s.saveSession("login", user.Id, session); err != nil {
		return nil, 0, err
	}
	return options, user.Id, nil
}

// FinishLoginParsed 验证登录响应，返回用户
func (s *WebauthnService) FinishLoginParsed(parsed *protocol.ParsedCredentialAssertionData) (*model.User, error) {
	if err := s.EnsureEnabled(); err != nil {
		return nil, err
	}
	session, sessionKey, err := s.peekLoginSession()
	if err != nil {
		return nil, err
	}
	if session.UserId == 0 {
		return nil, errcode.New(errcode.ERR3104.Code, errcode.ERR3104.Message)
	}
	var user model.User
	has, err := s.db.Where("id = ? and status > 0", session.UserId).Get(&user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errcode.New(errcode.ERR3107.Code, errcode.ERR3107.Message)
	}
	adapter, err := s.buildUserAdapter(&user)
	if err != nil {
		return nil, err
	}
	credential, err := s.wauthn.ValidateLogin(adapter, *session.SessionData, parsed)
	if err != nil {
		return nil, errcode.New(errcode.ERR3106.Code, errcode.ERR3106.Message)
	}

	credID := base64.RawURLEncoding.EncodeToString(credential.ID)
	_, _ = s.db.Where("user_id = ? and credential_id = ?", user.Id, credID).Cols("sign_count", "updated_at").Update(&model.WebauthnCredential{
		SignCount: credential.Authenticator.SignCount,
	})

	s.consumeSession(sessionKey)
	return &user, nil
}

// ListCredentials 列出用户已绑定的 Passkey
func (s *WebauthnService) ListCredentials(userID int) ([]model.WebauthnCredential, error) {
	var creds []model.WebauthnCredential
	err := s.db.Where("user_id = ?", userID).Desc("created_at").Find(&creds)
	return creds, err
}

// DeleteCredential 删除指定 Passkey
func (s *WebauthnService) DeleteCredential(userID int, id int) error {
	var cred model.WebauthnCredential
	has, err := s.db.Where("user_id = ? and id = ?", userID, id).Get(&cred)
	if err != nil {
		return err
	}
	if !has {
		return errcode.New(errcode.ERR3110.Code, errcode.ERR3110.Message)
	}
	_, err = s.db.Where("user_id = ? and id = ?", userID, id).Delete(&model.WebauthnCredential{})
	return err
}

// RenameCredential 重命名 Passkey
func (s *WebauthnService) RenameCredential(userID int, id int, name string) error {
	var cred model.WebauthnCredential
	has, err := s.db.Where("user_id = ? and id = ?", userID, id).Get(&cred)
	if err != nil {
		return err
	}
	if !has {
		return errcode.New(errcode.ERR3110.Code, errcode.ERR3110.Message)
	}
	_, err = s.db.Where("user_id = ? and id = ?", userID, id).Cols("name").Update(&model.WebauthnCredential{Name: name})
	return err
}

func (s *WebauthnService) saveSession(kind string, userID int, sessionData *webauthn.SessionData) error {
	challengeBytes, _ := json.Marshal(sessionData)
	record := &model.WebauthnSession{
		KeyHash:   util.Sha256Hex(util.RandomString(32)),
		Kind:      kind,
		UserId:    userID,
		Challenge: challengeBytes,
		ExpiresAt: time.Now().Add(5 * time.Minute),
		Status:    1,
	}
	if _, err := s.db.Insert(record); err != nil {
		return err
	}
	return nil
}

func (s *WebauthnService) popSession(kind string, userID int) (*webauthn.SessionData, error) {
	var session model.WebauthnSession
	has, err := s.db.Where("kind = ? and user_id = ? and status = 1 and expires_at > ?", kind, userID, time.Now()).Desc("created_at").Get(&session)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errcode.New(errcode.ERR3103.Code, errcode.ERR3103.Message)
	}
	_, _ = s.db.Where("id = ?", session.Id).Cols("status").Update(&model.WebauthnSession{Status: 0})
	var sd webauthn.SessionData
	if err = json.Unmarshal(session.Challenge, &sd); err != nil {
		return nil, err
	}
	return &sd, nil
}

type loginSession struct {
	*webauthn.SessionData
	UserId int
}

func (s *WebauthnService) peekLoginSession() (*loginSession, string, error) {
	var session model.WebauthnSession
	has, err := s.db.Where("kind = ? and status = 1 and expires_at > ?", "login", time.Now()).Desc("created_at").Get(&session)
	if err != nil {
		return nil, "", err
	}
	if !has {
		return nil, "", errcode.New(errcode.ERR3103.Code, errcode.ERR3103.Message)
	}
	var sd webauthn.SessionData
	if err = json.Unmarshal(session.Challenge, &sd); err != nil {
		return nil, "", err
	}
	return &loginSession{SessionData: &sd, UserId: session.UserId}, session.KeyHash, nil
}

func (s *WebauthnService) consumeSession(keyHash string) {
	_, _ = s.db.Where("key_hash = ?", keyHash).Cols("status").Update(&model.WebauthnSession{Status: 0})
}

// ParseCredentialCreationBody 从 io.Reader 解析注册响应
func ParseCredentialCreationBody(r io.Reader) (*protocol.ParsedCredentialCreationData, error) {
	return protocol.ParseCredentialCreationResponseBody(r)
}

// ParseCredentialRequestBody 从 io.Reader 解析登录响应
func ParseCredentialRequestBody(r io.Reader) (*protocol.ParsedCredentialAssertionData, error) {
	return protocol.ParseCredentialRequestResponseBody(r)
}

// UpdateConfig 动态更新 WebAuthn 配置（RPID/RPOrigins 从请求推导）
func (s *WebauthnService) UpdateConfig(rpID string, origins []string) error {
	if s.cfg == nil {
		return fmt.Errorf("config is nil")
	}
	if s.cfg.WebAuthn == nil {
		s.cfg.WebAuthn = &config.WebAuthnConfig{}
	}
	s.cfg.WebAuthn.RPID = rpID
	s.cfg.WebAuthn.RPOrigins = origins
	s.cfg.WebAuthn.Enabled = true

	rpName := strings.TrimSpace(s.cfg.WebAuthn.RPName)
	if rpName == "" {
		rpName = "RustDesk API Server Pro"
	}
	if rpID == "" || len(origins) == 0 {
		return errcode.New(errcode.ERR3102.Code, errcode.ERR3102.Message)
	}
	wconfig := &webauthn.Config{
		RPDisplayName: rpName,
		RPID:          rpID,
		RPOrigins:     origins,
	}
	w, err := webauthn.New(wconfig)
	if err != nil {
		return fmt.Errorf("webauthn init failed: %w", err)
	}
	s.wauthn = w
	return nil
}
