package repository

import (
	"testing"

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
