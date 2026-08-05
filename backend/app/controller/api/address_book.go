package api

import (
	"encoding/json"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/errcode"
	"rustdesk-api-server-pro/internal/transport/httpdto"
	"rustdesk-api-server-pro/util"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AddressBookController struct {
	basicController
}

func (c *AddressBookController) ownedAddressBook(guid string) (*model.AddressBook, error) {
	user := c.GetUser()
	if user == nil {
		return nil, errAddressBookNotFound
	}
	var ab model.AddressBook
	q := c.Db.Where("guid = ?", guid)
	if !user.IsAdmin {
		q = q.And("user_id = ?", user.Id)
	}
	has, err := q.Get(&ab)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errAddressBookNotFound
	}
	return &ab, nil
}

func (c *AddressBookController) HandleAbSharedAdd() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.fail(errcode.ErrUnauthorized)
	}
	var body map[string]any
	if err := c.readJSONBody(&body); err != nil {
		return c.fail(err)
	}
	name := stringFromAny(body["name"])
	if name == "" {
		return c.fail(errcode.New(errcode.ERR5002.Code, errcode.ERR5002.Message))
	}
	ab := model.AddressBook{UserId: user.Id, Guid: util.GetUUID(), Name: name, Owner: user.Username, Note: stringFromAny(body["note"]), Rule: 0, Shared: true}
	if _, err := c.Db.Insert(&ab); err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: iris.Map{"guid": ab.Guid, "name": ab.Name}}
}

func (c *AddressBookController) HandleAbSharedUpdate() mvc.Result {
	var body map[string]any
	if err := c.readJSONBody(&body); err != nil {
		return c.fail(err)
	}
	guid := stringFromAny(body["guid"])
	ab, err := c.ownedAddressBook(guid)
	if err != nil {
		return c.fail(err)
	}
	cols := []string{}
	update := model.AddressBook{}
	if v, ok := body["name"]; ok {
		update.Name = stringFromAny(v)
		cols = append(cols, "name")
	}
	if v, ok := body["note"]; ok {
		update.Note = stringFromAny(v)
		cols = append(cols, "note")
	}
	if v, ok := body["owner"]; ok && c.GetUser().IsAdmin {
		update.Owner = stringFromAny(v)
		cols = append(cols, "owner")
	}
	if len(cols) > 0 {
		if _, err = c.Db.Where("id = ?", ab.Id).Cols(cols...).Update(&update); err != nil {
			return c.fail(err)
		}
	}
	return c.ok()
}

func (c *AddressBookController) HandleAbSharedDelete() mvc.Result {
	var guids []string
	if err := c.Ctx.ReadBody(&guids); err != nil {
		return c.fail(err)
	}
	for _, guid := range guids {
		ab, err := c.ownedAddressBook(guid)
		if err != nil {
			return c.fail(err)
		}
		session := c.Db.NewSession()
		if err = session.Begin(); err != nil {
			return c.fail(err)
		}
		_, err = session.Where("ab_guid = ?", guid).Delete(&model.AddressBookRule{})
		if err == nil {
			_, err = session.Where("ab_id = ?", ab.Id).Delete(&model.AddressBookTag{})
		}
		if err == nil {
			_, err = session.Where("ab_id = ?", ab.Id).Delete(&model.Peer{})
		}
		if err == nil {
			_, err = session.Where("id = ?", ab.Id).Delete(&model.AddressBook{})
		}
		if err != nil {
			_ = session.Rollback()
			session.Close()
			return c.fail(err)
		}
		err = session.Commit()
		session.Close()
		if err != nil {
			return c.fail(err)
		}
	}
	return c.ok()
}

func (c *AddressBookController) HandleAbRules() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.fail(errcode.ErrUnauthorized)
	}
	abGuid := c.Ctx.URLParamDefault("ab", "")
	if abGuid == "" {
		abGuid = c.Ctx.URLParamDefault("guid", "")
	}
	if _, err := c.ownedAddressBook(abGuid); err != nil {
		return c.fail(err)
	}
	var rules []model.AddressBookRule
	if err := c.Db.Where("ab_guid = ?", abGuid).Find(&rules); err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: iris.Map{"total": len(rules), "data": rules}}
}

