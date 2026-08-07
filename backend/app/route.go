package app

import (
	"rustdesk-api-server-pro/app/controller/admin"
	"rustdesk-api-server-pro/app/controller/api"
	"rustdesk-api-server-pro/app/controller/userportal"
	"rustdesk-api-server-pro/app/middleware"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func SetRoute(app *iris.Application) {
	apiParty := app.Party("/api")
	apiMvc := mvc.New(apiParty)
	apiMvc.Handle(new(api.SystemController))
	apiMvc.Handle(new(api.LoginController))
	apiMvc.Handle(new(api.AuditController))
	apiMvc.Handle(new(api.CompatPublicController))
	apiMvc.Handle(new(api.OAuthController))
	apiMvc.Handle(new(api.ErrCodeController))

	licWebApiParty := app.Party("/lic/web/api")
	licWebApiMvc := mvc.New(licWebApiParty)
	licWebApiMvc.Handle(new(api.CompatLicController))

	apiWithAuthParty := app.Party("/api")
	apiWithAuthParty.Use(middleware.ApiAuth(app))
	{
		apiWithAuthMvc := mvc.New(apiWithAuthParty)
		apiWithAuthMvc.Handle(new(api.UserController))
		apiWithAuthMvc.Handle(new(api.PeerController))
		apiWithAuthMvc.Handle(new(api.CompatAuthController))
		apiWithAuthMvc.Handle(new(api.DeviceGroupController))
		apiWithAuthMvc.Handle(new(api.EnterpriseCompatController))
		apiWithAuthMvc.Handle(new(api.AddressBookController))
		apiWithAuthMvc.Handle(new(api.AddressBookPeerController))
		apiWithAuthMvc.Handle(new(api.AddressBookTagController))
	}

	adminParty := app.Party("/admin")
	adminMvc := mvc.New(adminParty)
	adminMvc.Handle(new(admin.AuthController))

	adminWithAuthParty := app.Party("/admin")
	adminWithAuthParty.Use(middleware.AdminAuth(app))
	{
		adminWithAuthMvc := mvc.New(adminWithAuthParty)
		adminWithAuthMvc.Handle(new(admin.UsersController))
		adminWithAuthMvc.Handle(new(admin.SessionsController))
		adminWithAuthMvc.Handle(new(admin.DevicesController))
		adminWithAuthMvc.Handle(new(admin.AddressBookController))
	}

	adminOrUserAuthParty := app.Party("/admin")
	adminOrUserAuthParty.Use(middleware.AdminOrUserAuth(app))
	{
		adminOrUserAuthMvc := mvc.New(adminOrUserAuthParty)
		adminOrUserAuthMvc.Handle(new(admin.DashboardController))
		adminOrUserAuthMvc.Handle(new(admin.AuditController))
		adminOrUserAuthMvc.Handle(new(admin.MailTemplateController))
		adminOrUserAuthMvc.Handle(new(admin.MaiLogsController))
		adminOrUserAuthMvc.Handle(new(admin.TokenController))
		adminOrUserAuthMvc.Handle(new(admin.OAuthController))
		adminOrUserAuthMvc.Handle(new(admin.SecurityAuditController))
		adminOrUserAuthMvc.Handle(new(admin.ErrorLogController))
		adminOrUserAuthMvc.Handle(new(admin.ContainerLogController))
	}

	adminUserAuthParty := app.Party("/admin")
	adminUserAuthParty.Use(middleware.UserAuth(app))
	{
		adminUserAuthMvc := mvc.New(adminUserAuthParty)
		adminUserAuthMvc.Handle(new(admin.IndexController))
	}

	userPortalWithAuthParty := app.Party("/user-portal")
	userPortalWithAuthParty.Use(middleware.UserAuth(app))
	{
		userPortalMvc := mvc.New(userPortalWithAuthParty)
		userPortalMvc.Handle(new(userportal.IndexController))
		userPortalMvc.Handle(new(userportal.DevicesController))
		userPortalMvc.Handle(new(userportal.AccountController))
	}
}
