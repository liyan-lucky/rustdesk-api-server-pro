package model

import "time"

type DeviceGroup struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Guid      string    `xorm:"'guid' varchar(64) unique"`
	Name      string    `xorm:"'name' varchar(255)"`
	OwnerId   int       `xorm:"'owner_id' int"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *DeviceGroup) TableName() string { return "device_group" }

type DeviceGroupDevice struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	GroupGuid  string    `xorm:"'group_guid' varchar(64) index"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(255) index"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created"`
}

func (m *DeviceGroupDevice) TableName() string { return "device_group_device" }

type UserGroup struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Guid      string    `xorm:"'guid' varchar(64) unique"`
	Name      string    `xorm:"'name' varchar(255)"`
	OwnerId   int       `xorm:"'owner_id' int"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *UserGroup) TableName() string { return "user_group" }

type UserGroupMember struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	GroupGuid string    `xorm:"'group_guid' varchar(64) index"`
	UserId    int       `xorm:"'user_id' int index"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
}

func (m *UserGroupMember) TableName() string { return "user_group_member" }

type Strategy struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	Guid      string    `xorm:"'guid' varchar(64) unique"`
	Name      string    `xorm:"'name' varchar(255)"`
	Content   string    `xorm:"'content' text"`
	Enabled   bool      `xorm:"'enabled' tinyint"`
	OwnerId   int       `xorm:"'owner_id' int"`
	CreatedAt time.Time `xorm:"'created_at' datetime created"`
	UpdatedAt time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Strategy) TableName() string { return "strategy" }

type StrategyAssignment struct {
	Id           int       `xorm:"'id' int notnull pk autoincr"`
	StrategyGuid string    `xorm:"'strategy_guid' varchar(64) index"`
	TargetType   string    `xorm:"'target_type' varchar(32)"`
	TargetGuid   string    `xorm:"'target_guid' varchar(255) index"`
	CreatedAt    time.Time `xorm:"'created_at' datetime created"`
}

func (m *StrategyAssignment) TableName() string { return "strategy_assignment" }
