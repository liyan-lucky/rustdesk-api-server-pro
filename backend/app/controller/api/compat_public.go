package api

import (
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

const compatSysinfoVersion = "rustdesk-api-server-pro-compat-1"
const compatRecordDir = "record_uploads"

// CompatPublicController provides compatibility endpoints used by newer RustDesk clients.
// Unsupported features return stable "error" payloads instead of 404.
type CompatPublicController struct {
	basicController
}

func (c *CompatPublicController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "sysinfo_ver", "HandleSysinfoVer")
	b.Handle("POST", "oidc/auth", "HandleOidcAuth")
	b.Handle("GET", "oidc/auth-query", "HandleOidcAuthQuery")
	b.Handle("POST", "record", "HandleRecord")
}

func (c *CompatPublicController) HandleSysinfoVer() mvc.Result {
	// Stable version string so clients can cache sysinfo uploads when payload is unchanged.
	return mvc.Response{
		Text: compatSysinfoVersion,
	}
}

func (c *CompatPublicController) HandleOidcAuth() mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"error":   "OIDC_NOT_SUPPORTED",
			"enabled": false,
			"url":     "",
		},
	}
}

func (c *CompatPublicController) HandleOidcAuthQuery() mvc.Result {
	return mvc.Response{
		Object: iris.Map{
			"error":   "OIDC_NOT_SUPPORTED",
			"enabled": false,
			"user":    nil,
		},
	}
}

func (c *CompatPublicController) HandleRecord() mvc.Result {
	op := strings.ToLower(strings.TrimSpace(c.Ctx.URLParamDefault("type", "")))
	fileName := sanitizeRecordFileName(c.Ctx.URLParamDefault("file", ""))
	if op == "" {
		_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
		return mvc.Response{
			Object: iris.Map{
				"error": "type required",
			},
		}
	}
	if fileName == "" {
		_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
		return mvc.Response{
			Object: iris.Map{
				"error": "file required",
			},
		}
	}

	fullPath, err := prepareRecordPath(fileName)
	if err != nil {
		_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	switch op {
	case "new":
		_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
		f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			return mvc.Response{Object: iris.Map{"error": err.Error()}}
		}
		_ = f.Close()
		return mvc.Response{}
	case "part":
		offset, err := strconv.ParseInt(c.Ctx.URLParamDefault("offset", "0"), 10, 64)
		if err != nil || offset < 0 {
			_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
			return mvc.Response{Object: iris.Map{"error": "invalid offset"}}
		}
		f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
			return mvc.Response{Object: iris.Map{"error": err.Error()}}
		}
		defer f.Close()
		if _, err = f.Seek(offset, io.SeekStart); err != nil {
			_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
			return mvc.Response{Object: iris.Map{"error": err.Error()}}
		}
		if _, err = io.Copy(f, c.Ctx.Request().Body); err != nil {
			return mvc.Response{Object: iris.Map{"error": err.Error()}}
		}
		return mvc.Response{}
	case "tail":
		// Final chunk/header is sent as plain body; append as received.
		f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
			return mvc.Response{Object: iris.Map{"error": err.Error()}}
		}
		defer f.Close()
		if _, err = io.Copy(f, c.Ctx.Request().Body); err != nil {
			return mvc.Response{Object: iris.Map{"error": err.Error()}}
		}
		return mvc.Response{}
	case "remove":
		_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
		if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
			return mvc.Response{Object: iris.Map{"error": err.Error()}}
		}
		return mvc.Response{}
	default:
		_, _ = io.Copy(io.Discard, c.Ctx.Request().Body)
		return mvc.Response{
			Object: iris.Map{
				"error": "unsupported record op",
			},
		}
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
