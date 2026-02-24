package api

import (
	"encoding/json"
	"io"
	"rustdesk-api-server-pro/app/model"
	"strings"

	"github.com/kataras/iris/v12"
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
	body, err := io.ReadAll(c.Ctx.Request().Body)
	if err != nil {
		return mvc.Response{
			Object: iris.Map{
				"error": err.Error(),
			},
		}
	}

	rustdeskID := gjson.GetBytes(body, "id").String()
	if rustdeskID == "" {
		return mvc.Response{
			Object: iris.Map{
				"error": "id required",
			},
		}
	}

	user := c.GetUser()

	// Update runtime device metadata if present.
	deviceUpdates := make([]string, 0, 2)
	device := model.Device{}
	if v := gjson.GetBytes(body, "device_name"); v.Exists() {
		device.Hostname = v.String()
		deviceUpdates = append(deviceUpdates, "hostname")
	}
	if v := gjson.GetBytes(body, "device_username"); v.Exists() {
		device.Username = v.String()
		deviceUpdates = append(deviceUpdates, "username")
	}
	if len(deviceUpdates) > 0 {
		_, _ = c.Db.Where("rustdesk_id = ?", rustdeskID).Cols(deviceUpdates...).Update(&device)
	}

	// Update persisted peer metadata for current user (legacy + address book peers share the same table).
	peer := model.Peer{}
	peerUpdates := make([]string, 0, 8)
	if v := gjson.GetBytes(body, "note"); v.Exists() {
		peer.Note = v.String()
		peerUpdates = append(peerUpdates, "note")
	}
	if v := gjson.GetBytes(body, "user_name"); v.Exists() {
		peer.LoginName = v.String()
		peerUpdates = append(peerUpdates, "loginName")
	}
	if v := gjson.GetBytes(body, "device_name"); v.Exists() {
		peer.Hostname = v.String()
		peerUpdates = append(peerUpdates, "hostname")
	}
	if v := gjson.GetBytes(body, "device_username"); v.Exists() {
		peer.Username = v.String()
		peerUpdates = append(peerUpdates, "username")
	}
	if v := gjson.GetBytes(body, "address_book_alias"); v.Exists() {
		peer.Alias = v.String()
		peerUpdates = append(peerUpdates, "alias")
	}
	if v := gjson.GetBytes(body, "address_book_password"); v.Exists() {
		peer.Password = v.String()
		peerUpdates = append(peerUpdates, "password")
	}
	if v := gjson.GetBytes(body, "address_book_note"); v.Exists() {
		peer.Note = v.String()
		if !contains(peerUpdates, "note") {
			peerUpdates = append(peerUpdates, "note")
		}
	}
	if v := gjson.GetBytes(body, "address_book_tag"); v.Exists() {
		tagText := strings.TrimSpace(v.String())
		tags := []string{}
		if tagText != "" {
			for _, t := range strings.Split(tagText, ",") {
				t = strings.TrimSpace(t)
				if t != "" {
					tags = append(tags, t)
				}
			}
		}
		if b, mErr := json.Marshal(tags); mErr == nil {
			peer.Tags = string(b)
			peerUpdates = append(peerUpdates, "tags")
		}
	}
	if len(peerUpdates) > 0 {
		_, _ = c.Db.Where("user_id = ? and rustdesk_id = ?", user.Id, rustdeskID).Cols(peerUpdates...).Update(&peer)
	}

	// Keep compatibility behavior with the official client CLI helper: empty body means success.
	return mvc.Response{Text: ""}
}

func contains(items []string, target string) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}
	return false
}
