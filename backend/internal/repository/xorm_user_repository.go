package repository

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/db"
	"rustdesk-api-server-pro/internal/core"

	"xorm.io/xorm"
)

type XormUserRepository struct {
	DB *xorm.Engine
}

func NewXormUserRepository(dbEngine *xorm.Engine) *XormUserRepository {
	return &XormUserRepository{DB: dbEngine}
}

func (r *XormUserRepository) List(query core.UserListQuery) (core.UserListResult, error) {
	sessionQuery := func() *xorm.Session {
		q := r.DB.Table(&model.User{}).Where("status = ?", query.Status)
		if query.HasAccessibleParam && !query.RequestUserIsAdmin {
			q.Where("id = ?", query.RequestUserID)
		}
		return q.Desc("id")
	}

	pagination := db.NewPagination(query.Current, query.PageSize)
	users := make([]model.User, 0)
	if err := pagination.Paginate(sessionQuery, &model.User{}, &users); err != nil {
		return core.UserListResult{}, err
	}

	items := make([]core.UserView, 0, len(users))
	for _, u := range users {
		items = append(items, core.UserView{
			Name:        u.Name,
			DisplayName: u.Name,
			Email:       u.Email,
			Note:        u.Note,
			Status:      u.Status,
			IsAdmin:     u.IsAdmin,
		})
	}
	return core.UserListResult{
		Total: pagination.TotalCount,
		Items: items,
	}, nil
}

func (r *XormUserRepository) ExpireAuthTokenByRustdeskID(rustdeskID string) error {
	_, err := r.DB.Where("rustdesk_id = ?", rustdeskID).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})
	return err
}
