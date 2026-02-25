package service

import (
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/repository"
)

type SystemService struct {
	repo repository.SystemRepository
}

func NewSystemService(repo repository.SystemRepository) *SystemService {
	return &SystemService{repo: repo}
}

func (s *SystemService) HandleHeartbeat(cmd core.HeartbeatCommand) (core.HeartbeatResult, error) {
	return s.repo.UpsertHeartbeat(cmd)
}

func (s *SystemService) UpdateSysinfo(cmd core.SysinfoUpdateCommand) (core.SysinfoUpdateResult, error) {
	return s.repo.UpdateSysinfo(cmd)
}
