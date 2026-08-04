package admin

import (
	"encoding/json"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/internal/repository"
	v2service "rustdesk-api-server-pro/internal/service"
	"rustdesk-api-server-pro/util"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

var errAddressBookNotFound = errcode.New(errcode.ERR5001.Code, errcode.ERR5001.Message)
var errUnauthorized = errcode.ErrUnauthorized

type AddressBookController struct {
	basicController
}

func (c *AddressBookController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/ab/peers", "HandleAbPeers")
	b.Handle("GET", "/ab/shared-profiles", "HandleAbSharedProfiles")
	b.Handle("GET", "/ab/personal", "HandleAbPersonal")
	b.Handle("GET", "/ab/list", "HandleAbList")
	b.Handle("POST", "/ab/tags/{guid:string}", "HandleAbTags")
	b.Handle("GET", "/ab/tags", "HandleAbAllTags")
	b.Handle("POST", "/ab/tag/add/{guid:string}", "HandleAbTagAdd")
	b.Handle("PUT", "/ab/tag/update/{guid:string}", "HandleAbTagUpdate")
	b.Handle("PUT", "/ab/tag/rename/{guid:string}", "HandleAbTagRename")
	b.Handle("POST", "/ab/peer/add/{guid:string}", "HandleAbPeerAdd")
	b.Handle("PUT", "/ab/peer/update/{guid:string}", "HandleAbPeerUpdate")
	b.Handle("DELETE", "/ab/peer/{guid:string}", "HandleAbPeerDelete")
	b.Handle("DELETE", "/ab/tag/{guid:string}", "HandleAbTagDelete")
	b.Handle("POST", "/ab/shared/add", "HandleAbSharedAdd")
	b.Handle("PUT", "/ab/shared/update", "HandleAbSharedUpdate")
	b.Handle("DELETE", "/ab/shared", "HandleAbSharedDelete")
	b.Handle("GET", "/ab/rules/{guid:string}", "HandleAbRules")
	b.Handle("POST", "/ab/rule/{guid:string}", "HandleAbRuleAdd")
	b.Handle("PUT", "/ab/rule/{guid:string}", "HandleAbRuleUpdate")
	b.Handle("DELETE", "/ab/rule/{guid:string}", "HandleAbRuleDelete")
}

func (c *AddressBookController) addressBookService() *v2service.AddressBookService {
	return v2service.NewAddressBookService(repository.NewXormAddressBookRepository(c.Db))
}

func (c *AddressBookController) HandleAbPeers() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	abGuid := c.Ctx.URLParamDefault("ab", "")

	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	if abGuid == "" {
		return c.listAccountPeers(user, current, pageSize)
	}

	if user.IsAdmin {
		return c.listAllPeers(abGuid, current, pageSize)
	}

	result, err := c.addressBookService().ListPeers(core.AddressBookPeerListQuery{
		UserID:   user.Id,
		AbGuid:   abGuid,
		Current:  current,
		PageSize: pageSize,
	})
	if err != nil {
		return c.dbError(err)
	}

	return c.Success(iris.Map{
		"total":   result.Total,
		"records": peersToRecords(result.Items),
		"current": current,
		"size":    pageSize,
	}, "ok")
}

