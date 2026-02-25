package core

type AddressBookPeerListQuery struct {
	UserID   int
	AbGuid   string
	Current  int
	PageSize int
}

type AddressBookPeerView struct {
	ID               string
	Hash             string
	Password         string
	Username         string
	Hostname         string
	Platform         string
	Alias            string
	Tags             []string
	Note             string
	ForceAlwaysRelay bool
	RdpPort          string
	RdpUsername      string
	LoginName        string
	SameServer       bool
}

type AddressBookPeerListResult struct {
	Total int64
	Items []AddressBookPeerView
}

type AddressBookTagListQuery struct {
	UserID int
	AbGuid string
}

type AddressBookTagView struct {
	Name  string
	Color int64
}

type SharedAddressBookListQuery struct {
	Current  int
	PageSize int
}

type SharedAddressBookView struct {
	Guid  string
	Name  string
	Owner string
	Note  string
	Rule  int
}

type SharedAddressBookListResult struct {
	Total int64
	Items []SharedAddressBookView
}

type AddressBookTagAddCommand struct {
	UserID int
	AbID   int
	Name   string
	Color  int64
}

type AddressBookTagUpdateColorCommand struct {
	UserID int
	AbID   int
	Name   string
	Color  int64
}

type AddressBookTagRenameCommand struct {
	UserID int
	AbID   int
	Old    string
	New    string
}

type AddressBookTagDeleteCommand struct {
	UserID int
	AbID   int
	Names  []string
}

type AddressBookPeerCreateCommand struct {
	UserID            int
	AbID              int
	RustdeskID        string
	Hash              string
	Username          string
	Password          string
	Hostname          string
	Platform          string
	Alias             string
	Tags              []string
	Note              string
	ForceAlwaysRelay  bool
	RdpPort           string
	RdpUsername       string
	LoginName         string
	SameServerPresent bool
}

type AddressBookPeerUpdateCommand struct {
	UserID     int
	AbID       int
	RustdeskID string

	Tags     *string
	Alias    *string
	Hash     *string
	Password *string
	Note     *string
	Username *string
	Hostname *string
	Platform *string
}

type AddressBookPeerDeleteCommand struct {
	UserID int
	AbID   int
	IDs    []string
}
