package core

type AuditConnNoteCommand struct {
	SessionID string
	Note      string
}

type AuditConnOpenCommand struct {
	ConnID        int
	RustdeskID    string
	PeerID        string
	IP            string
	PeerIP        string
	SessionID     string
	UUID          string
	Direction     string
	Status        string
	ClientVersion string
	Platform      string
	Hostname      string
	Raw           string
}

type AuditConnCloseCommand struct {
	ConnID int
}

type AuditConnSessionUpdateCommand struct {
	ConnID    int
	SessionID string
	Type      int
	Peer      string
	PeerID    string
}

type FileTransferCreateCommand struct {
	RustdeskID   string
	Info         string
	IsFile       bool
	Path         string
	FileName     string
	PeerID       string
	SessionID    string
	Type         int
	UUID         string
	Direction    string
	SizeBytes    int64
	Result       string
	ErrorMessage string
	Raw          string
}

type AuditGuidNoteUpdateCommand struct {
	Guid string
	Note string
}

type AlarmAuditCreateCommand struct {
	RustdeskID string
	PeerID     string
	SessionID  string
	AlarmType  string
	Severity   string
	Message    string
	Raw        string
}

type SecurityAuditCreateCommand struct {
	UserID    int
	Username  string
	Event     string
	IP        string
	UserAgent string
	Success   bool
	Reason    string
}

type OperationAuditCreateCommand struct {
	ActorUserID   int
	ActorUsername string
	Action        string
	ResourceType  string
	ResourceID    string
	BeforeData    string
	AfterData     string
	IP            string
	UserAgent     string
	Result        string
	ErrorMessage  string
}

type CompatAPIAuditCreateCommand struct {
	Method        string
	Path          string
	ClientVersion string
	RustdeskID    string
	IsStub        bool
	StatusCode    int
	IP            string
	UserAgent     string
}
