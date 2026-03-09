package service

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/repository"
)

// CompatSysinfoVersion is exposed to admin dashboard and compatibility probes.
// Keep this value aligned with the validated upstream RustDesk versions.
const CompatSysinfoVersion = "rustdesk-api-server-pro-compat-client-1.4.6-server-1.1.15-latest"
const compatRecordDir = "record_uploads"

type CompatService struct {
	repo repository.CompatRepository
}

func NewCompatService(repo repository.CompatRepository) *CompatService {
	return &CompatService{repo: repo}
}

func (s *CompatService) LoginOptions() core.CompatLoginOptionsResult {
	return core.CompatLoginOptionsResult{Options: []string{}}
}

func (s *CompatService) OidcAuth() core.CompatOidcAuthResult {
	return core.CompatOidcAuthResult{
		Error:   "OIDC_NOT_SUPPORTED",
		Enabled: false,
		URL:     "",
	}
}

func (s *CompatService) OidcAuthQuery() core.CompatOidcAuthQueryResult {
	return core.CompatOidcAuthQueryResult{
		Error:   "OIDC_NOT_SUPPORTED",
		Enabled: false,
		User:    nil,
	}
}

func (s *CompatService) PluginSign(msg []byte) core.CompatPluginSignResult {
	return core.CompatPluginSignResult{SignedMsg: msg}
}

func (s *CompatService) ApplyDevicesCli(cmd core.CompatDevicesCliCommand) error {
	return s.repo.ApplyDevicesCli(cmd)
}

func (s *CompatService) HandleRecord(cmd core.CompatRecordCommand) error {
	op := strings.ToLower(strings.TrimSpace(cmd.Op))
	fileName := sanitizeRecordFileName(cmd.FileName)
	if op == "" {
		return errors.New("type required")
	}
	if fileName == "" {
		return errors.New("file required")
	}

	fullPath, err := prepareRecordPath(fileName)
	if err != nil {
		return err
	}

	switch op {
	case "new":
		f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		return f.Close()
	case "part":
		if cmd.Offset < 0 {
			return errors.New("invalid offset")
		}
		f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err = f.Seek(cmd.Offset, io.SeekStart); err != nil {
			return err
		}
		_, err = io.Copy(f, bytes.NewReader(cmd.Body))
		return err
	case "tail":
		f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, bytes.NewReader(cmd.Body))
		return err
	case "remove":
		if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
			return err
		}
		return nil
	default:
		return errors.New("unsupported record op")
	}
}

func sanitizeRecordFileName(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return ""
	}
	base := filepath.Base(name)
	base = strings.ReplaceAll(base, "..", "")
	base = strings.TrimSpace(base)
	return base
}

func prepareRecordPath(fileName string) (string, error) {
	dir := filepath.Join(".", compatRecordDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(dir, fileName), nil
}
