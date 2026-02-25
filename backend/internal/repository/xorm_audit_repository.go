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
	_, err := r.DB.Insert(&model.Audit{
		ConnId:     cmd.ConnID,
		RustdeskId: cmd.RustdeskID,
		IP:         cmd.IP,
		SessionId:  cmd.SessionID,
		Uuid:       cmd.UUID,
	})
	return err
}

func (r *XormAuditRepository) CloseAuditConn(cmd core.AuditConnCloseCommand) error {
	_, err := r.DB.Where("conn_id = ?", cmd.ConnID).Update(&model.Audit{ClosedAt: time.Now()})
	return err
}

func (r *XormAuditRepository) UpdateAuditConnSession(cmd core.AuditConnSessionUpdateCommand) error {
	_, err := r.DB.Where("conn_id = ?", cmd.ConnID).Update(&model.Audit{
		SessionId: cmd.SessionID,
		Type:      cmd.Type,
		Peer:      cmd.Peer,
	})
	return err
}

func (r *XormAuditRepository) InsertFileTransfer(cmd core.FileTransferCreateCommand) error {
	_, err := r.DB.Insert(&model.FileTransfer{
		RustdeskId: cmd.RustdeskID,
		Info:       cmd.Info,
		IsFile:     cmd.IsFile,
		Path:       cmd.Path,
		PeerId:     cmd.PeerID,
		Type:       cmd.Type,
		Uuid:       cmd.UUID,
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
