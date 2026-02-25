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

	svc := NewCompatService(nil)
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
	svc := NewCompatService(nil)

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
