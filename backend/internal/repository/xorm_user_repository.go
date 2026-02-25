package repository

import (
	"rustdesk-api-server-pro/app/model"
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
	if query.Current < 1 {
		query.Current = 1
	}
	if query.PageSize <= 0 {
		query.PageSize = 10
	}

	countQ := r.DB.Table(&model.User{}).Where("status = ?", query.Status)
	listQ := r.DB.Table(&model.User{}).Where("status = ?", query.Status)
	if query.HasAccessibleParam {
		if query.RequestUserIsAdmin {
			// Admin: keep all active users visible.
		} else {
			ids, err := r.listAccessibleUserIDs(query.RequestUserID)
			if err != nil {
				return core.UserListResult{}, err
			}
			if len(ids) == 0 {
				ids = []int{query.RequestUserID}
			}
			countQ = countQ.In("id", ids)
			listQ = listQ.In("id", ids)
		}
	}

	total, err := countQ.Count(&model.User{})
	if err != nil {
		return core.UserListResult{}, err
	}

	users := make([]model.User, 0)
	if err := listQ.Desc("id").Limit(query.PageSize, (query.Current-1)*query.PageSize).Find(&users); err != nil {
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
		Total: total,
		Items: items,
	}, nil
}

func (r *XormUserRepository) listAccessibleUserIDs(requestUserID int) ([]int, error) {
	ids := map[int]struct{}{requestUserID: {}}

	var memberships []model.UserGroupMember
	if err := r.DB.Where("user_id = ?", requestUserID).Find(&memberships); err != nil {
		return nil, err
	}
	if len(memberships) == 0 {
		return []int{requestUserID}, nil
	}

	groupGuids := make([]string, 0, len(memberships))
	for _, m := range memberships {
		if m.GroupGuid != "" {
			groupGuids = append(groupGuids, m.GroupGuid)
		}
	}
	if len(groupGuids) == 0 {
		return []int{requestUserID}, nil
	}

	var members []model.UserGroupMember
	if err := r.DB.In("group_guid", groupGuids).Find(&members); err != nil {
		return nil, err
	}
	for _, m := range members {
		ids[m.UserId] = struct{}{}
	}

	out := make([]int, 0, len(ids))
	for id := range ids {
		out = append(out, id)
	}
	return out, nil
}

func (r *XormUserRepository) ExpireAuthTokenByRustdeskID(rustdeskID string) error {
	_, err := r.DB.Where("rustdesk_id = ?", rustdeskID).Cols("status").Update(&model.AuthToken{
		Status: 0,
	})
	return err
}
