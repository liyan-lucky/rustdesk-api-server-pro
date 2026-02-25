package core

type CompatLoginOptionsResult struct {
	Options []string
}

type CompatOidcAuthResult struct {
	Error   string
	Enabled bool
	URL     string
}

type CompatOidcAuthQueryResult struct {
	Error   string
	Enabled bool
	User    any
}

type CompatPluginSignResult struct {
	SignedMsg []byte
}

type CompatDevicesCliCommand struct {
	UserID     int
	RustdeskID string

	DeviceName     *string
	DeviceUsername *string
	UserName       *string
	Note           *string
	Alias          *string
	Password       *string

	Tags    []string
	HasTags bool
}

type CompatRecordCommand struct {
	Op       string
	FileName string
	Offset   int64
	Body     []byte
}
