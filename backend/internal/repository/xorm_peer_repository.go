package repository

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"

	"xorm.io/xorm"
)

type XormPeerRepository struct {
	DB *xorm.Engine
}

func NewXormPeerRepository(dbEngine *xorm.Engine) *XormPeerRepository {
	return &XormPeerRepository{DB: dbEngine}
}

func (r *XormPeerRepository) ListByUser(query core.PeerListQuery) (core.PeerListResult, error) {
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
			DeviceGroupName: "",
			Note:            p.Note,
		})
	}

	return core.PeerListResult{
		Total: pagination.TotalCount,
		Items: items,
	}, nil
}
