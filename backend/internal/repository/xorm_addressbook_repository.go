package repository

import (
	"encoding/json"
	"strconv"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"

	"github.com/beevik/guid"
	"xorm.io/xorm"
)

type XormAddressBookRepository struct {
	DB *xorm.Engine
}

func NewXormAddressBookRepository(dbEngine *xorm.Engine) *XormAddressBookRepository {
	return &XormAddressBookRepository{DB: dbEngine}
}

func (r *XormAddressBookRepository) GetLegacyAddressBook(query core.LegacyAddressBookGetQuery) (core.LegacyAddressBookGetResult, error) {
	tagList := make([]model.Tags, 0)
	if err := r.DB.Where("user_id = ?", query.UserID).Find(&tagList); err != nil {
		return core.LegacyAddressBookGetResult{}, err
	}

	tags := make([]string, 0, len(tagList))
	tagColors := make(map[string]int64, len(tagList))
	for _, tag := range tagList {
		tags = append(tags, tag.Tag)
		colorCode, err := strconv.ParseInt(tag.Color, 10, 64)
		if err != nil {
			continue
		}
		tagColors[tag.Tag] = colorCode
	}

	peerList := make([]model.Peer, 0)
	if err := r.DB.Where("user_id = ?", query.UserID).Find(&peerList); err != nil {
		return core.LegacyAddressBookGetResult{}, err
	}

	peers := make([]core.LegacyAddressBookPeerEntry, 0, len(peerList))
	for _, peer := range peerList {
		peerTags := make([]string, 0)
		if err := json.Unmarshal([]byte(peer.Tags), &peerTags); err != nil {
			continue
		}
		peers = append(peers, core.LegacyAddressBookPeerEntry{
			RustdeskID: peer.RustdeskId,
			Hash:       peer.Hash,
			Username:   peer.Username,
			Hostname:   peer.Hostname,
			Platform:   peer.Platform,
			Alias:      peer.Alias,
			Tags:       peerTags,
			Note:       peer.Note,
		})
	}

	return core.LegacyAddressBookGetResult{
		Tags:      tags,
		TagColors: tagColors,
		Peers:     peers,
	}, nil
}

func (r *XormAddressBookRepository) EnsurePersonalAddressBook(cmd core.PersonalAddressBookEnsureCommand) (core.PersonalAddressBookEnsureResult, error) {
	var ab model.AddressBook
	has, err := r.DB.Where("user_id = ?", cmd.UserID).Get(&ab)
	if err != nil {
		return core.PersonalAddressBookEnsureResult{}, err
	}
	if !has {
		g := guid.New()
		ab = model.AddressBook{
			UserId:  cmd.UserID,
			Guid:    g.String(),
			Name:    model.PersonalAddressBookName,
			Owner:   cmd.Username,
			MaxPeer: cmd.DefaultMaxPeer,
			Note:    cmd.DefaultNote,
			Rule:    cmd.DefaultRule,
		}
		if _, err := r.DB.Insert(&ab); err != nil {
			return core.PersonalAddressBookEnsureResult{}, err
		}
	}
	return core.PersonalAddressBookEnsureResult{Guid: ab.Guid}, nil
}

func (r *XormAddressBookRepository) GetAddressBookSettings(query core.AddressBookSettingsQuery) (core.AddressBookSettingsResult, error) {
	var ab model.AddressBook
	if _, err := r.DB.Where("user_id = ?", query.UserID).Get(&ab); err != nil {
		return core.AddressBookSettingsResult{}, err
	}
	return core.AddressBookSettingsResult{MaxPeerOneAB: ab.MaxPeer}, nil
}

