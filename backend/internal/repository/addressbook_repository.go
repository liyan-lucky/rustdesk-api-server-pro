package repository

import "rustdesk-api-server-pro/internal/core"

type AddressBookRepository interface {
	GetLegacyAddressBook(query core.LegacyAddressBookGetQuery) (core.LegacyAddressBookGetResult, error)
	EnsurePersonalAddressBook(cmd core.PersonalAddressBookEnsureCommand) (core.PersonalAddressBookEnsureResult, error)
	GetAddressBookSettings(query core.AddressBookSettingsQuery) (core.AddressBookSettingsResult, error)
	ListAddressBookPeers(query core.AddressBookPeerListQuery) (core.AddressBookPeerListResult, error)
	ListAddressBookTags(query core.AddressBookTagListQuery) ([]core.AddressBookTagView, error)
	ListSharedAddressBooks(query core.SharedAddressBookListQuery) (core.SharedAddressBookListResult, error)

	AddAddressBookTag(cmd core.AddressBookTagAddCommand) error
	UpdateAddressBookTagColor(cmd core.AddressBookTagUpdateColorCommand) error
	RenameAddressBookTag(cmd core.AddressBookTagRenameCommand) error
	DeleteAddressBookTags(cmd core.AddressBookTagDeleteCommand) error

	CountAddressBookPeers(userID, abID int) (int64, error)
	AddAddressBookPeer(cmd core.AddressBookPeerCreateCommand) error
	UpdateAddressBookPeer(cmd core.AddressBookPeerUpdateCommand) (bool, error)
	DeleteAddressBookPeers(cmd core.AddressBookPeerDeleteCommand) error

	ReplaceLegacyAddressBookData(cmd core.LegacyAddressBookReplaceCommand) error
}
