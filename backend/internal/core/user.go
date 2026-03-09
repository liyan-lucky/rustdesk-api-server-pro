package core

type UserView struct {
	Name        string
	DisplayName string
	Email       string
	Note        string
	Status      int
	IsAdmin     bool
}

type UserListQuery struct {
	RequestUserID      int
	RequestUserIsAdmin bool
	HasAccessibleParam bool
	Current            int
	PageSize           int
	Status             int
}

type UserListResult struct {
	Total int64
	Items []UserView
}