func (c *AddressBookController) listAccountPeers(user *model.User, current, pageSize int) mvc.Result {
	books := make([]model.AddressBook, 0)
	if user.IsAdmin {
		if err := c.Db.Find(&books); err != nil {
			return c.dbError(err)
		}
	} else if err := c.Db.Where("user_id = ?", user.Id).Find(&books); err != nil {
		return c.dbError(err)
	}
	bookIDs := make([]int, 0, len(books))
	bookByID := make(map[int]model.AddressBook, len(books))
	for _, book := range books {
		bookIDs = append(bookIDs, book.Id)
		bookByID[book.Id] = book
	}
	if len(bookIDs) == 0 {
		return c.Success(iris.Map{"total": 0, "records": []iris.Map{}, "current": current, "size": pageSize}, "ok")
	}
	query := func() *xorm.Session {
		q := c.Db.In("ab_id", bookIDs)
		filters := map[string]string{"rustdesk_id": "id", "username": "username", "hostname": "hostname", "platform": "platform", "alias": "alias", "hash": "hash", "note": "note", "tags": "tags"}
		for column, parameter := range filters {
			if value := c.Ctx.URLParamDefault(parameter, ""); value != "" {
				q = q.And(column+" like ?", "%"+value+"%")
			}
		}
		if guid := c.Ctx.URLParamDefault("ab_guid", ""); guid != "" {
			for _, book := range books {
				if book.Guid == guid {
					q = q.And("ab_id = ?", book.Id)
					return q
				}
			}
			return q.And("1 = 0")
		}
		if name := c.Ctx.URLParamDefault("ab_name", ""); name != "" {
			matchingIDs := make([]int, 0)
			for _, book := range books {
				if strings.Contains(strings.ToLower(book.Name), strings.ToLower(name)) {
					matchingIDs = append(matchingIDs, book.Id)
				}
			}
			if len(matchingIDs) == 0 {
				return q.And("1 = 0")
			}
			q = q.In("ab_id", matchingIDs)
		}
		return q.Desc("id")
	}
	pagination := db.NewPagination(current, pageSize)
	peers := make([]model.Peer, 0)
	if err := pagination.Paginate(query, &model.Peer{}, &peers); err != nil {
		return c.dbError(err)
	}
	records := make([]iris.Map, 0, len(peers))
	for _, peer := range peers {
		book := bookByID[peer.AbId]
		peerTags := make([]string, 0)
		_ = json.Unmarshal([]byte(peer.Tags), &peerTags)
		records = append(records, iris.Map{"id": peer.RustdeskId, "hash": peer.Hash, "username": peer.Username, "hostname": peer.Hostname, "platform": peer.Platform, "alias": peer.Alias, "tags": peerTags, "note": peer.Note, "ab_id": book.Id, "ab_guid": book.Guid, "ab_name": book.Name, "owner": book.Owner})
	}
	return c.Success(iris.Map{"total": pagination.TotalCount, "records": records, "current": current, "size": pageSize}, "ok")
}

func (c *AddressBookController) listAllPeers(abGuid string, current, pageSize int) mvc.Result {
	query := func() *xorm.Session {
		q := c.Db.Table(&model.Peer{})
		if abGuid != "" {
			var ab model.AddressBook
			has, err := c.Db.Where("guid = ?", abGuid).Get(&ab)
			if err != nil || !has {
				return q.Where("1 = 0")
			}
			q = q.Where("ab_id = ?", ab.Id)
		}
		return q
	}

	pagination := db.NewPagination(current, pageSize)
	peerList := make([]model.Peer, 0)
	if err := pagination.Paginate(query, &model.Peer{}, &peerList); err != nil {
		return c.dbError(err)
	}

	records := make([]iris.Map, 0, len(peerList))
	for _, peer := range peerList {
		peerTags := make([]string, 0)
		_ = json.Unmarshal([]byte(peer.Tags), &peerTags)
		records = append(records, iris.Map{
			"id":       peer.RustdeskId,
			"hash":     peer.Hash,
			"username": peer.Username,
			"hostname": peer.Hostname,
			"platform": peer.Platform,
			"alias":    peer.Alias,
			"tags":     peerTags,
			"note":     peer.Note,
		})
	}

	return c.Success(iris.Map{
		"total":   pagination.TotalCount,
		"records": records,
		"current": current,
		"size":    pageSize,
	}, "ok")
}

func (c *AddressBookController) HandleAbSharedProfiles() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)

	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	query := func() *xorm.Session {
		q := c.Db.Where("shared = ?", true)
		if !user.IsAdmin {
			q = q.And("user_id = ?", user.Id)
		}
		return q
	}
	var books []model.AddressBook
	total, err := query().Count(new(model.AddressBook))
	if err != nil {
		return c.dbError(err)
	}
	if err = query().Limit(pageSize, (current-1)*pageSize).Find(&books); err != nil {
		return c.dbError(err)
	}
	records := make([]iris.Map, 0, len(books))
	for _, ab := range books {
		records = append(records, iris.Map{
			"guid":  ab.Guid,
			"name":  ab.Name,
			"owner": ab.Owner,
			"note":  ab.Note,
			"rule":  ab.Rule,
		})
	}

	return c.Success(iris.Map{
		"total":   total,
		"records": records,
		"current": current,
		"size":    pageSize,
	}, "ok")
}

func (c *AddressBookController) HandleAbPersonal() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	result, err := c.addressBookService().EnsurePersonalAddressBook(core.PersonalAddressBookEnsureCommand{
		UserID:         user.Id,
		Username:       user.Username,
		DefaultNote:    "default address book",
		DefaultRule:    3,
		DefaultMaxPeer: 0,
	})
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(iris.Map{"guid": result.Guid}, "ok")
}

