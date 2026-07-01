package util

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHttpGetStringRejectsErrorStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer server.Close()

	if _, err := HttpGetString(server.URL); err == nil {
		t.Fatal("expected error status to be rejected")
	}
}

func TestDownloadFileRejectsErrorStatusAndLeavesNoFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	}))
	defer server.Close()

	filename := filepath.Join(t.TempDir(), "download.bin")
	if err := DownloadFile(server.URL, filename, false); err == nil {
		t.Fatal("expected error status to be rejected")
	}
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		t.Fatalf("download file should not exist after failed response, stat err=%v", err)
	}
}

func TestDownloadFileUsesSafePermissionsAndTruncatesOldFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	filename := filepath.Join(t.TempDir(), "download.bin")
	if err := os.WriteFile(filename, []byte("old-content"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := DownloadFile(server.URL, filename, false); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "ok" {
		t.Fatalf("unexpected file content: %q", string(content))
	}

	info, err := os.Stat(filename)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0644 {
		t.Fatalf("unexpected file mode: got %o want 0644", info.Mode().Perm())
	}
}
