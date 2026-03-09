package repository

import (
	"time"

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

	return core.HeartbeatResult{ModifiedAt: time.Now().Unix()}, nil
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