func (r *XormAddressBookRepository) ListAddressBookPeers(query core.AddressBookPeerListQuery) (core.AddressBookPeerListResult, error) {
	var ab model.AddressBook
	_, err := r.DB.Where("user_id = ? and guid = ?", query.UserID, query.AbGuid).Get(&ab)
	if err != nil {
		return core.AddressBookPeerListResult{}, err
	}

	sessionQuery := func() *xorm.Session {
		return r.DB.Table(&model.Peer{}).Where("user_id = ? and ab_id = ?", query.UserID, ab.Id)
	}
	pagination := db.NewPagination(query.Current, query.PageSize)
	peerList := make([]model.Peer, 0)
	if err := pagination.Paginate(sessionQuery, &model.Peer{}, &peerList); err != nil {
		return core.AddressBookPeerListResult{}, err
	}

	items := make([]core.AddressBookPeerView, 0, len(peerList))
	for _, peer := range peerList {
		peerTags := make([]string, 0)
		if err := json.Unmarshal([]byte(peer.Tags), &peerTags); err != nil {
			continue
		}
		items = append(items, core.AddressBookPeerView{
			ID:               peer.RustdeskId,
			Hash:             peer.Hash,
			Password:         peer.Password,
			Username:         peer.Username,
			Hostname:         peer.Hostname,
			Platform:         peer.Platform,
			Alias:            peer.Alias,
			Tags:             peerTags,
			Note:             peer.Note,
			ForceAlwaysRelay: peer.ForceAlwaysRelay,
			RdpPort:          peer.RdpPort,
			RdpUsername:      peer.RdpUsername,
			LoginName:        peer.LoginName,
			SameServer:       peer.SameServer,
		})
	}
	return core.AddressBookPeerListResult{Total: pagination.TotalCount, Items: items}, nil
}

func (r *XormAddressBookRepository) ListAddressBookTags(query core.AddressBookTagListQuery) ([]core.AddressBookTagView, error) {
	var ab model.AddressBook
	_, err := r.DB.Where("user_id = ? and guid = ?", query.UserID, query.AbGuid).Get(&ab)
	if err != nil {
		return nil, err
	}
	tags := make([]model.AddressBookTag, 0)
	if err := r.DB.Where("user_id = ? and ab_id = ?", query.UserID, ab.Id).Find(&tags); err != nil {
		return nil, err
	}
	out := make([]core.AddressBookTagView, 0, len(tags))
	for _, t := range tags {
		out = append(out, core.AddressBookTagView{Name: t.Name, Color: t.Color})
	}
	return out, nil
}

func (r *XormAddressBookRepository) ListSharedAddressBooks(query core.SharedAddressBookListQuery) (core.SharedAddressBookListResult, error) {
	sessionQuery := func() *xorm.Session {
		return r.DB.Table(&model.AddressBook{}).Where("shared = 1")
	}
	pagination := db.NewPagination(query.Current, query.PageSize)
	list := make([]model.AddressBook, 0)
	if err := pagination.Paginate(sessionQuery, &model.AddressBook{}, &list); err != nil {
		return core.SharedAddressBookListResult{}, err
	}
	items := make([]core.SharedAddressBookView, 0, len(list))
	for _, ab := range list {
		items = append(items, core.SharedAddressBookView{
			Guid:  ab.Guid,
			Name:  ab.Name,
			Owner: ab.Owner,
			Note:  ab.Note,
			Rule:  ab.Rule,
		})
	}
	return core.SharedAddressBookListResult{Total: pagination.TotalCount, Items: items}, nil
}

func (r *XormAddressBookRepository) AddAddressBookTag(cmd core.AddressBookTagAddCommand) error {
	_, err := r.DB.Insert(&model.AddressBookTag{
		UserId: cmd.UserID,
		AbId:   cmd.AbID,
		Name:   cmd.Name,
		Color:  cmd.Color,
	})
	return err
}

func (r *XormAddressBookRepository) UpdateAddressBookTagColor(cmd core.AddressBookTagUpdateColorCommand) error {
	_, err := r.DB.Where("user_id = ? and ab_id = ? and name = ?", cmd.UserID, cmd.AbID, cmd.Name).
		Update(&model.AddressBookTag{Color: cmd.Color})
	return err
}

func (r *XormAddressBookRepository) RenameAddressBookTag(cmd core.AddressBookTagRenameCommand) error {
	_, err := r.DB.Where("user_id = ? and ab_id = ? and name = ?", cmd.UserID, cmd.AbID, cmd.Old).
		Update(&model.AddressBookTag{Name: cmd.New})
	return err
}

func (r *XormAddressBookRepository) DeleteAddressBookTags(cmd core.AddressBookTagDeleteCommand) error {
	_, err := r.DB.Where("user_id = ? and ab_id = ?", cmd.UserID, cmd.AbID).
		In("name", cmd.Names).
		Delete(&model.AddressBookTag{})
	return err
}

func (r *XormAddressBookRepository) CountAddressBookPeers(userID, abID int) (int64, error) {
	return r.DB.Where("user_id = ? and ab_id = ?", userID, abID).Count(&model.Peer{})
}

