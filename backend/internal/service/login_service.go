package service

import (
	apiform "rustdesk-api-server-pro/app/form/api"
	legacyservice "rustdesk-api-server-pro/app/service"
)

// LoginService wraps the legacy login flow so controllers stay thin during migration.
// The underlying auth/token behavior remains unchanged.
type LoginService struct{}

func NewLoginService() *LoginService {
	return &LoginService{}
}

func (s *LoginService) HandleLogin(form apiform.LoginForm) any {
	userService := legacyservice.NewUserService()

	// email verification step
	if form.Type == "email_code" && form.VerificationCode != "" && form.TfaCode == "" {
		return userService.LoginVerifyByEmailCode(form)
	}

	// 2FA verification step
	if form.Type == "email_code" && form.TfaCode != "" && form.VerificationCode == form.TfaCode {
		return userService.LoginVerifyBy2FACode(form)
	}

	// regular account login
	return userService.Login(form)
}
