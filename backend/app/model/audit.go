package model

import "time"

type Audit struct {
	Id              int       `xorm:"'id' int notnull pk autoincr"`
	ConnId          int       `xorm:"'conn_id' int index"`
	RustdeskId      string    `xorm:"'rustdesk_id' varchar(100) index"`
	PeerId          string    `xorm:"'peer_id' varchar(100) index"`
	IP              string    `xorm:"'ip' varchar(64)"`
	PeerIP          string    `xorm:"'peer_ip' varchar(64)"`
	SessionId       string    `xorm:"'session_id' varchar(100) index"`
	Peer            string    `xorm:"'peer' text"`
	Uuid            string    `xorm:"'uuid' varchar(255) index"`
	Note            string    `xorm:"'note' varchar(1024)"`
	Type            int       `xorm:"'type' tinyint"`
	Direction       string    `xorm:"'direction' varchar(32) index"`
	Status          string    `xorm:"'status' varchar(32) index"`
	ClientVersion   string    `xorm:"'client_version' varchar(64)"`
	Platform        string    `xorm:"'platform' varchar(64)"`
	Hostname        string    `xorm:"'hostname' varchar(255)"`
	Raw             string    `xorm:"'raw' text"`
	StartedAt       time.Time `xorm:"'started_at' datetime index"`
	ClosedAt        time.Time `xorm:"'closed_at' datetime index"`
	DurationSeconds int64     `xorm:"'duration_seconds' bigint"`
	CreatedAt       time.Time `xorm:"'created_at' datetime created index"`
	UpdatedAt       time.Time `xorm:"'updated_at' datetime updated"`
}

func (m *Audit) TableName() string {
	return "audit"
}

type AlarmAudit struct {
	Id         int       `xorm:"'id' int notnull pk autoincr"`
	RustdeskId string    `xorm:"'rustdesk_id' varchar(100) index"`
	PeerId     string    `xorm:"'peer_id' varchar(100) index"`
	SessionId  string    `xorm:"'session_id' varchar(100) index"`
	AlarmType  string    `xorm:"'alarm_type' varchar(100) index"`
	Severity   string    `xorm:"'severity' varchar(32) index"`
	Message    string    `xorm:"'message' varchar(1024)"`
	Raw        string    `xorm:"'raw' text"`
	CreatedAt  time.Time `xorm:"'created_at' datetime created index"`
}

func (m *AlarmAudit) TableName() string {
	return "alarm_audit"
}

type OperationAudit struct {
	Id            int       `xorm:"'id' int notnull pk autoincr"`
	ActorUserId   int       `xorm:"'actor_user_id' int index"`
	ActorUsername string    `xorm:"'actor_username' varchar(100) index"`
	Action        string    `xorm:"'action' varchar(100) index"`
	ResourceType  string    `xorm:"'resource_type' varchar(100) index"`
	ResourceId    string    `xorm:"'resource_id' varchar(100) index"`
	BeforeData    string    `xorm:"'before_data' text"`
	AfterData     string    `xorm:"'after_data' text"`
	IP            string    `xorm:"'ip' varchar(64)"`
	UserAgent     string    `xorm:"'user_agent' varchar(512)"`
	Result        string    `xorm:"'result' varchar(32) index"`
	ErrorMessage  string    `xorm:"'error_message' varchar(1024)"`
	CreatedAt     time.Time `xorm:"'created_at' datetime created index"`
}

func (m *OperationAudit) TableName() string {
	return "operation_audit"
}

type SecurityAudit struct {
	Id        int       `xorm:"'id' int notnull pk autoincr"`
	UserId    int       `xorm:"'user_id' int index"`
	Username  string    `xorm:"'username' varchar(100) index"`
	Event     string    `xorm:"'event' varchar(100) index"`
	IP        string    `xorm:"'ip' varchar(64)"`
	UserAgent string    `xorm:"'user_agent' varchar(512)"`
	Success   bool      `xorm:"'success' bool index"`
	Reason    string    `xorm:"'reason' varchar(512)"`
	CreatedAt time.Time `xorm:"'created_at' datetime created index"`
}

func (m *SecurityAudit) TableName() string {
	return "security_audit"
}

type CompatAPIAudit struct {
	Id            int       `xorm:"'id' int notnull pk autoincr"`
	Method        string    `xorm:"'method' varchar(16) index"`
	Path          string    `xorm:"'path' varchar(255) index"`
	ClientVersion string    `xorm:"'client_version' varchar(64)"`
	RustdeskId    string    `xorm:"'rustdesk_id' varchar(100) index"`
	IsStub        bool      `xorm:"'is_stub' bool index"`
	StatusCode    int       `xorm:"'status_code' int"`
	IP            string    `xorm:"'ip' varchar(64)"`
	UserAgent     string    `xorm:"'user_agent' varchar(512)"`
	CreatedAt     time.Time `xorm:"'created_at' datetime created index"`
}

func (m *CompatAPIAudit) TableName() string {
	return "compat_api_audit"
}