func (c *AddressBookController) HandleAbRuleAdd() mvc.Result {
	var body map[string]any
	if err := c.readJSONBody(&body); err != nil {
		return c.fail(err)
	}
	abGuid := stringFromAny(body["ab_guid"])
	if abGuid == "" {
		abGuid = stringFromAny(body["ab"])
	}
	if _, err := c.ownedAddressBook(abGuid); err != nil {
		return c.fail(err)
	}
	typeName := stringFromAny(body["type"])
	if typeName == "" {
		typeName = stringFromAny(body["target_type"])
	}
	if typeName == "" {
		typeName = "everyone"
	}
	target := stringFromAny(body["target_guid"])
	if target == "" {
		target = stringFromAny(body["user"])
	}
	if target == "" {
		target = stringFromAny(body["group"])
	}
	rule := intFromAny(body["rule"])
	if rule < 1 || rule > 3 {
		return c.fail(errcode.New(errcode.ERR5012.Code, errcode.ERR5012.Message))
	}
	row := model.AddressBookRule{Guid: util.GetUUID(), AbGuid: abGuid, TargetType: typeName, TargetGuid: target, Rule: rule}
	if _, err := c.Db.Insert(&row); err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: row}
}

func (c *AddressBookController) HandleAbRuleUpdate() mvc.Result {
	var body map[string]any
	if err := c.readJSONBody(&body); err != nil {
		return c.fail(err)
	}
	guid := stringFromAny(body["guid"])
	var row model.AddressBookRule
	has, err := c.Db.Where("guid = ?", guid).Get(&row)
	if err != nil || !has {
		return c.fail(errcode.New(errcode.ERR5015.Code, errcode.ERR5015.Message))
	}
	if _, err = c.ownedAddressBook(row.AbGuid); err != nil {
		return c.fail(err)
	}
	rule := intFromAny(body["rule"])
	if rule < 1 || rule > 3 {
		return c.fail(errcode.New(errcode.ERR5012.Code, errcode.ERR5012.Message))
	}
	_, err = c.Db.Where("guid = ?", guid).Cols("rule").Update(&model.AddressBookRule{Rule: rule})
	if err != nil {
		return c.fail(err)
	}
	return c.ok()
}

func (c *AddressBookController) HandleAbRulesDelete() mvc.Result {
	var guids []string
	if err := c.Ctx.ReadBody(&guids); err != nil {
		return c.fail(err)
	}
	for _, guid := range guids {
		var row model.AddressBookRule
		has, err := c.Db.Where("guid = ?", guid).Get(&row)
		if err != nil || !has {
			continue
		}
		if _, err = c.ownedAddressBook(row.AbGuid); err != nil {
			return c.fail(err)
		}
		if _, err = c.Db.Where("guid = ?", guid).Delete(&model.AddressBookRule{}); err != nil {
			return c.fail(err)
		}
	}
	return c.ok()
}

func (c *AddressBookController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "ab/get", "HandleAbGet")
	b.Handle("POST", "ab/get", "HandleAbGet")
	b.Handle("GET", "ab/personal", "HandleAbPersonal")
	b.Handle("POST", "ab/personal", "HandleAbPersonal")
	b.Handle("GET", "ab/settings", "HandleAbSettings")
	b.Handle("POST", "ab/settings", "HandleAbSettings")
	b.Handle("GET", "ab/shared-profiles", "HandleAbSharedProfiles")
	b.Handle("POST", "ab/shared-profiles", "HandleAbSharedProfiles")
	b.Handle("GET", "ab/shared/profiles", "HandleAbSharedProfiles")
	b.Handle("POST", "ab/shared/profiles", "HandleAbSharedProfiles")
	b.Handle("GET", "ab/shared_profiles", "HandleAbSharedProfiles")
	b.Handle("POST", "ab/shared_profiles", "HandleAbSharedProfiles")
	b.Handle("POST", "ab/shared/add", "HandleAbSharedAdd")
	b.Handle("PUT", "ab/shared/update/profile", "HandleAbSharedUpdate")
	b.Handle("DELETE", "ab/shared", "HandleAbSharedDelete")
	b.Handle("GET", "ab/rules", "HandleAbRules")
	b.Handle("POST", "ab/rule", "HandleAbRuleAdd")
	b.Handle("PATCH", "ab/rule", "HandleAbRuleUpdate")
	b.Handle("DELETE", "ab/rules", "HandleAbRulesDelete")
}

// HandleAbGet keeps the legacy Sciter client's /api/ab/get pull endpoint
// equivalent to the Flutter client's GET /api/ab endpoint.
func (c *AddressBookController) HandleAbGet() mvc.Result {
	return c.GetAb()
}

func (c *AddressBookController) GetAb() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.fail(errcode.ErrUnauthorized)
	}
	result, err := c.addressBookService().GetLegacyAddressBook(core.LegacyAddressBookGetQuery{
		UserID: user.Id,
	})
	if err != nil {
		return c.fail(err)
	}
	resp, err := httpdto.NewLegacyAddressBookResponse(user.LicensedDevices, result)
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: resp}
}

