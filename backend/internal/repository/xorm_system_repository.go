package repository

import (
	"encoding/json"
	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"
	"time"

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
		Cols("is_online", "conns", "updated_at").
		Update(&model.Device{IsOnline: true, Conns: cmd.ConnCount, UpdatedAt: time.Now()}); err != nil {
		return core.HeartbeatResult{}, err
	}

	result := core.HeartbeatResult{ModifiedAt: cmd.ModifiedAt}
	strategy, modifiedAt, err := r.resolveDeviceStrategy(cmd.RustdeskID)
	if err != nil {
		return core.HeartbeatResult{}, err
	}
	if modifiedAt > cmd.ModifiedAt {
		result.ModifiedAt = modifiedAt
		result.Strategy = strategy
	}
	if device.Disabled && len(cmd.Conns) > 0 {
		result.Disconnect = append([]int(nil), cmd.Conns...)
	}
	return result, nil
}

func (r *XormSystemRepository) resolveDeviceStrategy(rustdeskID string) (map[string]string, int64, error) {
	hasAssignments, err := r.DB.IsTableExist(new(model.StrategyAssignment))
	if err != nil || !hasAssignments {
		return nil, 0, err
	}
	var assignments []model.StrategyAssignment
	err = r.DB.Where("(target_type = ? AND target_guid = ?) OR (target_type = ? AND target_guid IN (SELECT group_guid FROM device_group_device WHERE rustdesk_id = ?))", "device", rustdeskID, "device_group", rustdeskID).Find(&assignments)
	if err != nil || len(assignments) == 0 {
		return nil, 0, err
	}
	guids := make([]string, 0, len(assignments))
	for _, assignment := range assignments {
		guids = append(guids, assignment.StrategyGuid)
	}
	var strategies []model.Strategy
	if err := r.DB.In("guid", guids).Where("enabled = ?", true).Asc("updated_at", "id").Find(&strategies); err != nil {
		return nil, 0, err
	}
	merged := map[string]string{}
	var modifiedAt int64
	for _, strategy := range strategies {
		var payload map[string]any
		if json.Unmarshal([]byte(strategy.Content), &payload) != nil {
			continue
		}
		options := payload
		if nested, ok := payload["config_options"].(map[string]any); ok {
			options = nested
		}
		for key, value := range options {
			switch typed := value.(type) {
			case string:
				merged[key] = typed
			case nil:
				merged[key] = ""
			default:
				encoded, _ := json.Marshal(typed)
				merged[key] = string(encoded)
			}
		}
		if ts := strategy.UpdatedAt.Unix(); ts > modifiedAt {
			modifiedAt = ts
		}
	}
	if len(merged) == 0 {
		return nil, 0, nil
	}
	return merged, modifiedAt, nil
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
