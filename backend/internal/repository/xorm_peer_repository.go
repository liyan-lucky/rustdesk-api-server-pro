package repository

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"
	"strings"

	"xorm.io/xorm"
)

type XormPeerRepository struct {
	DB *xorm.Engine
}

func NewXormPeerRepository(dbEngine *xorm.Engine) *XormPeerRepository {
	return &XormPeerRepository{DB: dbEngine}
}

func (r *XormPeerRepository) ListByUser(query core.PeerListQuery) (core.PeerListResult, error) {
	if query.HasAccessibleParam {
		return r.listAccessible(query)
	}

	sessionQuery := func() *xorm.Session {
		q := r.DB.Table(&model.Peer{}).Where("user_id = ?", query.UserID)
		return q.Desc("id")
	}

	pagination := db.NewPagination(query.Current, query.PageSize)
	peerList := make([]model.Peer, 0)
	if err := pagination.Paginate(sessionQuery, &model.Peer{}, &peerList); err != nil {
		return core.PeerListResult{}, err
	}

	items := make([]core.PeerListItem, 0, len(peerList))
	for _, p := range peerList {
		items = append(items, core.PeerListItem{
			RustdeskID:      p.RustdeskId,
			Username:        p.Username,
			Platform:        p.Platform,
			Hostname:        p.Hostname,
			LoginName:       p.LoginName,
			OwnerName:       query.Username,
			DeviceGroupName: "",
			Note:            p.Note,
		})
	}

	return core.PeerListResult{
		Total: pagination.TotalCount,
		Items: items,
	}, nil
}

type accessiblePeerJoinRow struct {
	RustdeskId string `xorm:"rustdesk_id"`
	Username   string `xorm:"username"`
	Platform   string `xorm:"platform"`
	Hostname   string `xorm:"hostname"`
	LoginName  string `xorm:"loginName"`
	Note       string `xorm:"note"`
	OwnerName  string `xorm:"owner_name"`
	GroupName  string `xorm:"group_name"`
}

func (r *XormPeerRepository) listAccessible(query core.PeerListQuery) (core.PeerListResult, error) {
	if query.Current < 1 {
		query.Current = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	countQ := r.DB.Table("peer").Alias("p").
		Join("INNER", []string{"device_group_device", "dgd"}, "dgd.rustdesk_id = p.rustdesk_id").
		Join("INNER", []string{"device_group", "dg"}, "dg.guid = dgd.group_guid").
		Join("LEFT", []string{"user", "u"}, "u.id = p.user_id")
	listQ := r.DB.Table("peer").Alias("p").
		Join("INNER", []string{"device_group_device", "dgd"}, "dgd.rustdesk_id = p.rustdesk_id").
		Join("INNER", []string{"device_group", "dg"}, "dg.guid = dgd.group_guid").
		Join("LEFT", []string{"user", "u"}, "u.id = p.user_id")
	if !query.RequestUserIsAdmin {
		countQ = countQ.Where("dg.owner_id = ?", query.UserID)
		listQ = listQ.Where("dg.owner_id = ?", query.UserID)
	}

	total, err := countQ.Count(new(model.Peer))
	if err != nil {
		return core.PeerListResult{}, err
	}

	rows := make([]accessiblePeerJoinRow, 0)
	err = listQ.
		Select(strings.Join([]string{
			"p.rustdesk_id as rustdesk_id",
			"p.username as username",
			"p.platform as platform",
			"p.hostname as hostname",
			"p.loginName as loginName",
			"p.note as note",
			"u.name as owner_name",
			"dg.name as group_name",
		}, ",")).
		Desc("p.id").
		Limit(query.PageSize, (query.Current-1)*query.PageSize).
		Find(&rows)
	if err != nil {
		return core.PeerListResult{}, err
	}

	items := make([]core.PeerListItem, 0, len(rows))
	for _, p := range rows {
		items = append(items, core.PeerListItem{
			RustdeskID:      p.RustdeskId,
			Username:        p.Username,
			Platform:        p.Platform,
			Hostname:        p.Hostname,
			LoginName:       p.LoginName,
			OwnerName:       p.OwnerName,
			DeviceGroupName: p.GroupName,
			Note:            p.Note,
		})
	}
	return core.PeerListResult{Total: total, Items: items}, nil
}
