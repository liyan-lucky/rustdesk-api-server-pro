package repository

import (
	_ "modernc.org/sqlite"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"
	"testing"
	"xorm.io/xorm"
)

func TestDevicesCliAssignments(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	if err = engine.Sync(new(model.Device), new(model.Strategy), new(model.StrategyAssignment), new(model.DeviceGroup), new(model.DeviceGroupDevice), new(model.AddressBook), new(model.Peer)); err != nil {
		t.Fatal(err)
	}
	_, _ = engine.Insert(&model.Device{RustdeskId: "30001", Hostname: "host"})
	_, _ = engine.Insert(&model.Strategy{Guid: "s1", Name: "secure", Enabled: true})
	_, _ = engine.Insert(&model.DeviceGroup{Guid: "g1", Name: "ops"})
	ab := model.AddressBook{UserId: 1, Guid: "a1", Name: "team", Shared: true}
	_, _ = engine.Insert(&ab)
	strategy, addressBook, group := "secure", "team", "ops"
	err = NewXormCompatRepository(engine).ApplyDevicesCli(core.CompatDevicesCliCommand{UserID: 1, RustdeskID: "30001", StrategyName: &strategy, AddressBookName: &addressBook, DeviceGroupName: &group})
	if err != nil {
		t.Fatal(err)
	}
	if n, _ := engine.Count(new(model.StrategyAssignment)); n != 1 {
		t.Fatalf("strategy assignments=%d", n)
	}
	if n, _ := engine.Count(new(model.DeviceGroupDevice)); n != 1 {
		t.Fatalf("device group assignments=%d", n)
	}
	if n, _ := engine.Count(new(model.Peer)); n != 1 {
		t.Fatalf("address book peers=%d", n)
	}
}
