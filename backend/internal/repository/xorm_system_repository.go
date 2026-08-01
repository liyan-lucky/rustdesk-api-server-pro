package repository

import (
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"

	"xorm.io/xorm"
)

type XormSystemRepository struct {
	DB *xorm.Engine
}

func NewXormSystemRepository(dbEngine *xorm.Engine) *XormSystemRepository {
	return &XormSystemRepository{DB: dbEngine}
}

func (r *XormSystemRepository) UpsertHeartbeat(cmd core.HeartbeatCommand) (core.HeartbeatResult, error) {
	var device model.Device
	has, err := r.DB.Where("rustdesk_id = ?", cmd.RustdeskID).Get(&device)
	if err != nil {
		return core.HeartbeatResult{}, err
	}

	if !has {
		device = model.Device{
			RustdeskId: cmd.RustdeskID,
			Uuid:       cmd.UUID,
			Conns:      cmd.ConnCount,
			IsOnline:   true,
		}
		if _, err := r.DB.Insert(&device); err != nil {
			return core.HeartbeatResult{}, err
		}
	}

	if _, err := r.DB.Where("rustdesk_id = ?", cmd.RustdeskID).
		Cols("is_online", "conns").
		Update(&model.Device{IsOnline: true, Conns: cmd.ConnCount}); err != nil {
		return core.HeartbeatResult{}, err
	}

	// Until a strategy is actually assigned, echo the client's strategy
	// timestamp. Returning wall-clock time here makes every heartbeat look like
	// a strategy change and causes a permanent resynchronization loop.
	return core.HeartbeatResult{ModifiedAt: cmd.ModifiedAt}, nil
}

func (r *XormSystemRepository) UpdateSysinfo(cmd core.SysinfoUpdateCommand) (core.SysinfoUpdateResult, error) {
	var device model.Device
	has, err := r.DB.Where("rustdesk_id = ?", cmd.RustdeskID).Get(&device)
	if err != nil {
		return core.SysinfoUpdateResult{}, err
	}
	if !has {
		return core.SysinfoUpdateResult{Updated: false}, nil
	}

	device.Cpu = cmd.CPU
	device.Hostname = cmd.Hostname
	device.RustdeskId = cmd.RustdeskID
	device.Memory = cmd.Memory
	device.Os = cmd.OS
	device.Username = cmd.Username
	device.Uuid = cmd.UUID
	device.Version = cmd.Version

	_, err = r.DB.Where("id = ?", device.Id).Update(&device)
	if err != nil {
		return core.SysinfoUpdateResult{}, err
	}
	return core.SysinfoUpdateResult{Updated: true}, nil
}
