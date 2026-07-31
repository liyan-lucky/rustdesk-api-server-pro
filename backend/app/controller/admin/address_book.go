package admin

import (
	"encoding/json"
	"errors"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/repository"
	v2service "rustdesk-api-server-pro/internal/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

var errAddressBookNotFound = errors.New("address book not found")

type AddressBookController struct {
	basicController
}

func (c *AddressBookController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/ab/peers", "HandleAbPeers")
	b.Handle("GET", "/ab/shared-profiles", "HandleAbSharedProfiles")
	b.Handle("POST", "/ab/tags/{guid:string}", "HandleAbTags")
	b.Handle("DELETE", "/ab/peer/{guid:string}", "HandleAbPeerDelete")
	b.Handle("DELETE", "/ab/tag/{guid:string}", "HandleAbTagDelete")
}

func (c *AddressBookController) addressBookService() *v2service.AddressBookService {
	return v2service.NewAddressBookService(repository.NewXormAddressBookRepository(c.Db))
}

func (c *AddressBookController) HandleAbPeers() mvc.Result {
	current := c.Ctx.URLParamIntDefault("current", 1)
	pageSize := c.Ctx.URLParamIntDefault("size", 10)
	abGuid := c.Ctx.URLParamDefault("ab", "")

	user := c.GetUser()

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
		return c.Error(nil, err.Error())
	}

	return c.Success(iris.Map{
		"total":   result.Total,
		"records": peersToRecords(result.Items),
		"current": current,
		"size":    pageSize,
	}, "ok")
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
		return c.Error(nil, err.Error())
	}

	records := make([]iris.Map, 0, len(peerList))
	for _, peer := range peerList {
		peerTags := make([]string, 0)
		_ = json.Unmarshal([]byte(peer.Tags), &peerTags)
		records = append(records, iris.Map{
			"id":         peer.RustdeskId,
			"hash":       peer.Hash,
			"username":   peer.Username,
			"hostname":   peer.Hostname,
			"platform":   peer.Platform,
			"alias":      peer.Alias,
			"tags":       peerTags,
			"note":       peer.Note,
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

	result, err := c.addressBookService().ListSharedProfiles(core.SharedAddressBookListQuery{
		Current:  current,
		PageSize: pageSize,
	})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	records := make([]iris.Map, 0, len(result.Items))
	for _, ab := range result.Items {
		records = append(records, iris.Map{
			"guid":  ab.Guid,
			"name":  ab.Name,
			"owner": ab.Owner,
			"note":  ab.Note,
			"rule":  ab.Rule,
		})
	}

	return c.Success(iris.Map{
		"total":   result.Total,
		"records": records,
		"current": current,
		"size":    pageSize,
	}, "ok")
}

func (c *AddressBookController) HandleAbTags() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")
	user := c.GetUser()

	if user.IsAdmin {
		return c.listAllTags(abGuid)
	}

	tags, err := c.addressBookService().ListTags(core.AddressBookTagListQuery{
		UserID: user.Id,
		AbGuid: abGuid,
	})
	if err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(tagsToData(tags), "ok")
}

func (c *AddressBookController) listAllTags(abGuid string) mvc.Result {
	var ab model.AddressBook
	has, err := c.Db.Where("guid = ?", abGuid).Get(&ab)
	if err != nil {
		return c.Error(nil, err.Error())
	}
	if !has {
		return c.Error(nil, errAddressBookNotFound.Error())
	}

	tags := make([]model.AddressBookTag, 0)
	if err := c.Db.Where("ab_id = ?", ab.Id).Find(&tags); err != nil {
		return c.Error(nil, err.Error())
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
		return c.Error(nil, err.Error())
	}
	if len(ids) == 0 {
		return c.Error(nil, "NoPeerIds")
	}

	user := c.GetUser()
	ab, err := c.findAddressBook(user, abGuid)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if err := c.addressBookService().DeletePeers(core.AddressBookPeerDeleteCommand{
		UserID: ab.UserId,
		AbID:   ab.Id,
		IDs:    ids,
	}); err != nil {
		return c.Error(nil, err.Error())
	}

	return c.Success(nil, "ok")
}

func (c *AddressBookController) HandleAbTagDelete() mvc.Result {
	abGuid := c.Ctx.Params().Get("guid")

	var names []string
	if err := c.Ctx.ReadBody(&names); err != nil {
		return c.Error(nil, err.Error())
	}
	if len(names) == 0 {
		return c.Error(nil, "NoTagNames")
	}

	user := c.GetUser()
	ab, err := c.findAddressBook(user, abGuid)
	if err != nil {
		return c.Error(nil, err.Error())
	}

	if err := c.addressBookService().DeleteTags(core.AddressBookTagDeleteCommand{
		UserID: ab.UserId,
		AbID:   ab.Id,
		Names:  names,
	}); err != nil {
		return c.Error(nil, err.Error())
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