func (c *AddressBookController) HandleAbList() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	list := make([]model.AddressBook, 0)
	if user.IsAdmin {
		if err := c.Db.Find(&list); err != nil {
			return c.dbError(err)
		}
	} else {
		if err := c.Db.Where("user_id = ?", user.Id).Find(&list); err != nil {
			return c.dbError(err)
		}
	}
	records := make([]iris.Map, 0, len(list))
	for _, ab := range list {
		var owner model.User
		hasOwner, ownerErr := c.Db.Where("id = ?", ab.UserId).Get(&owner)
		if ownerErr != nil {
			return c.Error(nil, ownerErr.Error())
		}
		if !hasOwner {
			continue
		}
		records = append(records, iris.Map{
			"id":               ab.Id,
			"user_id":          ab.UserId,
			"guid":             ab.Guid,
			"name":             ab.Name,
			"owner":            owner.Username,
			"note":             ab.Note,
			"rule":             ab.Rule,
			"max_peer":         ab.MaxPeer,
			"shared":           ab.Shared,
			"created_by_admin": ab.CreatedByAdmin,
		})
	}
	return c.Success(records, "ok")
}

func (c *AddressBookController) HandleAbTagAdd() mvc.Result {
	ab, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), true)
	if err != nil {
		return c.dbError(err)
	}
	var body struct {
		Name  string `json:"name"`
		Color int64  `json:"color"`
	}
	if err = c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	if body.Name == "" {
		return c.Error(nil, "name required")
	}
	count, err := c.Db.Where("ab_id = ? and name = ?", ab.Id, body.Name).Count(new(model.AddressBookTag))
	if err != nil {
		return c.dbError(err)
	}
	if count > 0 {
		return c.Error(nil, "tag already exists")
	}
	_, err = c.Db.Insert(&model.AddressBookTag{UserId: ab.UserId, AbId: ab.Id, Name: body.Name, Color: body.Color})
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}

func (c *AddressBookController) HandleAbTagUpdate() mvc.Result {
	ab, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), true)
	if err != nil {
		return c.dbError(err)
	}
	var body struct {
		Name  string `json:"name"`
		Color int64  `json:"color"`
	}
	if err = c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	affected, err := c.Db.Where("ab_id = ? and name = ?", ab.Id, body.Name).Cols("color").Update(&model.AddressBookTag{Color: body.Color})
	if err != nil {
		return c.dbError(err)
	}
	if affected == 0 {
		return c.Error(nil, "tag not found")
	}
	return c.Success(nil, "ok")
}

func (c *AddressBookController) HandleAbTagRename() mvc.Result {
	ab, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), true)
	if err != nil {
		return c.dbError(err)
	}
	var body struct {
		Old string `json:"old"`
		New string `json:"new"`
	}
	if err = c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	if body.Old == "" || body.New == "" {
		return c.Error(nil, "old and new required")
	}
	if err = c.addressBookService().RenameTag(core.AddressBookTagRenameCommand{UserID: ab.UserId, AbID: ab.Id, Old: body.Old, New: body.New}); err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}

