package core

type AuditConnNoteCommand struct {
	SessionID string
	Note      string
}

type AuditConnOpenCommand struct {
	ConnID     int
	RustdeskID string
	IP         string
	SessionID  string
	UUID       string
}

type AuditConnCloseCommand struct {
	ConnID int
}

type AuditConnSessionUpdateCommand struct {
	ConnID    int
	SessionID string
	Type      int
	Peer      string
}

type FileTransferCreateCommand struct {
	RustdeskID string
	Info       string
	IsFile     bool
	Path       string
	PeerID     string
	Type       int
	UUID       string
}

type AuditGuidNoteUpdateCommand struct {
	Guid string
	Note string
}
