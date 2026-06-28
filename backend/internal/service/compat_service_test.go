package service

import (
	"os"
	"path/filepath"
	"testing"

	"rustdesk-api-server-pro/internal/core"
)

func TestCompatServiceHandleRecordLifecycle(t *testing.T) {
	tmp := t.TempDir()
	oldWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir tmp: %v", err)
	}
	t.Cleanup(func() { _ = os.Chdir(oldWd) })

	svc := NewCompatService(nil, nil, nil)
	file := "demo.rec"

	if err := svc.HandleRecord(core.CompatRecordCommand{
		Op:       "new",
		FileName: file,
	}); err != nil {
		t.Fatalf("new op failed: %v", err)
	}

	if err := svc.HandleRecord(core.CompatRecordCommand{
		Op:       "part",
		FileName: file,
		Offset:   0,
		Body:     []byte("abc"),
	}); err != nil {
		t.Fatalf("part op failed: %v", err)
	}

	if err := svc.HandleRecord(core.CompatRecordCommand{
		Op:       "tail",
		FileName: file,
		Body:     []byte("XYZ"),
	}); err != nil {
		t.Fatalf("tail op failed: %v", err)
	}

	fullPath := filepath.Join(tmp, "record_uploads", file)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		t.Fatalf("read record file: %v", err)
	}
	if string(data) != "abcXYZ" {
		t.Fatalf("unexpected file content: %q", string(data))
	}

	if err := svc.HandleRecord(core.CompatRecordCommand{
		Op:       "remove",
		FileName: file,
	}); err != nil {
		t.Fatalf("remove op failed: %v", err)
	}

	if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
		t.Fatalf("expected record file removed, stat err=%v", err)
	}
}

func TestCompatServiceHandleRecordValidation(t *testing.T) {
	svc := NewCompatService(nil, nil, nil)

	tests := []struct {
		name string
		cmd  core.CompatRecordCommand
		want string
	}{
		{
			name: "missing type",
			cmd:  core.CompatRecordCommand{FileName: "a.rec"},
			want: "type required",
		},
		{
			name: "missing file",
			cmd:  core.CompatRecordCommand{Op: "new"},
			want: "file required",
		},
		{
			name: "invalid offset",
			cmd:  core.CompatRecordCommand{Op: "part", FileName: "a.rec", Offset: -1},
			want: "invalid offset",
		},
		{
			name: "unsupported op",
			cmd:  core.CompatRecordCommand{Op: "noop", FileName: "a.rec"},
			want: "unsupported record op",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.HandleRecord(tt.cmd)
			if err == nil || err.Error() != tt.want {
				t.Fatalf("got err=%v, want %q", err, tt.want)
			}
		})
	}
}

func TestCompatServiceTargetContract(t *testing.T) {
	svc := NewCompatService(nil, nil, nil)
	target := svc.Target()

	if got := target["project"]; got != "rustdesk-api-server-pro" {
		t.Fatalf("unexpected project: %v", got)
	}
	if got := target["sysinfo_version"]; got != CompatSysinfoVersion {
		t.Fatalf("unexpected sysinfo_version: %v", got)
	}

	client, ok := target["client"].(map[string]any)
	if !ok {
		t.Fatalf("client target should be map[string]any")
	}
	if got := client["name"]; got != CompatClientName {
		t.Fatalf("unexpected client name: %v", got)
	}
	if got := client["version"]; got != CompatClientVersion {
		t.Fatalf("unexpected client version: %v", got)
	}
	if got := client["release_date"]; got != CompatClientReleaseDate {
		t.Fatalf("unexpected client release date: %v", got)
	}

	server, ok := target["server"].(map[string]any)
	if !ok {
		t.Fatalf("server target should be map[string]any")
	}
	if got := server["version"]; got != CompatServerVersion {
		t.Fatalf("unexpected server version: %v", got)
	}
	if got := server["status"]; got != CompatTargetStatus {
		t.Fatalf("unexpected server status: %v", got)
	}

	features, ok := target["features"].(map[string]bool)
	if !ok {
		t.Fatalf("features should be map[string]bool")
	}
	for _, key := range []string{
		"address_book",
		"audit",
		"file_transfer_audit",
		"alarm_audit",
		"device_group",
		"user_group",
		"strategy",
		"record",
		"plugin_sign_passthrough",
	} {
		if !features[key] {
			t.Fatalf("expected feature %s to be enabled", key)
		}
	}

	focus, ok := target["official_focus"].([]string)
	if !ok {
		t.Fatalf("official_focus should be []string")
	}
	for _, key := range []string{"windows_arm64_support", "remote_restart_autoconnect", "oidc_microsoft_icon_compat"} {
		if !containsString(focus, key) {
			t.Fatalf("expected official_focus to contain %s", key)
		}
	}

	endpoints, ok := target["probe_endpoints"].([]string)
	if !ok {
		t.Fatalf("probe_endpoints should be []string")
	}
	for _, path := range []string{"/api/health", "/api/compat-target", "/api/server/info", "/api/devices/deploy"} {
		if !containsString(endpoints, path) {
			t.Fatalf("expected probe_endpoints to contain %s", path)
		}
	}
}

func TestCompatServiceHandleDeviceDeploy(t *testing.T) {
	svc := NewCompatService(nil, nil, nil)

	tests := []struct {
		name string
		cmd  core.CompatDeviceDeployCommand
		want string
	}{
		{
			name: "missing id",
			cmd: core.CompatDeviceDeployCommand{
				UUID:      "uuid",
				PublicKey: "pk",
			},
			want: "INVALID_INPUT",
		},
		{
			name: "missing uuid",
			cmd: core.CompatDeviceDeployCommand{
				RustdeskID: "123456789",
				PublicKey: "pk",
			},
			want: "INVALID_INPUT",
		},
		{
			name: "missing public key",
			cmd: core.CompatDeviceDeployCommand{
				RustdeskID: "123456789",
				UUID:      "uuid",
			},
			want: "INVALID_INPUT",
		},
		{
			name: "deployment not required",
			cmd: core.CompatDeviceDeployCommand{
				RustdeskID: "123456789",
				UUID:      "uuid",
				PublicKey: "pk",
			},
			want: "NOT_ENABLED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := svc.HandleDeviceDeploy(tt.cmd)
			if got.Result != tt.want {
				t.Fatalf("got result=%q, want %q", got.Result, tt.want)
			}
		})
	}
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