type adminPeerPayload struct {
	ID       string   `json:"id"`
	Hash     string   `json:"hash"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Hostname string   `json:"hostname"`
	Platform string   `json:"platform"`
	Alias    string   `json:"alias"`
	Tags     []string `json:"tags"`
	Note     string   `json:"note"`
}

func (c *AddressBookController) HandleAbPeerAdd() mvc.Result {
	ab, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), true)
	if err != nil {
		return c.dbError(err)
	}
	var body adminPeerPayload
	if err = c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	if body.ID == "" {
		return c.Error(nil, "device id required")
	}
	if err = c.addressBookService().AddPeer(core.AddressBookPeerCreateCommand{UserID: ab.UserId, AbID: ab.Id, RustdeskID: body.ID, Hash: body.Hash, Username: body.Username, Password: body.Password, Hostname: body.Hostname, Platform: body.Platform, Alias: body.Alias, Tags: body.Tags, Note: body.Note}); err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}

func (c *AddressBookController) HandleAbPeerUpdate() mvc.Result {
	ab, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), true)
	if err != nil {
		return c.dbError(err)
	}
	var body adminPeerPayload
	if err = c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	tags, err := json.Marshal(body.Tags)
	if err != nil {
		return c.dbError(err)
	}
	tagsJSON := string(tags)
	cmd := core.AddressBookPeerUpdateCommand{UserID: ab.UserId, AbID: ab.Id, RustdeskID: body.ID, Tags: &tagsJSON, Alias: &body.Alias, Hash: &body.Hash, Password: &body.Password, Note: &body.Note, Username: &body.Username, Hostname: &body.Hostname, Platform: &body.Platform}
	has, err := c.addressBookService().UpdatePeer(cmd)
	if err != nil {
		return c.dbError(err)
	}
	if !has {
		return c.Error(nil, "peer not found")
	}
	return c.Success(nil, "ok")
}

func (c *AddressBookController) HandleAbSharedAdd() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	var body struct {
		UserID  int    `json:"user_id"`
		Name    string `json:"name"`
		Note    string `json:"note"`
		Rule    int    `json:"rule"`
		MaxPeer int    `json:"max_peer"`
	}
	if err := c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	if body.Name == "" {
		return c.Error(nil, "name required")
	}
	owner := user
	createdByAdmin := false
	if user.IsAdmin && body.UserID > 0 && body.UserID != user.Id {
		var target model.User
		has, targetErr := c.Db.Where("id = ?", body.UserID).Get(&target)
		if targetErr != nil {
			return c.Error(nil, targetErr.Error())
		}
		if !has {
			return c.Error(nil, "user not found")
		}
		owner = &target
		createdByAdmin = true
	}
	book := model.AddressBook{UserId: owner.Id, Guid: util.GetUUID(), Name: body.Name, Owner: owner.Username, Note: body.Note, Rule: body.Rule, MaxPeer: body.MaxPeer, Shared: true, CreatedByAdmin: createdByAdmin}
	if _, err := c.Db.Insert(&book); err != nil {
		return c.dbError(err)
	}
	return c.Success(iris.Map{"guid": book.Guid}, "ok")
}

func (c *AddressBookController) HandleAbSharedUpdate() mvc.Result {
	var body struct {
		Guid    string `json:"guid"`
		Name    string `json:"name"`
		Note    string `json:"note"`
		Rule    int    `json:"rule"`
		MaxPeer int    `json:"max_peer"`
	}
	if err := c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	book, err := c.manageableAddressBook(body.Guid, false)
	if err != nil {
		return c.dbError(err)
	}
	if !book.Shared {
		return c.Error(nil, "personal address book is read-only")
	}
	if body.Name == "" {
		return c.Error(nil, "name required")
	}
	_, err = c.Db.Where("id = ?", book.Id).Cols("name", "note", "rule", "max_peer").Update(&model.AddressBook{Name: body.Name, Note: body.Note, Rule: body.Rule, MaxPeer: body.MaxPeer})
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}

func (c *AddressBookController) HandleAbSharedDelete() mvc.Result {
	var guids []string
	if err := c.Ctx.ReadJSON(&guids); err != nil {
		return c.dbError(err)
	}
	for _, guid := range guids {
		book, err := c.manageableAddressBook(guid, false)
		if err != nil {
			return c.dbError(err)
		}
		if !book.Shared {
			return c.Error(nil, "personal address book cannot be deleted")
		}
		user := c.GetUser()
		if !canDeleteAddressBook(user, book) {
			return c.Error(nil, "administrator-created address book cannot be deleted by user")
		}
		if err = c.deleteAddressBook(book); err != nil {
			return c.dbError(err)
		}
	}
	return c.Success(nil, "ok")
}

func canDeleteAddressBook(user *model.User, book *model.AddressBook) bool {
	return user != nil && book != nil && (user.IsAdmin || !book.CreatedByAdmin)
}

func (c *AddressBookController) HandleAbRules() mvc.Result {
	book, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), false)
	if err != nil {
		return c.dbError(err)
	}
	var rules []model.AddressBookRule
	if err = c.Db.Where("ab_guid = ?", book.Guid).Find(&rules); err != nil {
		return c.dbError(err)
	}
	data := make([]iris.Map, 0, len(rules))
	for _, rule := range rules {
		data = append(data, addressBookRuleData(rule))
	}
	return c.Success(data, "ok")
}

type addressBookRulePayload struct {
	Guid       string `json:"guid"`
	TargetType string `json:"target_type"`
	TargetGuid string `json:"target_guid"`
	Rule       int    `json:"rule"`
}

func (c *AddressBookController) HandleAbRuleAdd() mvc.Result {
	book, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), false)
	if err != nil {
		return c.dbError(err)
	}
	var body addressBookRulePayload
	if err = c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	if body.Rule < 1 || body.Rule > 3 {
		return c.Error(nil, "invalid rule")
	}
	rule := model.AddressBookRule{Guid: util.GetUUID(), AbGuid: book.Guid, TargetType: body.TargetType, TargetGuid: body.TargetGuid, Rule: body.Rule}
	if rule.TargetType == "" || rule.TargetGuid == "" {
		return c.Error(nil, "target type and target guid required")
	}
	if _, err = c.Db.Insert(&rule); err != nil {
		return c.dbError(err)
	}
	return c.Success(addressBookRuleData(rule), "ok")
}
func (c *AddressBookController) HandleAbRuleUpdate() mvc.Result {
	book, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), false)
	if err != nil {
		return c.dbError(err)
	}
	var body addressBookRulePayload
	if err = c.Ctx.ReadJSON(&body); err != nil {
		return c.dbError(err)
	}
	if body.Rule < 1 || body.Rule > 3 {
		return c.Error(nil, "invalid rule")
	}
	if body.Guid == "" || body.TargetType == "" || body.TargetGuid == "" {
		return c.Error(nil, "guid and target required")
	}
	update := model.AddressBookRule{TargetType: body.TargetType, TargetGuid: body.TargetGuid, Rule: body.Rule}
	affected, err := c.Db.Where("guid = ? and ab_guid = ?", body.Guid, book.Guid).Cols("target_type", "target_guid", "rule").Update(&update)
	if err != nil {
		return c.dbError(err)
	}
	if affected == 0 {
		return c.Error(nil, "rule not found")
	}
	return c.Success(nil, "ok")
}

func addressBookRuleData(rule model.AddressBookRule) iris.Map {
	return iris.Map{
		"guid":        rule.Guid,
		"ab_guid":     rule.AbGuid,
		"target_type": rule.TargetType,
		"target_guid": rule.TargetGuid,
		"rule":        rule.Rule,
		"created_at":  rule.CreatedAt,
		"updated_at":  rule.UpdatedAt,
	}
}
func (c *AddressBookController) HandleAbRuleDelete() mvc.Result {
	book, err := c.manageableAddressBook(c.Ctx.Params().Get("guid"), false)
	if err != nil {
		return c.dbError(err)
	}
	var guids []string
	if err = c.Ctx.ReadJSON(&guids); err != nil {
		return c.dbError(err)
	}
	_, err = c.Db.Where("ab_guid = ?", book.Guid).In("guid", guids).Delete(new(model.AddressBookRule))
	if err != nil {
		return c.dbError(err)
	}
	return c.Success(nil, "ok")
}

func (c *AddressBookController) manageableAddressBook(guid string, allowSharedWriter bool) (*model.AddressBook, error) {
	user := c.GetUser()
	if user == nil {
		return nil, errAddressBookNotFound
	}
	var book model.AddressBook
	has, err := c.Db.Where("guid = ?", guid).Get(&book)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errAddressBookNotFound
	}
	if user.IsAdmin || book.UserId == user.Id {
		return &book, nil
	}
	if allowSharedWriter && book.Shared && book.Rule >= 2 {
		return &book, nil
	}
	return nil, errAddressBookNotFound
}
func (c *AddressBookController) deleteAddressBook(book *model.AddressBook) error {
	session := c.Db.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	if _, err := session.Where("ab_guid = ?", book.Guid).Delete(new(model.AddressBookRule)); err != nil {
		_ = session.Rollback()
		return err
	}
	for _, value := range []any{new(model.AddressBookTag), new(model.Peer)} {
		if _, err := session.Where("ab_id = ?", book.Id).Delete(value); err != nil {
			_ = session.Rollback()
			return err
		}
	}
	if _, err := session.Where("id = ?", book.Id).Delete(new(model.AddressBook)); err != nil {
		_ = session.Rollback()
		return err
	}
	return session.Commit()
}

func (c *AddressBookController) HandleAbTags() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}

	if user.IsAdmin {
		return c.listAllTags(abGuid)
	}

	tags, err := c.addressBookService().ListTags(core.AddressBookTagListQuery{
		UserID: user.Id,
		AbGuid: abGuid,
	})
	if err != nil {
		return c.dbError(err)
	}

	return c.Success(tagsToData(tags), "ok")
}

func (c *AddressBookController) HandleAbAllTags() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	books := make([]model.AddressBook, 0)
	if user.IsAdmin {
		if err := c.Db.Find(&books); err != nil {
			return c.dbError(err)
		}
	} else if err := c.Db.Where("user_id = ?", user.Id).Find(&books); err != nil {
		return c.dbError(err)
	}
	data := make([]iris.Map, 0)
	for _, book := range books {
		tags := make([]model.AddressBookTag, 0)
		if err := c.Db.Where("ab_id = ?", book.Id).Find(&tags); err != nil {
			return c.dbError(err)
		}
		for _, tag := range tags {
			data = append(data, iris.Map{
				"id": tag.Id, "ab_id": tag.AbId, "ab_guid": book.Guid,
				"ab_name": book.Name, "owner": book.Owner, "name": tag.Name, "color": tag.Color,
			})
		}
	}
	return c.Success(data, "ok")
}

func (c *AddressBookController) listAllTags(abGuid string) mvc.Result {
	var ab model.AddressBook
	has, err := c.Db.Where("guid = ?", abGuid).Get(&ab)
	if err != nil {
		return c.dbError(err)
	}
	if !has {
		return c.Error(nil, errAddressBookNotFound.Error())
	}

	tags := make([]model.AddressBookTag, 0)
	if err := c.Db.Where("ab_id = ?", ab.Id).Find(&tags); err != nil {
		return c.dbError(err)
	}

	data := make([]iris.Map, 0, len(tags))
	for _, t := range tags {
		data = append(data, iris.Map{
			"id":    t.Id,
			"ab_id": t.AbId,
			"name":  t.Name,
			"color": t.Color,
		})
	}

	return c.Success(data, "ok")
}

func (c *AddressBookController) HandleAbPeerDelete() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var ids []string
	if err := c.Ctx.ReadBody(&ids); err != nil {
		return c.dbError(err)
	}
	if len(ids) == 0 {
		return c.Error(nil, "NoPeerIds")
	}

	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	ab, err := c.findAddressBook(user, abGuid)
	if err != nil {
		return c.dbError(err)
	}

	if err := c.addressBookService().DeletePeers(core.AddressBookPeerDeleteCommand{
		UserID: ab.UserId,
		AbID:   ab.Id,
		IDs:    ids,
	}); err != nil {
		return c.dbError(err)
	}

	return c.Success(nil, "ok")
}

func (c *AddressBookController) HandleAbTagDelete() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var names []string
	if err := c.Ctx.ReadBody(&names); err != nil {
		return c.dbError(err)
	}
	if len(names) == 0 {
		return c.Error(nil, "NoTagNames")
	}

	user := c.GetUser()
	if user == nil {
		return c.Error(nil, errUnauthorized.Error())
	}
	ab, err := c.findAddressBook(user, abGuid)
	if err != nil {
		return c.dbError(err)
	}

	if err := c.addressBookService().DeleteTags(core.AddressBookTagDeleteCommand{
		UserID: ab.UserId,
		AbID:   ab.Id,
		Names:  names,
	}); err != nil {
		return c.dbError(err)
	}

	return c.Success(nil, "ok")
}

func (c *AddressBookController) findAddressBook(user *model.User, guid string) (*model.AddressBook, error) {
	var ab model.AddressBook
	if user.IsAdmin {
		has, err := c.Db.Where("guid = ?", guid).Get(&ab)
		if err != nil {
			return nil, err
		}
		if !has {
			return nil, errAddressBookNotFound
		}
	} else {
		has, err := c.Db.Where("user_id = ? and guid = ?", user.Id, guid).Get(&ab)
		if err != nil {
			return nil, err
		}
		if !has {
			return nil, errAddressBookNotFound
		}
	}
	return &ab, nil
}

func peersToRecords(items []core.AddressBookPeerView) []iris.Map {
	records := make([]iris.Map, 0, len(items))
	for _, p := range items {
		records = append(records, iris.Map{
			"id":       p.ID,
			"hash":     p.Hash,
			"username": p.Username,
			"hostname": p.Hostname,
			"platform": p.Platform,
			"alias":    p.Alias,
			"tags":     p.Tags,
			"note":     p.Note,
		})
	}
	return records
}

func tagsToData(tags []core.AddressBookTagView) []iris.Map {
	data := make([]iris.Map, 0, len(tags))
	for _, t := range tags {
		data = append(data, iris.Map{
			"name":  t.Name,
			"color": t.Color,
		})
	}
	return data
}
