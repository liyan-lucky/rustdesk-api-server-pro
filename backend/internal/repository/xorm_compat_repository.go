package repository

import (
	"encoding/json"

	"rustdesk-api-server-pro/app/model"
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/errcode"

	"xorm.io/xorm"
)

type XormCompatRepository struct {
	DB *xorm.Engine
}

func NewXormCompatRepository(dbEngine *xorm.Engine) *XormCompatRepository {
	return &XormCompatRepository{DB: dbEngine}
}

func (r *XormCompatRepository) ApplyDevicesCli(cmd core.CompatDevicesCliCommand) error {
	var existingDevice model.Device
	hasDevice, err := r.DB.Where("rustdesk_id = ?", cmd.RustdeskID).Get(&existingDevice)
	if err != nil {
		return err
	}
	if !hasDevice {
		return errcode.New(errcode.ERR6006.Code, errcode.ERR6006.Message)
	}
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
	if cmd.StrategyName != nil {
		var strategy model.Strategy
		has, err := r.DB.Where("name = ? and enabled = ?", *cmd.StrategyName, true).Get(&strategy)
		if err != nil {
			return err
		}
		if !has {
			return errcode.New(errcode.ERR6003.Code, errcode.ERR6003.Message)
		}
		if _, err = r.DB.Where("target_type = ? and target_guid = ?", "device", cmd.RustdeskID).Delete(&model.StrategyAssignment{}); err != nil {
			return err
		}
		if _, err = r.DB.Insert(&model.StrategyAssignment{StrategyGuid: strategy.Guid, TargetType: "device", TargetGuid: cmd.RustdeskID}); err != nil {
			return err
		}
	}
	if cmd.DeviceGroupName != nil {
		var group model.DeviceGroup
		has, err := r.DB.Where("name = ?", *cmd.DeviceGroupName).Get(&group)
		if err != nil {
			return err
		}
		if !has {
			return errcode.New(errcode.ERR6001.Code, errcode.ERR6001.Message)
		}
		count, err := r.DB.Where("group_guid = ? and rustdesk_id = ?", group.Guid, cmd.RustdeskID).Count(&model.DeviceGroupDevice{})
		if err != nil {
			return err
		}
		if count == 0 {
			if _, err = r.DB.Insert(&model.DeviceGroupDevice{GroupGuid: group.Guid, RustdeskId: cmd.RustdeskID}); err != nil {
				return err
			}
		}
	}
	if cmd.AddressBookName != nil {
		var ab model.AddressBook
		has, err := r.DB.Where("name = ? and (user_id = ? or shared = ?)", *cmd.AddressBookName, cmd.UserID, true).Get(&ab)
		if err != nil {
			return err
		}
		if !has {
			return errcode.New(errcode.ERR5001.Code, errcode.ERR5001.Message)
		}
		count, err := r.DB.Where("user_id = ? and ab_id = ? and rustdesk_id = ?", ab.UserId, ab.Id, cmd.RustdeskID).Count(&model.Peer{})
		if err != nil {
			return err
		}
		if count == 0 {
			if _, err = r.DB.Insert(&model.Peer{UserId: ab.UserId, AbId: ab.Id, RustdeskId: cmd.RustdeskID, Hostname: existingDevice.Hostname, Username: existingDevice.Username, Platform: existingDevice.Os, Tags: "[]"}); err != nil {
				return err
			}
		}
	}
	return nil
}
