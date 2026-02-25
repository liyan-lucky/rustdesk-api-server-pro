package core

// PeerListQuery is the service-level query for listing peers.
type PeerListQuery struct {
	UserID   int
	Username string
	Current  int
	PageSize int
	Status   int
}

// PeerListItem is the normalized peer view used by services and transport mappers.
type PeerListItem struct {
	RustdeskID      string
	Username        string
	Platform        string
	Hostname        string
	LoginName       string
	DeviceGroupName string
	Note            string
}

// PeerListResult is the normalized paginated result for peers.
type PeerListResult struct {
	Total int64
	Items []PeerListItem
}
