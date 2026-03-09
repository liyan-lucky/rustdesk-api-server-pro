package service

import (
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/repository"
)

type AddressBookService struct {
	repo repository.AddressBookRepository
}

func NewAddressBookService(repo repository.AddressBookRepository) *AddressBookService {
	return &AddressBookService{repo: repo}
}

func (s *AddressBookService) GetLegacyAddressBook(query core.LegacyAddressBookGetQuery) (core.LegacyAddressBookGetResult, error) {
	return s.repo.GetLegacyAddressBook(query)
}

func (s *AddressBookService) EnsurePersonalAddressBook(cmd core.PersonalAddressBookEnsureCommand) (core.PersonalAddressBookEnsureResult, error) {
	return s.repo.EnsurePersonalAddressBook(cmd)
}

func (s *AddressBookService) GetSettings(query core.AddressBookSettingsQuery) (core.AddressBookSettingsResult, error) {
	return s.repo.GetAddressBookSettings(query)
}

func (s *AddressBookService) ListPeers(query core.AddressBookPeerListQuery) (core.AddressBookPeerListResult, error) {
	return s.repo.ListAddressBookPeers(query)
}

func (s *AddressBookService) ListTags(query core.AddressBookTagListQuery) ([]core.AddressBookTagView, error) {
	return s.repo.ListAddressBookTags(query)
}

func (s *AddressBookService) ListSharedProfiles(query core.SharedAddressBookListQuery) (core.SharedAddressBookListResult, error) {
	return s.repo.ListSharedAddressBooks(query)
}

func (s *AddressBookService) AddTag(cmd core.AddressBookTagAddCommand) error {
	return s.repo.AddAddressBookTag(cmd)
}

func (s *AddressBookService) UpdateTagColor(cmd core.AddressBookTagUpdateColorCommand) error {
	return s.repo.UpdateAddressBookTagColor(cmd)
}

func (s *AddressBookService) RenameTag(cmd core.AddressBookTagRenameCommand) error {
	return s.repo.RenameAddressBookTag(cmd)
}

func (s *AddressBookService) DeleteTags(cmd core.AddressBookTagDeleteCommand) error {
	return s.repo.DeleteAddressBookTags(cmd)
}

func (s *AddressBookService) CountPeers(userID, abID int) (int64, error) {
	return s.repo.CountAddressBookPeers(userID, abID)
}

func (s *AddressBookService) AddPeer(cmd core.AddressBookPeerCreateCommand) error {
	return s.repo.AddAddressBookPeer(cmd)
}

func (s *AddressBookService) UpdatePeer(cmd core.AddressBookPeerUpdateCommand) (bool, error) {
	return s.repo.UpdateAddressBookPeer(cmd)
}

func (s *AddressBookService) DeletePeers(cmd core.AddressBookPeerDeleteCommand) error {
	return s.repo.DeleteAddressBookPeers(cmd)
}

func (s *AddressBookService) ReplaceLegacyAddressBookData(cmd core.LegacyAddressBookReplaceCommand) error {
	return s.repo.ReplaceLegacyAddressBookData(cmd)
}