func (c *AddressBookController) PostAb() mvc.Result {
	var abForm api.AbForm
	err := c.readJSONBody(&abForm)
	if err != nil {
		c.recordAPIOperationAudit("ab_legacy_replace", "address_book", "legacy", nil, nil, "failure", err.Error())
		return c.fail(err)
	}
	var abData api.AbData
	err = json.Unmarshal([]byte(abForm.Data), &abData)
	if err != nil {
		c.recordAPIOperationAudit("ab_legacy_replace", "address_book", "legacy", nil, nil, "failure", err.Error())
		return c.fail(err)
	}
	var tagColors map[string]int64
	err = json.Unmarshal([]byte(abData.TagColors), &tagColors)
	if err != nil {
		c.recordAPIOperationAudit("ab_legacy_replace", "address_book", "legacy", nil, sanitizeLegacyAddressBookForAudit(abData), "failure", err.Error())
		return c.fail(err)
	}

	user := c.GetUser()
	if user == nil {
		return c.fail(errcode.ErrUnauthorized)
	}
	if user.LicensedDevices > 0 && len(abData.Peers) > user.LicensedDevices {
		c.recordAPIOperationAudit("ab_legacy_replace", "address_book", "legacy", nil, sanitizeLegacyAddressBookForAudit(abData), "failure", "Number of equipment in excess of licenses")
		return c.fail(errcode.New(errcode.ERR5018.Code, errcode.ERR5018.Message))
	}

	cmd := core.LegacyAddressBookReplaceCommand{
		UserID: user.Id,
		Tags:   make([]core.LegacyAddressBookTagEntry, 0, len(abData.Tags)),
		Peers:  make([]core.LegacyAddressBookPeerEntry, 0, len(abData.Peers)),
	}
	for _, tag := range abData.Tags {
		cmd.Tags = append(cmd.Tags, core.LegacyAddressBookTagEntry{
			Name:  tag,
			Color: tagColors[tag],
		})
	}
	for _, peer := range abData.Peers {
		cmd.Peers = append(cmd.Peers, core.LegacyAddressBookPeerEntry{
			RustdeskID: peer.Id,
			Hash:       peer.Hash,
			Username:   peer.Username,
			Hostname:   peer.Hostname,
			Platform:   peer.Platform,
			Alias:      peer.Alias,
			Tags:       peer.Tags,
			Note:       peer.Note,
		})
	}
	if err = c.addressBookService().ReplaceLegacyAddressBookData(cmd); err != nil {
		c.recordAPIOperationAudit("ab_legacy_replace", "address_book", "legacy", nil, sanitizeLegacyAddressBookForAudit(abData), "failure", err.Error())
		return c.fail(err)
	}

	c.recordAPIOperationAudit("ab_legacy_replace", "address_book", "legacy", nil, sanitizeLegacyAddressBookForAudit(abData), "success", "")
	return c.ok()
}

func (c *AddressBookController) HandleAbPersonal() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.fail(errcode.ErrUnauthorized)
	}
	result, err := c.addressBookService().EnsurePersonalAddressBook(core.PersonalAddressBookEnsureCommand{
		UserID:         user.Id,
		Username:       user.Username,
		DefaultNote:    "default address book",
		DefaultRule:    3,
		DefaultMaxPeer: 0,
	})
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: httpdto.NewPersonalAddressBookResponse(result)}
}

func (c *AddressBookController) HandleAbSettings() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.fail(errcode.ErrUnauthorized)
	}
	result, err := c.addressBookService().GetSettings(core.AddressBookSettingsQuery{UserID: user.Id})
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: httpdto.NewAddressBookSettingsResponse(result)}
}

func (c *AddressBookController) HandleAbSharedProfiles() mvc.Result {
	user := c.GetUser()
	if user == nil {
		return c.fail(errcode.ErrUnauthorized)
	}
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	result, err := c.addressBookService().ListSharedProfiles(core.SharedAddressBookListQuery{
		UserID:   user.Id,
		Current:  current,
		PageSize: pageSize,
	})
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{
		Object: httpdto.NewSharedAddressBookProfileListResponse(result),
	}
}

func sanitizeLegacyAddressBookForAudit(data api.AbData) map[string]any {
	peerIDs := make([]string, 0, len(data.Peers))
	for _, peer := range data.Peers {
		peerIDs = append(peerIDs, peer.Id)
	}
	return map[string]any{
		"tag_count":  len(data.Tags),
		"peer_count": len(data.Peers),
		"tags":       data.Tags,
		"peer_ids":   peerIDs,
	}
}
