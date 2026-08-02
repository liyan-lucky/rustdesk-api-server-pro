package repository

import (
	"reflect"
	"testing"
	"time"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

func TestHeartbeatDoesNotInventStrategyVersion(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("new sqlite engine: %v", err)
	}
	defer engine.Close()
	if err := engine.Sync(new(model.Device)); err != nil {
		t.Fatalf("sync schema: %v", err)
	}

	const clientVersion int64 = 1725698100
	result, err := NewXormSystemRepository(engine).UpsertHeartbeat(core.HeartbeatCommand{
		RustdeskID: "10001",
		UUID:       "uuid-1",
		ModifiedAt: clientVersion,
	})
	if err != nil {
		t.Fatalf("heartbeat: %v", err)
	}
	if result.ModifiedAt != clientVersion {
		t.Fatalf("modified_at = %d, want %d", result.ModifiedAt, clientVersion)
	}
}

func TestHeartbeatDeliversAssignedStrategyAndDisconnect(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	if err = engine.Sync(new(model.Device), new(model.Strategy), new(model.StrategyAssignment), new(model.DeviceGroupDevice)); err != nil {
		t.Fatal(err)
	}
	now := time.Now().Add(time.Second).Truncate(time.Second)
	device := model.Device{RustdeskId: "10002", Disabled: true}
	if _, err = engine.Insert(&device); err != nil {
		t.Fatal(err)
	}
	strategy := model.Strategy{Guid: "strategy-1", Name: "secure", Content: `{"config_options":{"enable-file-transfer":"N","custom-rendezvous-server":"hbbs.local"}}`, Enabled: true, UpdatedAt: now}
	if _, err = engine.Insert(&strategy); err != nil {
		t.Fatal(err)
	}
	if _, err = engine.Insert(&model.StrategyAssignment{StrategyGuid: strategy.Guid, TargetType: "device", TargetGuid: device.RustdeskId}); err != nil {
		t.Fatal(err)
	}
	result, err := NewXormSystemRepository(engine).UpsertHeartbeat(core.HeartbeatCommand{RustdeskID: device.RustdeskId, ModifiedAt: 1, Conns: []int{7, 9}, ConnCount: 2})
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]string{"enable-file-transfer": "N", "custom-rendezvous-server": "hbbs.local"}
	if !reflect.DeepEqual(result.Strategy, want) {
		t.Fatalf("strategy=%v want=%v", result.Strategy, want)
	}
	if result.ModifiedAt <= 1 {
		t.Fatalf("modified_at=%d", result.ModifiedAt)
	}
	if !reflect.DeepEqual(result.Disconnect, []int{7, 9}) {
		t.Fatalf("disconnect=%v", result.Disconnect)
	}
}
