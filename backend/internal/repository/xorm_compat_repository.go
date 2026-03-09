package repository

import (
	"encoding/json"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"

	"xorm.io/xorm"
)

type XormCompatRepository struct {
	DB *xorm.Engine
}

func NewXormCompatRepository(dbEngine *xorm.Engine) *XormCompatRepository {
	return &XormCompatRepository{DB: dbEngine}
}

func (r *XormCompatRepository) ApplyDevicesCli(cmd core.CompatDevicesCliCommand) error {
	device := model.Device{}
	deviceCols := make([]string, 0, 2)
	if cmd.DeviceName != nil {
		device.Hostname = *cmd.DeviceName
		deviceCols = append(deviceCols, "hostname")
	}
	if cmd.DeviceUsername != nil {
		device.Username = *cmd.DeviceUsername
		deviceCols = append(deviceCols, "username")
	}
	if len(deviceCols) > 0 {
		if _, err := r.DB.Where("rustdesk_id = ?", cmd.RustdeskID).Cols(deviceCols...).Update(&device); err != nil {
			return err
		}
	}

	peer := model.Peer{}
	peerCols := make([]string, 0, 8)
	appendCol := func(col string) {
		for _, c := range peerCols {
			if c == col {
				return
			}
		}
		peerCols = append(peerCols, col)
	}

	if cmd.Note != nil {
		peer.Note = *cmd.Note
		appendCol("note")
	}
	if cmd.UserName != nil {
		peer.LoginName = *cmd.UserName
		appendCol("loginName")
	}
	if cmd.DeviceName != nil {
		peer.Hostname = *cmd.DeviceName
		appendCol("hostname")
	}
	if cmd.DeviceUsername != nil {
		peer.Username = *cmd.DeviceUsername
		appendCol("username")
	}
	if cmd.Alias != nil {
		peer.Alias = *cmd.Alias
		appendCol("alias")
	}
	if cmd.Password != nil {
		peer.Password = *cmd.Password
		appendCol("password")
	}
	if cmd.HasTags {
		tagsJSON := "[]"
		if b, err := json.Marshal(cmd.Tags); err == nil {
			tagsJSON = string(b)
		}
		peer.Tags = tagsJSON
		appendCol("tags")
	}
	if len(peerCols) > 0 {
		if _, err := r.DB.Where("user_id = ? and rustdesk_id = ?", cmd.UserID, cmd.RustdeskID).Cols(peerCols...).Update(&peer); err != nil {
			return err
		}
	}
	return nil
}
