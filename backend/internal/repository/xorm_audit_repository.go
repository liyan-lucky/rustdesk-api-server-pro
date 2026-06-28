package repository

import (
	"strconv"
	"time"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"

	"xorm.io/xorm"
)

type XormAuditRepository struct {
	DB *xorm.Engine
}

func NewXormAuditRepository(dbEngine *xorm.Engine) *XormAuditRepository {
	return &XormAuditRepository{DB: dbEngine}
}

func (r *XormAuditRepository) UpdateAuditConnNote(cmd core.AuditConnNoteCommand) error {
	_, err := r.DB.Where("session_id = ?", cmd.SessionID).Update(&model.Audit{Note: cmd.Note})
	return err
}

func (r *XormAuditRepository) InsertAuditConnOpen(cmd core.AuditConnOpenCommand) error {
	status := cmd.Status
	if status == "" {
		status = "open"
	}

	now := time.Now()
	_, err := r.DB.Insert(&model.Audit{
		ConnId:        cmd.ConnID,
		RustdeskId:    cmd.RustdeskID,
		PeerId:        cmd.PeerID,
		IP:            cmd.IP,
		PeerIP:        cmd.PeerIP,
		SessionId:     cmd.SessionID,
		Uuid:          cmd.UUID,
		Direction:     cmd.Direction,
		Status:        status,
		ClientVersion: cmd.ClientVersion,
		Platform:      cmd.Platform,
		Hostname:      cmd.Hostname,
		Raw:           cmd.Raw,
		StartedAt:     now,
	})
	return err
}

func (r *XormAuditRepository) CloseAuditConn(cmd core.AuditConnCloseCommand) error {
	var audit model.Audit
	has, err := r.DB.Where("conn_id = ?", cmd.ConnID).Get(&audit)
	if err != nil {
		return err
	}

	now := time.Now()
	update := &model.Audit{ClosedAt: now, Status: "closed"}
	if has && !audit.StartedAt.IsZero() {
		update.DurationSeconds = int64(now.Sub(audit.StartedAt).Seconds())
	}

	_, err = r.DB.Where("conn_id = ?", cmd.ConnID).Update(update)
	return err
}

func (r *XormAuditRepository) UpdateAuditConnSession(cmd core.AuditConnSessionUpdateCommand) error {
	_, err := r.DB.Where("conn_id = ?", cmd.ConnID).Update(&model.Audit{
		SessionId: cmd.SessionID,
		Type:      cmd.Type,
		Peer:      cmd.Peer,
		PeerId:    cmd.PeerID,
	})
	return err
}

func (r *XormAuditRepository) InsertFileTransfer(cmd core.FileTransferCreateCommand) error {
	result := cmd.Result
	if result == "" {
		result = "unknown"
	}

	_, err := r.DB.Insert(&model.FileTransfer{
		RustdeskId:   cmd.RustdeskID,
		Info:         cmd.Info,
		IsFile:       cmd.IsFile,
		Path:         cmd.Path,
		FileName:     cmd.FileName,
		PeerId:       cmd.PeerID,
		SessionId:    cmd.SessionID,
		Type:         cmd.Type,
		Uuid:         cmd.UUID,
		Direction:    cmd.Direction,
		SizeBytes:    cmd.SizeBytes,
		Result:       result,
		ErrorMessage: cmd.ErrorMessage,
		Raw:          cmd.Raw,
	})
	return err
}

func (r *XormAuditRepository) UpdateAuditNoteByGuid(cmd core.AuditGuidNoteUpdateCommand) error {
	affected, err := r.DB.Where("session_id = ?", cmd.Guid).Cols("note").Update(&model.Audit{Note: cmd.Note})
	if err != nil {
		return err
	}
	if affected == 0 {
		if id, convErr := strconv.Atoi(cmd.Guid); convErr == nil {
			_, err = r.DB.Where("id = ?", id).Cols("note").Update(&model.Audit{Note: cmd.Note})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *XormAuditRepository) InsertAlarmAudit(cmd core.AlarmAuditCreateCommand) error {
	_, err := r.DB.Insert(&model.AlarmAudit{
		RustdeskId: cmd.RustdeskID,
		PeerId:     cmd.PeerID,
		SessionId:  cmd.SessionID,
		AlarmType:  cmd.AlarmType,
		Severity:   cmd.Severity,
		Message:    cmd.Message,
		Raw:        cmd.Raw,
	})
	return err
}

func (r *XormAuditRepository) InsertSecurityAudit(cmd core.SecurityAuditCreateCommand) error {
	_, err := r.DB.Insert(&model.SecurityAudit{
		UserId:    cmd.UserID,
		Username:  cmd.Username,
		Event:     cmd.Event,
		IP:        cmd.IP,
		UserAgent: cmd.UserAgent,
		Success:   cmd.Success,
		Reason:    cmd.Reason,
	})
	return err
}

func (r *XormAuditRepository) InsertOperationAudit(cmd core.OperationAuditCreateCommand) error {
	_, err := r.DB.Insert(&model.OperationAudit{
		ActorUserId:   cmd.ActorUserID,
		ActorUsername: cmd.ActorUsername,
		Action:        cmd.Action,
		ResourceType:  cmd.ResourceType,
		ResourceId:    cmd.ResourceID,
		BeforeData:    cmd.BeforeData,
		AfterData:     cmd.AfterData,
		IP:            cmd.IP,
		UserAgent:     cmd.UserAgent,
		Result:        cmd.Result,
		ErrorMessage:  cmd.ErrorMessage,
	})
	return err
}

func (r *XormAuditRepository) InsertCompatAPIAudit(cmd core.CompatAPIAuditCreateCommand) error {
	_, err := r.DB.Insert(&model.CompatAPIAudit{
		Method:        cmd.Method,
		Path:          cmd.Path,
		ClientVersion: cmd.ClientVersion,
		RustdeskId:    cmd.RustdeskID,
		IsStub:        cmd.IsStub,
		StatusCode:    cmd.StatusCode,
		IP:            cmd.IP,
		UserAgent:     cmd.UserAgent,
	})
	return err
}
