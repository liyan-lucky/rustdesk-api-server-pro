package api

import (
	"encoding/json"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/transport/httpdto"

	"github.com/kataras/iris/v12/mvc"
)

type AddressBookController struct {
	basicController
}

func (c *AddressBookController) GetAb() mvc.Result {
	user := c.GetUser()
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
		return c.fail(err)
	}
	var abData api.AbData
	err = json.Unmarshal([]byte(abForm.Data), &abData)
	if err != nil {
		return c.fail(err)
	}
	var tagColors map[string]int64
	err = json.Unmarshal([]byte(abData.TagColors), &tagColors)
	if err != nil {
		return c.fail(err)
	}

	user := c.GetUser()
	if user.LicensedDevices > 0 && len(abData.Peers) > user.LicensedDevices {
		return c.failMsg("Number of equipment in excess of licenses")
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
	err = c.addressBookService().ReplaceLegacyAddressBookData(cmd)
	if err != nil {
		return c.fail(err)
	}

	return c.ok()
}

func (c *AddressBookController) PostAbPersonal() mvc.Result {
	user := c.GetUser()
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

func (c *AddressBookController) PostAbSettings() mvc.Result {
	user := c.GetUser()
	result, err := c.addressBookService().GetSettings(core.AddressBookSettingsQuery{UserID: user.Id})
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{Object: httpdto.NewAddressBookSettingsResponse(result)}
}

func (c *AddressBookController) PostAbSharedProfiles() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("pageSize", 10)
	result, err := c.addressBookService().ListSharedProfiles(core.SharedAddressBookListQuery{
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
