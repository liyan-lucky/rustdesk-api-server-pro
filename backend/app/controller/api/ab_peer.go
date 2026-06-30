package api

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/tidwall/gjson"

	abform "rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/transport/httpdto"
)

type AddressBookPeerController struct {
	basicController
}

func (c *AddressBookPeerController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "ab/peer/add/{guid:string}", "HandleAbPeerAdd")
	b.Handle("PUT", "ab/peer/update/{guid:string}", "HandleAbPeerUpdate")
	b.Handle("DELETE", "ab/peer/{guid:string}", "HandleAbPeerDelete")
}

func (c *AddressBookPeerController) PostAbPeers() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	abGuid := c.Ctx.URLParamDefault("ab", "")

	user := c.GetUser()
	result, err := c.addressBookService().ListPeers(core.AddressBookPeerListQuery{
		UserID:   user.Id,
		AbGuid:   abGuid,
		Current:  current,
		PageSize: pageSize,
	})
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: httpdto.NewAddressBookPeerListResponse(result)}
}

func (c *AddressBookPeerController) HandleAbPeerAdd() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form abform.AbPeer
	if err := c.readJSONBody(&form); err != nil {
		c.recordAPIOperationAudit("ab_peer_add", "address_book_peer", abGuid, nil, nil, "failure", err.Error())
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		c.recordAPIOperationAudit("ab_peer_add", "address_book_peer", abGuid, nil, sanitizeAddressBookPeerForAudit(form), "failure", err.Error())
		return c.fail(err)
	}

	totalPeers, err := c.addressBookService().CountPeers(user.Id, ab.Id)
	if err != nil {
		c.recordAPIOperationAudit("ab_peer_add", "address_book_peer", abGuid, nil, sanitizeAddressBookPeerForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	if ab.MaxPeer > 0 && totalPeers >= int64(ab.MaxPeer) {
		c.recordAPIOperationAudit("ab_peer_add", "address_book_peer", abGuid, nil, sanitizeAddressBookPeerForAudit(form), "failure", "exceed_max_devices")
		return c.failMsg("exceed_max_devices")
	}

	forceAlwaysRelay := form.ForceAlwaysRelay == "true"
	sameServerPresent := form.SameServer != ""

	if err := c.addressBookService().AddPeer(core.AddressBookPeerCreateCommand{
		UserID:            user.Id,
		AbID:              ab.Id,
		RustdeskID:        form.Id,
		Hash:              form.Hash,
		Username:          form.Username,
		Password:          form.Password,
		Hostname:          form.Hostname,
		Platform:          form.Platform,
		Alias:             form.Alias,
		Tags:              form.Tags,
		Note:              form.Note,
		ForceAlwaysRelay:  forceAlwaysRelay,
		RdpPort:           form.RdpPort,
		RdpUsername:       form.RdpUsername,
		LoginName:         form.LoginName,
		SameServerPresent: sameServerPresent,
	}); err != nil {
		c.recordAPIOperationAudit("ab_peer_add", "address_book_peer", form.Id, nil, sanitizeAddressBookPeerForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	c.recordAPIOperationAudit("ab_peer_add", "address_book_peer", form.Id, nil, sanitizeAddressBookPeerForAudit(form), "success", "")
	return c.okText("")
}

func (c *AddressBookPeerController) HandleAbPeerUpdate() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	body, err := c.readBodyBytes()
	if err != nil {
		c.recordAPIOperationAudit("ab_peer_update", "address_book_peer", abGuid, nil, nil, "failure", err.Error())
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		c.recordAPIOperationAudit("ab_peer_update", "address_book_peer", abGuid, nil, sanitizeAddressBookPeerUpdateBodyForAudit(body), "failure", err.Error())
		return c.fail(err)
	}

	cmd := parseAddressBookPeerUpdate(body)
	cmd.UserID = user.Id
	cmd.AbID = ab.Id

	has, err := c.addressBookService().UpdatePeer(cmd)
	if err != nil {
		c.recordAPIOperationAudit("ab_peer_update", "address_book_peer", cmd.RustdeskID, nil, sanitizeAddressBookPeerUpdateBodyForAudit(body), "failure", err.Error())
		return c.fail(err)
	}
	if !has {
		c.recordAPIOperationAudit("ab_peer_update", "address_book_peer", cmd.RustdeskID, nil, sanitizeAddressBookPeerUpdateBodyForAudit(body), "failure", "peer not found")
		return c.failMsg("peer not found")
	}
	c.recordAPIOperationAudit("ab_peer_update", "address_book_peer", cmd.RustdeskID, nil, sanitizeAddressBookPeerUpdateBodyForAudit(body), "success", "")
	return c.okText("")
}

func (c *AddressBookPeerController) HandleAbPeerDelete() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var ids []string
	if err := c.Ctx.ReadBody(&ids); err != nil {
		c.recordAPIOperationAudit("ab_peer_delete", "address_book_peer", abGuid, nil, nil, "failure", err.Error())
		return c.fail(err)
	}
	if len(ids) == 0 {
		c.recordAPIOperationAudit("ab_peer_delete", "address_book_peer", abGuid, map[string]any{"ids": ids}, nil, "failure", "NoPeerIds")
		return c.failMsg("NoPeerIds")
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		c.recordAPIOperationAudit("ab_peer_delete", "address_book_peer", abGuid, map[string]any{"ids": ids}, nil, "failure", err.Error())
		return c.fail(err)
	}

	if err := c.addressBookService().DeletePeers(core.AddressBookPeerDeleteCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		IDs:    ids,
	}); err != nil {
		c.recordAPIOperationAudit("ab_peer_delete", "address_book_peer", abGuid, map[string]any{"ids": ids}, nil, "failure", err.Error())
		return c.fail(err)
	}

	c.recordAPIOperationAudit("ab_peer_delete", "address_book_peer", abGuid, map[string]any{"ids": ids}, nil, "success", "")
	return c.ok()
}

