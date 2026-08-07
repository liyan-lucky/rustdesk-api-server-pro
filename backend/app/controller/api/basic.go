package api

import (
	"io"
	"log"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/app/service"
	"rustdesk-api-server-pro/config"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/internal/repository"
	v2service "rustdesk-api-server-pro/internal/service"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

var errAddressBookNotFound = errcode.New(errcode.ERR5001.Code, errcode.ERR5001.Message)

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
	v := c.Ctx.Values().Get(config.CurrentAuthTokenString)
	if v == nil {
		return ""
	}
	return v.(string)
}

func (c *basicController) GetAuthToken() *model.AuthToken {
	v := c.Ctx.Values().Get(config.CurrentAuthToken)
	if v == nil {
		return nil
	}
	return v.(*model.AuthToken)
}

func (c *basicController) ok() mvc.Result {
	return mvc.Response{}
}

func (c *basicController) okText(text string) mvc.Result {
	return mvc.Response{Text: text}
}

func extractErrCode(message string) string {
	message = strings.TrimSpace(message)
	if message == "" || !strings.HasPrefix(message, "ERR-") {
		return ""
	}
	idx := strings.Index(message, ":")
	if idx <= 0 {
		return message
	}
	code := message[:idx]
	if len(code) < 5 {
		return ""
	}
	return code
}

func (c *basicController) fail(err error) mvc.Result {
	user := c.GetUser()
	userId := 0
	userName := ""
	if user != nil {
		userId = user.Id
		userName = user.Username
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic in RecordErrorLog (api/fail): %v", r)
			}
		}()
		service.RecordErrorLog(c.Db, extractErrCode(err.Error()), err.Error(), "api", c.Ctx.Path(), c.Ctx.Method(), userId, userName, c.Ctx.RemoteAddr(), c.Ctx.GetHeader("User-Agent"))
	}()
	return mvc.Response{
		Object: iris.Map{
			"error": err.Error(),
		},
	}
}

func (c *basicController) failMsg(msg string) mvc.Result {
	user := c.GetUser()
	userId := 0
	userName := ""
	if user != nil {
		userId = user.Id
		userName = user.Username
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic in RecordErrorLog (api/failMsg): %v", r)
			}
		}()
		service.RecordErrorLog(c.Db, extractErrCode(msg), msg, "api", c.Ctx.Path(), c.Ctx.Method(), userId, userName, c.Ctx.RemoteAddr(), c.Ctx.GetHeader("User-Agent"))
	}()
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
	has, err := c.Db.Where("guid = ?", guid).Get(&ab)
	if err != nil || !has {
		if err != nil {
			return nil, err
		}
		return nil, errAddressBookNotFound
	}
	if ab.UserId == userID {
		return &ab, nil
	}
	if !ab.Shared {
		return nil, errAddressBookNotFound
	}
	if ab.Rule >= 2 {
		return &ab, nil
	}
	hasRules, err := c.Db.IsTableExist(new(model.AddressBookRule))
	if err != nil {
		return nil, err
	}
	if !hasRules {
		return nil, errAddressBookNotFound
	}
	allowed, err := c.Db.Where("ab_guid = ? and rule >= ? and (target_type = ? or (target_type = ? and target_guid = ?))", ab.Guid, 2, "everyone", "user", strconv.Itoa(userID)).Exist(new(model.AddressBookRule))
	if err != nil {
		return nil, err
	}
	if !allowed {
		hasGroups, groupErr := c.Db.IsTableExist(new(model.UserGroupMember))
		if groupErr != nil {
			return nil, groupErr
		}
		if hasGroups {
			allowed, err = c.Db.Where("ab_guid = ? and rule >= ? and target_type = ? and target_guid in (select group_guid from user_group_member where user_id = ?)", ab.Guid, 2, "user_group", userID).Exist(new(model.AddressBookRule))
			if err != nil {
				return nil, err
			}
		}
	}
	if !allowed {
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
