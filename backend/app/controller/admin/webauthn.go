package admin

import (
	"bytes"
	"encoding/json"
	"io"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"
	v2service "rustdesk-api-server-pro/internal/service"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type WebauthnController struct {
	basicController
}

func (c *WebauthnController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/webauthn/register/begin", "PostRegisterBegin")
	b.Handle("POST", "/webauthn/register/finish", "PostRegisterFinish")
	b.Handle("GET", "/webauthn/credentials", "GetCredentials")
	b.Handle("DELETE", "/webauthn/credentials/{id:int}", "DeleteCredentials")
	b.Handle("PUT", "/webauthn/credentials/{id:int}", "PutCredentials")
}

func (c *WebauthnController) ensureWebauthnService() *v2service.WebauthnService {
	service := v2service.NewWebauthnService(config.GetServerConfig(), c.Db)
	if !service.IsEnabled() {
		rpID := strings.TrimSpace(c.Ctx.Host())
		if i := strings.IndexByte(rpID, ':'); i > 0 {
			rpID = rpID[:i]
		}
		if rpID != "" {
			scheme := "https"
			if c.Ctx.Request().TLS == nil {
				scheme = "http"
			}
			if forwardedProto := strings.TrimSpace(c.Ctx.GetHeader("X-Forwarded-Proto")); forwardedProto == "https" {
				scheme = "https"
			}
			origin := scheme + "://" + c.Ctx.Host()
			_ = service.UpdateConfig(rpID, []string{origin})
		}
	}
	return service
}

func (c *WebauthnController) PostRegisterBegin() mvc.Result {
	currentUser := c.GetUser()
	if currentUser == nil {
		return c.Error(nil, errcode.New(errcode.ERR1005.Code, errcode.ERR1005.Message).Error())
	}

	service := c.ensureWebauthnService()
	if err := service.EnsureEnabled(); err != nil {
		return c.Error(nil, err.Error())
	}

	options, err := service.BeginRegistration(currentUser)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(options, "ok")
}

func (c *WebauthnController) PostRegisterFinish() mvc.Result {
	currentUser := c.GetUser()
	if currentUser == nil {
		return c.Error(nil, errcode.New(errcode.ERR1005.Code, errcode.ERR1005.Message).Error())
	}

	service := c.ensureWebauthnService()
	if err := service.EnsureEnabled(); err != nil {
		return c.Error(nil, err.Error())
	}

	bodyBytes, err := io.ReadAll(c.Ctx.Request().Body)
	if err != nil {
		return c.Error(nil, errcode.New(errcode.ERR3105.Code, errcode.ERR3105.Message).Error())
	}

	parsed, err := v2service.ParseCredentialCreationBody(bytes.NewReader(bodyBytes))
	if err != nil {
		return c.Error(nil, errcode.New(errcode.ERR3105.Code, errcode.ERR3105.Message).Error())
	}

	var form struct {
		Name string `json:"name"`
	}
	_ = json.Unmarshal(bodyBytes, &form)
	credentialName := strings.TrimSpace(form.Name)
	if credentialName == "" {
		credentialName = "Passkey"
	}

	cred, err := service.FinishRegistrationParsed(currentUser, parsed, credentialName)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(iris.Map{
		"id":   cred.Id,
		"name": cred.Name,
	}, "ok")
}

func (c *WebauthnController) GetCredentials() mvc.Result {
	currentUser := c.GetUser()
	if currentUser == nil {
		return c.Error(nil, errcode.New(errcode.ERR1005.Code, errcode.ERR1005.Message).Error())
	}

	service := v2service.NewWebauthnService(config.GetServerConfig(), c.Db)
	creds, err := service.ListCredentials(currentUser.Id)
	if err != nil {
		return c.dbError(err)
	}

	type credentialItem struct {
		Id        int    `json:"id"`
		Name      string `json:"name"`
		AAGUID    string `json:"aaguid"`
		CreatedAt string `json:"createdAt"`
	}
	items := make([]credentialItem, 0, len(creds))
	for _, cred := range creds {
		items = append(items, credentialItem{
			Id:        cred.Id,
			Name:      cred.Name,
			AAGUID:    cred.AAGUID,
			CreatedAt: cred.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return c.Success(items, "ok")
}

func (c *WebauthnController) DeleteCredentials() mvc.Result {
	currentUser := c.GetUser()
	if currentUser == nil {
		return c.Error(nil, errcode.New(errcode.ERR1005.Code, errcode.ERR1005.Message).Error())
	}

	credentialID := c.Ctx.Params().GetIntDefault("id", 0)
	if credentialID == 0 {
		return c.Error(nil, errcode.New(errcode.ERR3110.Code, errcode.ERR3110.Message).Error())
	}

	service := v2service.NewWebauthnService(config.GetServerConfig(), c.Db)
	if err := service.DeleteCredential(currentUser.Id, credentialID); err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(nil, "ok")
}

func (c *WebauthnController) PutCredentials() mvc.Result {
	currentUser := c.GetUser()
	if currentUser == nil {
		return c.Error(nil, errcode.New(errcode.ERR1005.Code, errcode.ERR1005.Message).Error())
	}

	credentialID := c.Ctx.Params().GetIntDefault("id", 0)
	if credentialID == 0 {
		return c.Error(nil, errcode.New(errcode.ERR3110.Code, errcode.ERR3110.Message).Error())
	}

	var form struct {
		Name string `json:"name"`
	}
	if err := c.Ctx.ReadJSON(&form); err != nil {
		return c.Error(nil, err.Error())
	}

	service := v2service.NewWebauthnService(config.GetServerConfig(), c.Db)
	if err := service.RenameCredential(currentUser.Id, credentialID, form.Name); err != nil {
		return c.Error(nil, err.Error())
	}
	return c.Success(nil, "ok")
}

var _ model.User
