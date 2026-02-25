package api

import (
	"github.com/kataras/iris/v12/mvc"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/transport/httpdto"
)

type AddressBookTagController struct {
	basicController
}

func (c *AddressBookTagController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "ab/tags/{guid:string}", "HandleAbTags")
	b.Handle("POST", "ab/tag/add/{guid:string}", "HandleAbTagAdd")
	b.Handle("PUT", "ab/tag/update/{guid:string}", "HandleAbTagUpdate")
	b.Handle("PUT", "ab/tag/rename/{guid:string}", "HandleAbTagRename")
	b.Handle("DELETE", "ab/tag/{guid:string}", "HandleAbTagDelete")
}

func (c *AddressBookTagController) HandleAbTags() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	user := c.GetUser()
	tags, err := c.addressBookService().ListTags(core.AddressBookTagListQuery{
		UserID: user.Id,
		AbGuid: abGuid,
	})
	if err != nil {
		return c.fail(err)
	}
	return mvc.Response{
		Object: httpdto.NewAddressBookTagListResponse(tags),
	}
}

func (c *AddressBookTagController) HandleAbTagAdd() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagForm
	err := c.readJSONBody(&form)
	if err != nil {
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		return c.fail(err)
	}
	err = c.addressBookService().AddTag(core.AddressBookTagAddCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Name:   form.Name,
		Color:  form.Color,
	})
	if err != nil {
		return c.fail(err)
	}
	return c.okText("")
}

func (c *AddressBookTagController) HandleAbTagUpdate() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagForm
	err := c.readJSONBody(&form)
	if err != nil {
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		return c.fail(err)
	}
	err = c.addressBookService().UpdateTagColor(core.AddressBookTagUpdateColorCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Name:   form.Name,
		Color:  form.Color,
	})
	if err != nil {
		return c.fail(err)
	}
	return c.okText("")
}

func (c *AddressBookTagController) HandleAbTagRename() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagRenameForm
	err := c.readJSONBody(&form)
	if err != nil {
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		return c.fail(err)
	}
	err = c.addressBookService().RenameTag(core.AddressBookTagRenameCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Old:    form.Old,
		New:    form.New,
	})
	if err != nil {
		return c.fail(err)
	}
	return c.okText("")
}

func (c *AddressBookTagController) HandleAbTagDelete() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var names []string
	err := c.Ctx.ReadBody(&names)
	if err != nil {
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		return c.fail(err)
	}
	if err := c.addressBookService().DeleteTags(core.AddressBookTagDeleteCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Names:  names,
	}); err != nil {
		return c.fail(err)
	}

	return c.ok()
}
