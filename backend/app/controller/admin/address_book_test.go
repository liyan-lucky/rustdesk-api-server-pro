package admin

import (
	"testing"

	"rustdesk-api-server-pro/app/model"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func TestCanDeleteAddressBook(t *testing.T) {
	admin := &model.User{IsAdmin: true}
	user := &model.User{}
	adminCreated := &model.AddressBook{CreatedByAdmin: true}
	userCreated := &model.AddressBook{}

	if !canDeleteAddressBook(admin, adminCreated) {
		t.Fatal("administrator must be able to delete an administrator-created address book")
	}
	if canDeleteAddressBook(user, adminCreated) {
		t.Fatal("ordinary user must not delete an administrator-created address book")
	}
	if !canDeleteAddressBook(user, userCreated) {
		t.Fatal("ordinary user must be able to delete their own address book")
	}
	if canDeleteAddressBook(nil, userCreated) {
		t.Fatal("unauthenticated caller must not delete an address book")
	}
}

func TestFindManageableAddressBooksIncludesWritableSharedBooks(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	if err = engine.Sync(new(model.AddressBook), new(model.AddressBookRule), new(model.UserGroupMember)); err != nil {
		t.Fatal(err)
	}
	books := []model.AddressBook{
		{UserId: 2, Guid: "owned", Name: "Owned"},
		{UserId: 1, Guid: "global-write", Name: "Global write", Shared: true, Rule: 2},
		{UserId: 1, Guid: "global-read", Name: "Global read", Shared: true, Rule: 1},
		{UserId: 1, Guid: "user-write", Name: "User write", Shared: true},
	}
	if _, err = engine.Insert(&books); err != nil {
		t.Fatal(err)
	}
	if _, err = engine.Insert(&model.AddressBookRule{Guid: "rule-user", AbGuid: "user-write", TargetType: "user", TargetGuid: "2", Rule: 2}); err != nil {
		t.Fatal(err)
	}
	list, err := findManageableAddressBooks(engine, &model.User{Id: 2})
	if err != nil {
		t.Fatal(err)
	}
	guids := make(map[string]bool, len(list))
	for _, book := range list {
		guids[book.Guid] = true
	}
	if !guids["owned"] || !guids["global-write"] || !guids["user-write"] || guids["global-read"] {
		t.Fatalf("unexpected manageable books: %+v", guids)
	}
}
