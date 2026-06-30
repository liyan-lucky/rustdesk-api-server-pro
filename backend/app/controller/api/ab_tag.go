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
		c.recordAPIOperationAudit("ab_tag_add", "address_book_tag", abGuid, nil, nil, "failure", err.Error())
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_add", "address_book_tag", abGuid, nil, sanitizeAddressBookTagFormForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	cmd := core.AddressBookTagAddCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Name:   form.Name,
		Color:  form.Color,
	}
	err = c.addressBookService().AddTag(cmd)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_add", "address_book_tag", abGuid, nil, sanitizeAddressBookTagFormForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	c.recordAPIOperationAudit("ab_tag_add", "address_book_tag", abGuid, nil, sanitizeAddressBookTagFormForAudit(form), "success", "")
	return c.okText("")
}

func (c *AddressBookTagController) HandleAbTagUpdate() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagForm
	err := c.readJSONBody(&form)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_update", "address_book_tag", abGuid, nil, nil, "failure", err.Error())
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_update", "address_book_tag", abGuid, nil, sanitizeAddressBookTagFormForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	cmd := core.AddressBookTagUpdateColorCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Name:   form.Name,
		Color:  form.Color,
	}
	err = c.addressBookService().UpdateTagColor(cmd)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_update", "address_book_tag", abGuid, nil, sanitizeAddressBookTagFormForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	c.recordAPIOperationAudit("ab_tag_update", "address_book_tag", abGuid, nil, sanitizeAddressBookTagFormForAudit(form), "success", "")
	return c.okText("")
}

func (c *AddressBookTagController) HandleAbTagRename() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var form api.AbTagRenameForm
	err := c.readJSONBody(&form)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_rename", "address_book_tag", abGuid, nil, nil, "failure", err.Error())
		return c.fail(err)
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_rename", "address_book_tag", abGuid, nil, sanitizeAddressBookTagRenameForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	cmd := core.AddressBookTagRenameCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Old:    form.Old,
		New:    form.New,
	}
	err = c.addressBookService().RenameTag(cmd)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_rename", "address_book_tag", abGuid, nil, sanitizeAddressBookTagRenameForAudit(form), "failure", err.Error())
		return c.fail(err)
	}
	c.recordAPIOperationAudit("ab_tag_rename", "address_book_tag", abGuid, sanitizeAddressBookTagRenameBeforeForAudit(form), sanitizeAddressBookTagRenameForAudit(form), "success", "")
	return c.okText("")
}

func (c *AddressBookTagController) HandleAbTagDelete() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var names []string
	err := c.Ctx.ReadBody(&names)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_delete", "address_book_tag", abGuid, nil, nil, "failure", err.Error())
		return c.fail(err)
	}
	if len(names) == 0 {
		c.recordAPIOperationAudit("ab_tag_delete", "address_book_tag", abGuid, nil, map[string]any{"names": names}, "failure", "NoTagNames")
		return c.failMsg("NoTagNames")
	}

	user := c.GetUser()
	ab, err := c.getAddressBookByGuid(user.Id, abGuid)
	if err != nil {
		c.recordAPIOperationAudit("ab_tag_delete", "address_book_tag", abGuid, map[string]any{"names": names}, nil, "failure", err.Error())
		return c.fail(err)
	}
	if err := c.addressBookService().DeleteTags(core.AddressBookTagDeleteCommand{
		UserID: user.Id,
		AbID:   ab.Id,
		Names:  names,
	}); err != nil {
		c.recordAPIOperationAudit("ab_tag_delete", "address_book_tag", abGuid, map[string]any{"names": names}, nil, "failure", err.Error())
		return c.fail(err)
	}

	c.recordAPIOperationAudit("ab_tag_delete", "address_book_tag", abGuid, map[string]any{"names": names}, nil, "success", "")
	return c.ok()
}

func sanitizeAddressBookTagFormForAudit(form api.AbTagForm) map[string]any {
	return map[string]any{
		"name":  form.Name,
		"color": form.Color,
	}
}

func sanitizeAddressBookTagRenameForAudit(form api.AbTagRenameForm) map[string]any {
	return map[string]any{
		"old": form.Old,
		"new": form.New,
	}
}

func sanitizeAddressBookTagRenameBeforeForAudit(form api.AbTagRenameForm) map[string]any {
	return map[string]any{
		"name": form.Old,
	}
}