func (r *XormAddressBookRepository) AddAddressBookPeer(cmd core.AddressBookPeerCreateCommand) error {
	peerTags, err := json.Marshal(cmd.Tags)
	if err != nil {
		return err
	}
	tagsJSON := string(peerTags)
	if tagsJSON == "null" {
		tagsJSON = "[]"
	}
	_, err = r.DB.Insert(&model.Peer{
		UserId:           cmd.UserID,
		AbId:             cmd.AbID,
		RustdeskId:       cmd.RustdeskID,
		Hash:             cmd.Hash,
		Username:         cmd.Username,
		Password:         cmd.Password,
		Hostname:         cmd.Hostname,
		Platform:         cmd.Platform,
		Alias:            cmd.Alias,
		Tags:             tagsJSON,
		Note:             cmd.Note,
		ForceAlwaysRelay: cmd.ForceAlwaysRelay,
		RdpPort:          cmd.RdpPort,
		RdpUsername:      cmd.RdpUsername,
		LoginName:        cmd.LoginName,
		SameServer:       cmd.SameServerPresent,
	})
	return err
}

func (r *XormAddressBookRepository) UpdateAddressBookPeer(cmd core.AddressBookPeerUpdateCommand) (bool, error) {
	var peer model.Peer
	has, err := r.DB.Where("user_id = ? and ab_id = ? and rustdesk_id = ?", cmd.UserID, cmd.AbID, cmd.RustdeskID).Get(&peer)
	if err != nil || !has {
		return has, err
	}

	updateCols := make([]string, 0, 8)
	if cmd.Tags != nil {
		peer.Tags = *cmd.Tags
		updateCols = append(updateCols, "tags")
	}
	if cmd.Alias != nil {
		peer.Alias = *cmd.Alias
		updateCols = append(updateCols, "alias")
	}
	if cmd.Hash != nil {
		peer.Hash = *cmd.Hash
		updateCols = append(updateCols, "hash")
	}
	if cmd.Password != nil {
		peer.Password = *cmd.Password
		updateCols = append(updateCols, "password")
	}
	if cmd.Note != nil {
		peer.Note = *cmd.Note
		updateCols = append(updateCols, "note")
	}
	if cmd.Username != nil {
		peer.Username = *cmd.Username
		updateCols = append(updateCols, "username")
	}
	if cmd.Hostname != nil {
		peer.Hostname = *cmd.Hostname
		updateCols = append(updateCols, "hostname")
	}
	if cmd.Platform != nil {
		peer.Platform = *cmd.Platform
		updateCols = append(updateCols, "platform")
	}

	if len(updateCols) == 0 {
		return true, nil
	}
	_, err = r.DB.Where("id = ?", peer.Id).Cols(updateCols...).Update(&peer)
	return true, err
}

func (r *XormAddressBookRepository) DeleteAddressBookPeers(cmd core.AddressBookPeerDeleteCommand) error {
	_, err := r.DB.Where("user_id = ? and ab_id = ?", cmd.UserID, cmd.AbID).
		In("rustdesk_id", cmd.IDs).
		Delete(&model.Peer{})
	return err
}

func (r *XormAddressBookRepository) ReplaceLegacyAddressBookData(cmd core.LegacyAddressBookReplaceCommand) error {
	session := r.DB.NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		return err
	}

	if _, err := session.Where("user_id = ?", cmd.UserID).Delete(&model.Tags{}); err != nil {
		_ = session.Rollback()
		return err
	}
	if _, err := session.Where("user_id = ?", cmd.UserID).Delete(&model.Peer{}); err != nil {
		_ = session.Rollback()
		return err
	}

	tags := make([]*model.Tags, 0, len(cmd.Tags))
	for _, tag := range cmd.Tags {
		tags = append(tags, &model.Tags{
			UserId: cmd.UserID,
			Tag:    tag.Name,
			Color:  strconv.FormatInt(tag.Color, 10),
		})
	}
	if len(tags) > 0 {
		if _, err := session.Insert(tags); err != nil {
			_ = session.Rollback()
			return err
		}
	}

	peers := make([]*model.Peer, 0, len(cmd.Peers))
	for _, peer := range cmd.Peers {
		peerTags := "[]"
		if b, err := json.Marshal(peer.Tags); err == nil {
			peerTags = string(b)
		}
		peers = append(peers, &model.Peer{
			UserId:     cmd.UserID,
			RustdeskId: peer.RustdeskID,
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
			_ = session.Rollback()
			return err
		}
	}

	return session.Commit()
}
