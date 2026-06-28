package api

import (
	"rustdesk-api-server-pro/internal/core"

	"github.com/kataras/iris/v12/mvc"
	"github.com/tidwall/gjson"
)

type AuditController struct {
	basicController
}

func (c *AuditController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("PUT", "audit", "HandleAuditUpdate")
}

func (c *AuditController) PostAuditConn() mvc.Result {
	body, err := c.readBodyBytes()
	if err != nil {
		return c.fail(err)
	}

	rustdeskID := gjson.GetBytes(body, "id").String()
	sessionID := gjson.GetBytes(body, "session_id").String()

	// Note-only update branch.
	if gjson.GetBytes(body, "note").Exists() {
		if err := c.auditService().UpdateConnNote(core.AuditConnNoteCommand{
			SessionID: sessionID,
			Note:      gjson.GetBytes(body, "note").String(),
		}); err != nil {
			return c.fail(err)
		}
		return c.ok()
	}

	connID := int(gjson.GetBytes(body, "conn_id").Int())
	uuid := gjson.GetBytes(body, "uuid").String()

	if actionResult := gjson.GetBytes(body, "action"); actionResult.Exists() {
		switch actionResult.String() {
		case "new":
			if err := c.auditService().OpenConn(core.AuditConnOpenCommand{
				ConnID:        connID,
				RustdeskID:    rustdeskID,
				PeerID:        firstJSONValue(body, "peer_id", "peerId", "peer"),
				IP:            gjson.GetBytes(body, "ip").String(),
				PeerIP:        firstJSONValue(body, "peer_ip", "peerIp"),
				SessionID:     sessionID,
				UUID:          uuid,
				Direction:     firstJSONValue(body, "direction"),
				Status:        "open",
				ClientVersion: firstJSONValue(body, "version", "client_version", "clientVersion"),
				Platform:      firstJSONValue(body, "platform", "os"),
				Hostname:      firstJSONValue(body, "hostname", "host_name"),
				Raw:           string(body),
			}); err != nil {
				return c.fail(err)
			}
		case "close":
			if err := c.auditService().CloseConn(core.AuditConnCloseCommand{ConnID: connID}); err != nil {
				return c.fail(err)
			}
		}
		return c.ok()
	}

	if peerResult := gjson.GetBytes(body, "peer"); peerResult.Exists() {
		if err := c.auditService().UpdateConnSession(core.AuditConnSessionUpdateCommand{
			ConnID:    connID,
			SessionID: sessionID,
			Type:      int(gjson.GetBytes(body, "type").Int()),
			Peer:      peerResult.String(),
			PeerID:    firstJSONValue(body, "peer_id", "peerId"),
		}); err != nil {
			return c.fail(err)
		}
	}

	return c.ok()
}

func (c *AuditController) PostAuditFile() mvc.Result {
	body, err := c.readBodyBytes()
	if err != nil {
		return c.fail(err)
	}

	// Note: fix an old bug here: uuid should come from "uuid", not "type".
	if err := c.auditService().CreateFileTransfer(core.FileTransferCreateCommand{
		RustdeskID:   gjson.GetBytes(body, "id").String(),
		Info:         gjson.GetBytes(body, "info").String(),
		IsFile:       gjson.GetBytes(body, "is_file").Bool(),
		Path:         gjson.GetBytes(body, "path").String(),
		FileName:     firstJSONValue(body, "file_name", "filename", "name"),
		PeerID:       gjson.GetBytes(body, "peer_id").String(),
		SessionID:    firstJSONValue(body, "session_id", "guid"),
		Type:         int(gjson.GetBytes(body, "type").Int()),
		UUID:         gjson.GetBytes(body, "uuid").String(),
		Direction:    firstJSONValue(body, "direction"),
		SizeBytes:    gjson.GetBytes(body, "size").Int(),
		Result:       firstJSONValue(body, "result", "status"),
		ErrorMessage: firstJSONValue(body, "error", "error_message", "message"),
		Raw:          string(body),
	}); err != nil {
		return c.fail(err)
	}

	return c.ok()
}

func (c *AuditController) PostAuditAlarm() mvc.Result {
	body, err := c.readBodyBytes()
	if err != nil {
		return c.fail(err)
	}

	if err := c.auditService().CreateAlarmAudit(core.AlarmAuditCreateCommand{
		RustdeskID: firstJSONValue(body, "id", "rustdesk_id", "rustdeskId"),
		PeerID:     firstJSONValue(body, "peer_id", "peerId", "peer"),
		SessionID:  firstJSONValue(body, "session_id", "guid"),
		AlarmType:  firstJSONValue(body, "alarm_type", "alarmType", "type", "event"),
		Severity:   firstJSONValue(body, "severity", "level"),
		Message:    firstJSONValue(body, "message", "msg", "info"),
		Raw:        string(body),
	}); err != nil {
		return c.fail(err)
	}

	return c.ok()
}

func (c *AuditController) HandleAuditUpdate() mvc.Result {
	body, err := c.readBodyBytes()
	if err != nil {
		return c.fail(err)
	}

	guid := gjson.GetBytes(body, "guid").String()
	if guid == "" {
		return c.failMsg("guid required")
	}

	if err := c.auditService().UpdateNoteByGuid(core.AuditGuidNoteUpdateCommand{
		Guid: guid,
		Note: gjson.GetBytes(body, "note").String(),
	}); err != nil {
		return c.fail(err)
	}

	return c.ok()
}

func firstJSONValue(body []byte, keys ...string) string {
	for _, key := range keys {
		value := gjson.GetBytes(body, key)
		if value.Exists() {
			return value.String()
		}
	}
	return ""
}
