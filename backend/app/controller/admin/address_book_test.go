package admin

import (
	"testing"

	"rustdesk-api-server-pro/app/model"
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
