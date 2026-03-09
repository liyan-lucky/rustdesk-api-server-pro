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
				ConnID:     connID,
				RustdeskID: rustdeskID,
				IP:         gjson.GetBytes(body, "ip").String(),
				SessionID:  sessionID,
				UUID:       uuid,
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
		RustdeskID: gjson.GetBytes(body, "id").String(),
		Info:       gjson.GetBytes(body, "info").String(),
		IsFile:     gjson.GetBytes(body, "is_file").Bool(),
		Path:       gjson.GetBytes(body, "path").String(),
		PeerID:     gjson.GetBytes(body, "peer_id").String(),
		Type:       int(gjson.GetBytes(body, "type").Int()),
		UUID:       gjson.GetBytes(body, "uuid").String(),
	}); err != nil {
		return c.fail(err)
	}

	return c.ok()
}

func (c *AuditController) PostAuditAlarm() mvc.Result {
	// Keep compatibility with official clients: accept alarm reports even though
	// this server does not persist alarm payloads yet.
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
