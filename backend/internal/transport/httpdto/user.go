package httpdto

import "rustdesk-api-server-pro/internal/core"

type UserResponse struct {
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Note        string `json:"note"`
	Status      int    `json:"status"`
	IsAdmin     bool   `json:"is_admin"`
}

type UserListResponse struct {
	Total int64          `json:"total"`
	Data  []UserResponse `json:"data"`
}

func NewUserResponse(u core.UserView) UserResponse {
	return UserResponse{
		Name:        u.Name,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		Note:        u.Note,
		Status:      u.Status,
		IsAdmin:     u.IsAdmin,
	}
}

func NewUserListResponse(result core.UserListResult) UserListResponse {
	data := make([]UserResponse, 0, len(result.Items))
	for _, u := range result.Items {
		data = append(data, NewUserResponse(u))
	}
	return UserListResponse{
		Total: result.Total,
		Data:  data,
	}
}
