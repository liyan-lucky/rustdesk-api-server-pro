package repository

import "rustdesk-api-server-pro/internal/core"

type SystemRepository interface {
	UpsertHeartbeat(cmd core.HeartbeatCommand) (core.HeartbeatResult, error)
	UpdateSysinfo(cmd core.SysinfoUpdateCommand) (core.SysinfoUpdateResult, error)
}
