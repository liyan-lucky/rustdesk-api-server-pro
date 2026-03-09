package api

import (
	"strings"

	"rustdesk-api-server-pro/internal/core"

	"github.com/kataras/iris/v12/mvc"
	"github.com/tidwall/gjson"
)

// CompatAuthController contains authenticated compatibility endpoints for the desktop client.
type CompatAuthController struct {
	basicController
}

func (c *CompatAuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "devices/cli", "HandleDevicesCli")
}

func (c *CompatAuthController) HandleDevicesCli() mvc.Result {
	body, err := c.readBodyBytes()
	if err != nil {
		return c.fail(err)
	}

	rustdeskID := gjson.GetBytes(body, "id").String()
	if rustdeskID == "" {
		return c.failMsg("id required")
	}

	user := c.GetUser()
	cmd := core.CompatDevicesCliCommand{
		UserID:     user.Id,
		RustdeskID: rustdeskID,
	}

	if v := gjson.GetBytes(body, "device_name"); v.Exists() {
		value := v.String()
		cmd.DeviceName = &value
	}
	if v := gjson.GetBytes(body, "device_username"); v.Exists() {
		value := v.String()
		cmd.DeviceUsername = &value
	}
	if v := gjson.GetBytes(body, "user_name"); v.Exists() {
		value := v.String()
		cmd.UserName = &value
	}

	if v := gjson.GetBytes(body, "note"); v.Exists() {
		value := v.String()
		cmd.Note = &value
	}
	if v := gjson.GetBytes(body, "address_book_note"); v.Exists() {
		// address_book_note should override generic note when both exist.
		value := v.String()
		cmd.Note = &value
	}
	if v := gjson.GetBytes(body, "address_book_alias"); v.Exists() {
		value := v.String()
		cmd.Alias = &value
	}
	if v := gjson.GetBytes(body, "address_book_password"); v.Exists() {
		value := v.String()
		cmd.Password = &value
	}
	if v := gjson.GetBytes(body, "address_book_tag"); v.Exists() {
		cmd.HasTags = true
		tagText := strings.TrimSpace(v.String())
		if tagText != "" {
			for _, t := range strings.Split(tagText, ",") {
				t = strings.TrimSpace(t)
				if t != "" {
					cmd.Tags = append(cmd.Tags, t)
				}
			}
		}
	}

	if err := c.compatService().ApplyDevicesCli(cmd); err != nil {
		return c.fail(err)
	}

	// Keep compatibility behavior with the official client CLI helper: empty body means success.
	return mvc.Response{Text: ""}
}
