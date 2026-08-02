package repository

import (
	"errors"
	"testing"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func TestSharedAddressBookCanBeReadByAnotherUser(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", "file::memory:?cache=shared")
	if err != nil {
		t.Fatalf("new sqlite engine: %v", err)
	}
	defer engine.Close()
	if err := engine.Sync(new(model.AddressBook), new(model.AddressBookTag), new(model.Peer)); err != nil {
		t.Fatalf("sync schema: %v", err)
	}

	ab := model.AddressBook{UserId: 1, Guid: "shared-guid", Name: "Shared", Shared: true, Rule: 1}
	if _, err := engine.Insert(&ab); err != nil {
		t.Fatalf("insert address book: %v", err)
	}
	if _, err := engine.Insert(&model.Peer{UserId: 1, AbId: ab.Id, RustdeskId: "10001", Tags: `["ops"]`}); err != nil {
		t.Fatalf("insert shared peer: %v", err)
	}
	if _, err := engine.Insert(&model.AddressBookTag{UserId: 1, AbId: ab.Id, Name: "ops", Color: 42}); err != nil {
		t.Fatalf("insert shared tag: %v", err)
	}

	repo := NewXormAddressBookRepository(engine)
	peers, err := repo.ListAddressBookPeers(core.AddressBookPeerListQuery{UserID: 2, AbGuid: ab.Guid, Current: 1, PageSize: 100})
	if err != nil {
		t.Fatalf("list shared peers: %v", err)
	}
	if peers.Total != 1 || len(peers.Items) != 1 || peers.Items[0].ID != "10001" {
		t.Fatalf("unexpected shared peers: %+v", peers)
	}
	tags, err := repo.ListAddressBookTags(core.AddressBookTagListQuery{UserID: 2, AbGuid: ab.Guid})
	if err != nil {
		t.Fatalf("list shared tags: %v", err)
	}
	if len(tags) != 1 || tags[0].Name != "ops" || tags[0].Color != 42 {
		t.Fatalf("unexpected shared tags: %+v", tags)
	}
}

func TestPrivateAddressBookIsNotReadableByAnotherUser(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("new sqlite engine: %v", err)
	}
	defer engine.Close()
	if err := engine.Sync(new(model.AddressBook), new(model.AddressBookTag), new(model.Peer)); err != nil {
		t.Fatalf("sync schema: %v", err)
	}
	if _, err := engine.Insert(&model.AddressBook{UserId: 1, Guid: "private-guid", Name: "Private"}); err != nil {
		t.Fatalf("insert address book: %v", err)
	}

	repo := NewXormAddressBookRepository(engine)
	_, err = repo.ListAddressBookPeers(core.AddressBookPeerListQuery{UserID: 2, AbGuid: "private-guid", Current: 1, PageSize: 100})
	if !errors.Is(err, ErrAddressBookNotFound) {
		t.Fatalf("error = %v, want ErrAddressBookNotFound", err)
	}
}

func TestSharedAddressBookUserRuleLimitsAccess(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	if err = engine.Sync(new(model.AddressBook), new(model.AddressBookRule), new(model.Peer)); err != nil {
		t.Fatal(err)
	}
	ab := model.AddressBook{UserId: 1, Guid: "ruled", Name: "Ruled", Shared: true, Rule: 0}
	if _, err = engine.Insert(&ab); err != nil {
		t.Fatal(err)
	}
	if _, err = engine.Insert(&model.AddressBookRule{Guid: "rule-1", AbGuid: ab.Guid, TargetType: "user", TargetGuid: "2", Rule: 1}); err != nil {
		t.Fatal(err)
	}
	if _, err = engine.Insert(&model.Peer{UserId: 1, AbId: ab.Id, RustdeskId: "20001", Tags: "[]"}); err != nil {
		t.Fatal(err)
	}
	repo := NewXormAddressBookRepository(engine)
	if _, err = repo.ListAddressBookPeers(core.AddressBookPeerListQuery{UserID: 2, AbGuid: ab.Guid, Current: 1, PageSize: 10}); err != nil {
		t.Fatalf("allowed user: %v", err)
	}
	if _, err = repo.ListAddressBookPeers(core.AddressBookPeerListQuery{UserID: 3, AbGuid: ab.Guid, Current: 1, PageSize: 10}); !errors.Is(err, ErrAddressBookNotFound) {
		t.Fatalf("denied user error=%v", err)
	}
}

func TestModernTagsDriveLegacySyncAndPeerReferences(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	if err = engine.Sync(new(model.AddressBook), new(model.AddressBookTag), new(model.Tags), new(model.Peer)); err != nil {
		t.Fatal(err)
	}
	ab := model.AddressBook{UserId: 1, Guid: "personal", Name: model.PersonalAddressBookName}
	if _, err = engine.Insert(&ab); err != nil {
		t.Fatal(err)
	}
	_, _ = engine.Insert(&model.Tags{UserId: 1, Tag: "stale", Color: "1"})
	_, _ = engine.Insert(&model.AddressBookTag{UserId: 1, AbId: ab.Id, Name: "ops", Color: 42})
	_, _ = engine.Insert(&model.Peer{UserId: 1, AbId: ab.Id, RustdeskId: "10001", Tags: `["ops","keep"]`})

	repo := NewXormAddressBookRepository(engine)
	legacy, err := repo.GetLegacyAddressBook(core.LegacyAddressBookGetQuery{UserID: 1})
	if err != nil {
		t.Fatal(err)
	}
	if len(legacy.Tags) != 1 || legacy.Tags[0] != "ops" || legacy.TagColors["ops"] != 42 {
		t.Fatalf("modern tags were not authoritative: %+v", legacy)
	}
	if err = repo.RenameAddressBookTag(core.AddressBookTagRenameCommand{UserID: 1, AbID: ab.Id, Old: "ops", New: "prod"}); err != nil {
		t.Fatal(err)
	}
	var peer model.Peer
	_, _ = engine.Where("rustdesk_id = ?", "10001").Get(&peer)
	if peer.Tags != `["prod","keep"]` {
		t.Fatalf("rename did not rewrite peer tags: %s", peer.Tags)
	}
	if err = repo.DeleteAddressBookTags(core.AddressBookTagDeleteCommand{UserID: 1, AbID: ab.Id, Names: []string{"prod"}}); err != nil {
		t.Fatal(err)
	}
	peer = model.Peer{}
	_, _ = engine.Where("rustdesk_id = ?", "10001").Get(&peer)
	if peer.Tags != `["keep"]` {
		t.Fatalf("delete did not remove peer tag: %s", peer.Tags)
	}
}
