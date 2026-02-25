package service

import (
	"errors"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/repository"
)

var ErrAdminRequired = errors.New("Admin required!")

type UserService struct {
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) *UserService {
	return &UserService{users: users}
}

func (s *UserService) CurrentUserView(name, email, note string, status int, isAdmin bool) core.UserView {
	return core.UserView{
		Name:        name,
		DisplayName: name,
		Email:       email,
		Note:        note,
		Status:      status,
		IsAdmin:     isAdmin,
	}
}

func (s *UserService) ListUsers(query core.UserListQuery) (core.UserListResult, error) {
	if !query.RequestUserIsAdmin && !query.HasAccessibleParam {
		return core.UserListResult{}, ErrAdminRequired
	}
	return s.users.List(query)
}

func (s *UserService) LogoutByRustdeskID(rustdeskID string) error {
	return s.users.ExpireAuthTokenByRustdeskID(rustdeskID)
}
