package repository

import "rustdesk-api-server-pro/internal/core"

type AuditRepository interface {
	UpdateAuditConnNote(cmd core.AuditConnNoteCommand) error
	InsertAuditConnOpen(cmd core.AuditConnOpenCommand) error
	CloseAuditConn(cmd core.AuditConnCloseCommand) error
	UpdateAuditConnSession(cmd core.AuditConnSessionUpdateCommand) error

	InsertFileTransfer(cmd core.FileTransferCreateCommand) error
	UpdateAuditNoteByGuid(cmd core.AuditGuidNoteUpdateCommand) error

	InsertAlarmAudit(cmd core.AlarmAuditCreateCommand) error
	InsertSecurityAudit(cmd core.SecurityAuditCreateCommand) error
	InsertOperationAudit(cmd core.OperationAuditCreateCommand) error
	InsertCompatAPIAudit(cmd core.CompatAPIAuditCreateCommand) error
}
