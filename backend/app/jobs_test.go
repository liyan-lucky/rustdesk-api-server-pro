package app

import (
	"testing"
	"time"

	"rustdesk-api-server-pro/app/model"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

// TestMarkExpiredDevicesOffline 验证官方三十秒边界只下线过期设备。
func TestMarkExpiredDevicesOffline(t *testing.T) {
	engine, err := xorm.NewEngine("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()
	if err = engine.Sync(new(model.Device)); err != nil {
		t.Fatal(err)
	}

	now := time.Now().Truncate(time.Second)
	if _, err = engine.Exec("INSERT INTO device (rustdesk_id, is_online, updated_at) VALUES (?, 1, ?), (?, 1, ?)",
		"expired", now.Add(-31*time.Second), "fresh", now.Add(-29*time.Second)); err != nil {
		t.Fatal(err)
	}
	var before []model.Device
	if err = engine.Asc("rustdesk_id").Find(&before); err != nil {
		t.Fatal(err)
	}
	if len(before) != 2 || !before[0].IsOnline || !before[1].IsOnline {
		t.Fatalf("测试前置在线状态错误: %+v", before)
	}

	cutoff := now.Add(-30 * time.Second).UTC().Format("2006-01-02 15:04:05")
	if _, err = markExpiredDevicesOffline(engine, cutoff); err != nil {
		t.Fatal(err)
	}

	var devices []model.Device
	if err = engine.Asc("rustdesk_id").Find(&devices); err != nil {
		t.Fatal(err)
	}
	if len(devices) != 2 || devices[0].RustdeskId != "expired" || devices[0].IsOnline {
		t.Fatalf("过期设备未下线: %+v", devices)
	}
	if devices[1].RustdeskId != "fresh" || !devices[1].IsOnline {
		t.Fatalf("新鲜设备被误下线: %+v", devices)
	}
}
