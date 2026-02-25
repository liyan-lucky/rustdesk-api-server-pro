package api

import (
	"encoding/json"
	"rustdesk-api-server-pro/app/form/api"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/transport/httpdto"
	"strconv"

	"github.com/beevik/guid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xorm.io/xorm"
)

type AddressBookController struct {
	basicController
}

func (c *AddressBookController) GetAb() mvc.Result {
	user := c.GetUser()
	tagList := make([]model.Tags, 0)
	err := c.Db.Where("user_id = ?", user.Id).Find(&tagList)
	if err != nil {
		return c.fail(err)
	}
	tags := make([]string, 0)
	tagColors := make(map[string]int64)
	for _, tag := range tagList {
		tags = append(tags, tag.Tag)
		colorCode, err := strconv.ParseInt(tag.Color, 10, 64)
		if err != nil {
			continue
		}
		tagColors[tag.Tag] = colorCode
	}

	peerList := make([]model.Peer, 0)
	err = c.Db.Where("user_id = ?", user.Id).Find(&peerList)
	if err != nil {
		return c.fail(err)
	}
	peers := make([]iris.Map, 0)
	for _, peer := range peerList {
		var peerTags []string
		err := json.Unmarshal([]byte(peer.Tags), &peerTags)
		if err != nil {
			continue
		}
		peers = append(peers, iris.Map{
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

	tagColorsJson, err := json.Marshal(tagColors)
	if err != nil {
		return c.fail(err)
	}

	dataJson, err := json.Marshal(iris.Map{
		"tags":       tags,
		"peers":      peers,
		"tag_colors": string(tagColorsJson),
	})
	if err != nil {
		return c.fail(err)
	}

	return mvc.Response{
		Object: iris.Map{
			"licensed_devices": user.LicensedDevices,
			"data":             string(dataJson),
		},
	}
}

func (c *AddressBookController) PostAb() mvc.Result {
	var abForm api.AbForm
	err := c.Ctx.ReadJSON(&abForm)
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

	err = c.withTx(func(session *xorm.Session) error {
		if _, err := session.Where("user_id = ?", user.Id).Delete(&model.Tags{}); err != nil {
			return err
		}
		if _, err := session.Where("user_id = ?", user.Id).Delete(&model.Peer{}); err != nil {
			return err
		}

		tags := make([]*model.Tags, 0)
		for _, tag := range abData.Tags {
			tags = append(tags, &model.Tags{
				UserId: user.Id,
				Tag:    tag,
				Color:  strconv.FormatInt(tagColors[tag], 10),
			})
		}
		if len(tags) > 0 {
			if _, err := session.Insert(tags); err != nil {
				return err
			}
		}

		peers := make([]*model.Peer, 0)
		for _, peer := range abData.Peers {
			peerTags := ""
			b, err := json.Marshal(peer.Tags)
			if err == nil {
				peerTags = string(b)
			}
			peers = append(peers, &model.Peer{
				UserId:     user.Id,
				RustdeskId: peer.Id,
				Hash:       peer.Hash,
				Username:   peer.Username,
				Hostname:   peer.Hostname,
				Platform:   peer.Platform,
				Alias:      peer.Alias,
				Tags:       peerTags,
				Note:       peer.Note,
			})
		}
		if len(peers) > 0 {
			if _, err := session.Insert(peers); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return c.fail(err)
	}

	return c.ok()
}

func (c *AddressBookController) PostAbPersonal() mvc.Result {
	user := c.GetUser()
	var ab model.AddressBook
	has, err := c.Db.Where("user_id = ?", user.Id).Get(&ab)
	if err != nil {
		return c.fail(err)
	}

	if !has {
		g := guid.New()
		ab.UserId = user.Id
		ab.Guid = g.String()
		ab.Name = model.PersonalAddressBookName
		ab.Owner = user.Username
		ab.MaxPeer = model.MaxPeer
		ab.Note = "default address book"
		ab.Rule = 3 // full control
		c.Db.Insert(&ab)
	}

	return mvc.Response{
		Object: iris.Map{
			"guid": ab.Guid,
		},
	}
}

func (c *AddressBookController) PostAbSettings() mvc.Result {
	user := c.GetUser()
	var ab model.AddressBook
	_, _ = c.Db.Where("user_id = ?", user.Id).Get(&ab)
	return mvc.Response{
		Object: iris.Map{
			"max_peer_one_ab": ab.MaxPeer,
		},
	}
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
