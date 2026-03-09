package api

import (
	"errors"
	"io"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/repository"
	v2service "rustdesk-api-server-pro/internal/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

var errAddressBookNotFound = errors.New("address book not found")

type basicController struct {
	Ctx iris.Context
	Db  *xorm.Engine
	Cfg *config.ServerConfig
}

func (c *basicController) GetUser() *model.User {
	user := c.Ctx.Values().Get(config.CurrentUserKey)
	if user != nil {
		return user.(*model.User)
	}
	return nil
}

func (c *basicController) GetToken() string {
	return c.Ctx.Values().Get(config.CurrentAuthTokenString).(string)
}

func (c *basicController) GetAuthToken() *model.AuthToken {
	return c.Ctx.Values().Get(config.CurrentAuthToken).(*model.AuthToken)
}

func (c *basicController) ok() mvc.Result {
	return mvc.Response{}
}

func (c *basicController) okText(text string) mvc.Result {
	return mvc.Response{Text: text}
}

func (c *basicController) fail(err error) mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"error": err.Error(),
		},
	}
}

func (c *basicController) failMsg(msg string) mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"error": msg,
		},
	}
}

func (c *basicController) readJSONBody(target interface{}) error {
	return c.Ctx.ReadJSON(target)
}

func (c *basicController) readBodyBytes() ([]byte, error) {
	return io.ReadAll(c.Ctx.Request().Body)
}

func (c *basicController) withTx(fn func(session *xorm.Session) error) error {
	session := c.Db.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}
	if err := fn(session); err != nil {
		_ = session.Rollback()
		return err
	}
	return session.Commit()
}

func (c *basicController) getAddressBookByGuid(userID int, guid string) (*model.AddressBook, error) {
	var ab model.AddressBook
	has, err := c.Db.Where("user_id = ? and guid = ?", userID, guid).Get(&ab)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errAddressBookNotFound
	}
	return &ab, nil
}

func (c *basicController) peerService() *v2service.PeerService {
	return v2service.NewPeerService(repository.NewXormPeerRepository(c.Db))
}

func (c *basicController) userService() *v2service.UserService {
	return v2service.NewUserService(repository.NewXormUserRepository(c.Db))
}

func (c *basicController) addressBookService() *v2service.AddressBookService {
	return v2service.NewAddressBookService(repository.NewXormAddressBookRepository(c.Db))
}

func (c *basicController) systemService() *v2service.SystemService {
	return v2service.NewSystemService(repository.NewXormSystemRepository(c.Db))
}

func (c *basicController) auditService() *v2service.AuditService {
	return v2service.NewAuditService(repository.NewXormAuditRepository(c.Db))
}

func (c *basicController) compatService() *v2service.CompatService {
	return v2service.NewCompatService(repository.NewXormCompatRepository(c.Db), c.Cfg, c.Db)
}

func (c *basicController) loginService() *v2service.LoginService {
	return v2service.NewLoginService()
}
