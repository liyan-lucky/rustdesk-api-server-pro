package httpdto

import (
	"encoding/json"
	"testing"

	"rustdesk-api-server-pro/internal/core"
)

func TestNewLegacyAddressBookResponse(t *testing.T) {
	result := core.LegacyAddressBookGetResult{
		Tags: []string{"prod", "ops"},
		TagColors: map[string]int64{
			"prod": 1,
			"ops":  2,
		},
		Peers: []core.LegacyAddressBookPeerEntry{
			{
				RustdeskID: "10001",
				Hash:       "h1",
				Username:   "u1",
				Hostname:   "host1",
				Platform:   "Windows",
				Alias:      "pc-1",
				Tags:       []string{"prod"},
				Note:       "note-1",
			},
		},
	}

	resp, err := NewLegacyAddressBookResponse(99, result)
	if err != nil {
		t.Fatalf("NewLegacyAddressBookResponse err: %v", err)
	}

	if resp.LicensedDevices != 99 {
		t.Fatalf("licensed_devices = %d, want 99", resp.LicensedDevices)
	}
	if resp.Data == "" {
		t.Fatalf("data should not be empty")
	}

	var payload struct {
		Tags  []string `json:"tags"`
		Peers []struct {
			ID   string `json:"id"`
			Note string `json:"note"`
		} `json:"peers"`
		TagColors string `json:"tag_colors"`
	}
	if err := json.Unmarshal([]byte(resp.Data), &payload); err != nil {
		t.Fatalf("unmarshal data payload: %v", err)
	}

	if len(payload.Tags) != 2 {
		t.Fatalf("tags len = %d, want 2", len(payload.Tags))
	}
	if len(payload.Peers) != 1 || payload.Peers[0].ID != "10001" || payload.Peers[0].Note != "note-1" {
		t.Fatalf("unexpected peers payload: %+v", payload.Peers)
	}

	var tagColors map[string]int64
	if err := json.Unmarshal([]byte(payload.TagColors), &tagColors); err != nil {
		t.Fatalf("unmarshal tag_colors string: %v", err)
	}
	if tagColors["prod"] != 1 || tagColors["ops"] != 2 {
		t.Fatalf("unexpected tag colors: %+v", tagColors)
	}
}
