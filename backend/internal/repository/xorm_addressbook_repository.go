package repository

import (
	"encoding/json"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"

	"xorm.io/xorm"
)

type XormAddressBookRepository struct {
	DB *xorm.Engine
}

func NewXormAddressBookRepository(dbEngine *xorm.Engine) *XormAddressBookRepository {
	return &XormAddressBookRepository{DB: dbEngine}
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
