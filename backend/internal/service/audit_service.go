package service

import (
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/repository"
)

type AuditService struct {
	repo repository.AuditRepository
}

func NewAuditService(repo repository.AuditRepository) *AuditService {
	return &AuditService{repo: repo}
}

func (s *AuditService) UpdateConnNote(cmd core.AuditConnNoteCommand) error {
	return s.repo.UpdateAuditConnNote(cmd)
}

func (s *AuditService) OpenConn(cmd core.AuditConnOpenCommand) error {
	return s.repo.InsertAuditConnOpen(cmd)
}

func (s *AuditService) CloseConn(cmd core.AuditConnCloseCommand) error {
	return s.repo.CloseAuditConn(cmd)
}

func (s *AuditService) UpdateConnSession(cmd core.AuditConnSessionUpdateCommand) error {
	return s.repo.UpdateAuditConnSession(cmd)
}

func (s *AuditService) CreateFileTransfer(cmd core.FileTransferCreateCommand) error {
	return s.repo.InsertFileTransfer(cmd)
}

func (s *AuditService) UpdateNoteByGuid(cmd core.AuditGuidNoteUpdateCommand) error {
	return s.repo.UpdateAuditNoteByGuid(cmd)
}
