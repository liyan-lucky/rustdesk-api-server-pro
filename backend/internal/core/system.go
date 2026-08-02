package core

type HeartbeatCommand struct {
	RustdeskID string
	UUID       string
	ConnCount  int
	Conns      []int
	ModifiedAt int64
}

type HeartbeatResult struct {
	ModifiedAt int64
	Strategy   map[string]string
	Sysinfo    bool
	Disconnect []int
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
