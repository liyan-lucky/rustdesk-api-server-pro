package core

type HeartbeatCommand struct {
	RustdeskID string
	UUID       string
	ConnCount  int
}

type HeartbeatResult struct {
	ModifiedAt int64
}

type SysinfoUpdateCommand struct {
	RustdeskID string
	CPU        string
	Hostname   string
	Memory     string
	OS         string
	Username   string
	UUID       string
	Version    string
}

type SysinfoUpdateResult struct {
	Updated bool
}