func parseAddressBookPeerUpdate(body []byte) core.AddressBookPeerUpdateCommand {
	cmd := core.AddressBookPeerUpdateCommand{
		RustdeskID: gjson.GetBytes(body, "id").String(),
	}

	if v := gjson.GetBytes(body, "tags"); v.Exists() {
		tags := v.String()
		if tags == "null" {
			tags = "[]"
		}
		cmd.Tags = &tags
	}
	if v := gjson.GetBytes(body, "alias"); v.Exists() {
		value := v.String()
		cmd.Alias = &value
	}
	if v := gjson.GetBytes(body, "hash"); v.Exists() {
		value := v.String()
		cmd.Hash = &value
	}
	if v := gjson.GetBytes(body, "password"); v.Exists() {
		value := v.String()
		cmd.Password = &value
	}
	if v := gjson.GetBytes(body, "note"); v.Exists() {
		value := v.String()
		cmd.Note = &value
	}
	if v := gjson.GetBytes(body, "username"); v.Exists() {
		value := v.String()
		cmd.Username = &value
	}
	if v := gjson.GetBytes(body, "hostname"); v.Exists() {
		value := v.String()
		cmd.Hostname = &value
	}
	if v := gjson.GetBytes(body, "platform"); v.Exists() {
		value := v.String()
		cmd.Platform = &value
	}

	return cmd
}

func sanitizeAddressBookPeerForAudit(form abform.AbPeer) map[string]any {
	return map[string]any{
		"id":                 form.Id,
		"hash":               form.Hash,
		"username":           form.Username,
		"password_changed":   form.Password != "",
		"hostname":           form.Hostname,
		"platform":           form.Platform,
		"alias":              form.Alias,
		"tags":               form.Tags,
		"note":               form.Note,
		"force_always_relay": form.ForceAlwaysRelay,
		"rdp_port":           form.RdpPort,
		"rdp_username":       form.RdpUsername,
		"login_name":         form.LoginName,
		"same_server":        form.SameServer != "",
	}
}

func sanitizeAddressBookPeerUpdateBodyForAudit(body []byte) map[string]any {
	result := map[string]any{
		"id": gjson.GetBytes(body, "id").String(),
	}
	for _, key := range []string{"tags", "alias", "hash", "note", "username", "hostname", "platform"} {
		if v := gjson.GetBytes(body, key); v.Exists() {
			result[key] = v.Value()
		}
	}
	if v := gjson.GetBytes(body, "password"); v.Exists() {
		result["password_changed"] = v.String() != ""
	}
	return result
}
