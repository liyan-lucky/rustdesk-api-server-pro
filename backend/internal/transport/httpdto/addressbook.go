package httpdto

import (
	"encoding/json"

	"rustdesk-api-server-pro/internal/core"
)

type AddressBookPeerRow struct {
	ID               string   `json:"id"`
	Hash             string   `json:"hash"`
	Password         string   `json:"password"`
	Username         string   `json:"username"`
	Hostname         string   `json:"hostname"`
	Platform         string   `json:"platform"`
	Alias            string   `json:"alias"`
	Tags             []string `json:"tags"`
	Note             string   `json:"note"`
	ForceAlwaysRelay string   `json:"forceAlwaysRelay"`
	RdpPort          string   `json:"rdpPort"`
	RdpUsername      string   `json:"rdpUsername"`
	LoginName        string   `json:"loginName"`
	SameServer       bool     `json:"same_server"`
}

type LegacyAddressBookPeerRow struct {
	ID       string   `json:"id"`
	Hash     string   `json:"hash"`
	Username string   `json:"username"`
	Hostname string   `json:"hostname"`
	Platform string   `json:"platform"`
	Alias    string   `json:"alias"`
	Tags     []string `json:"tags"`
	Note     string   `json:"note"`
}

type LegacyAddressBookDataPayload struct {
	Tags      []string                   `json:"tags"`
	Peers     []LegacyAddressBookPeerRow `json:"peers"`
	TagColors string                     `json:"tag_colors"`
}

type LegacyAddressBookResponse struct {
	LicensedDevices int    `json:"licensed_devices"`
	Data            string `json:"data"`
}

type PersonalAddressBookResponse struct {
	Guid string `json:"guid"`
}

type AddressBookSettingsResponse struct {
	MaxPeerOneAB int `json:"max_peer_one_ab"`
}

func NewLegacyAddressBookResponse(licensedDevices int, result core.LegacyAddressBookGetResult) (LegacyAddressBookResponse, error) {
	tagColorsJSON, err := json.Marshal(result.TagColors)
	if err != nil {
		return LegacyAddressBookResponse{}, err
	}
	peers := make([]LegacyAddressBookPeerRow, 0, len(result.Peers))
	for _, p := range result.Peers {
		peers = append(peers, LegacyAddressBookPeerRow{
			ID:       p.RustdeskID,
			Hash:     p.Hash,
			Username: p.Username,
			Hostname: p.Hostname,
			Platform: p.Platform,
			Alias:    p.Alias,
			Tags:     p.Tags,
			Note:     p.Note,
		})
	}
	dataBytes, err := json.Marshal(LegacyAddressBookDataPayload{
		Tags:      result.Tags,
		Peers:     peers,
		TagColors: string(tagColorsJSON),
	})
	if err != nil {
		return LegacyAddressBookResponse{}, err
	}
	return LegacyAddressBookResponse{
		LicensedDevices: licensedDevices,
		Data:            string(dataBytes),
	}, nil
}

func NewPersonalAddressBookResponse(result core.PersonalAddressBookEnsureResult) PersonalAddressBookResponse {
	return PersonalAddressBookResponse{Guid: result.Guid}
}

func NewAddressBookSettingsResponse(result core.AddressBookSettingsResult) AddressBookSettingsResponse {
	return AddressBookSettingsResponse{MaxPeerOneAB: result.MaxPeerOneAB}
}

type AddressBookPeerListResponse struct {
	Total int64                `json:"total"`
	Data  []AddressBookPeerRow `json:"data"`
}

func NewAddressBookPeerListResponse(result core.AddressBookPeerListResult) AddressBookPeerListResponse {
	data := make([]AddressBookPeerRow, 0, len(result.Items))
	for _, p := range result.Items {
		forceAlwaysRelay := "false"
		if p.ForceAlwaysRelay {
			forceAlwaysRelay = "true"
		}
		data = append(data, AddressBookPeerRow{
			ID:               p.ID,
			Hash:             p.Hash,
			Password:         p.Password,
			Username:         p.Username,
			Hostname:         p.Hostname,
			Platform:         p.Platform,
			Alias:            p.Alias,
			Tags:             p.Tags,
			Note:             p.Note,
			ForceAlwaysRelay: forceAlwaysRelay,
			RdpPort:          p.RdpPort,
			RdpUsername:      p.RdpUsername,
			LoginName:        p.LoginName,
			SameServer:       p.SameServer,
		})
	}
	return AddressBookPeerListResponse{Total: result.Total, Data: data}
}

type AddressBookTagRow struct {
	Name  string `json:"name"`
	Color int64  `json:"color"`
}

func NewAddressBookTagListResponse(tags []core.AddressBookTagView) []AddressBookTagRow {
	out := make([]AddressBookTagRow, 0, len(tags))
	for _, t := range tags {
		out = append(out, AddressBookTagRow{Name: t.Name, Color: t.Color})
	}
	return out
}

type SharedAddressBookProfileRow struct {
	Guid  string `json:"guid"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Note  string `json:"note"`
	Rule  int    `json:"rule"`
}

type SharedAddressBookProfileListResponse struct {
	Total int64                         `json:"total"`
	Data  []SharedAddressBookProfileRow `json:"data"`
}

func NewSharedAddressBookProfileListResponse(result core.SharedAddressBookListResult) SharedAddressBookProfileListResponse {
	out := make([]SharedAddressBookProfileRow, 0, len(result.Items))
	for _, ab := range result.Items {
		out = append(out, SharedAddressBookProfileRow{
			Guid:  ab.Guid,
			Name:  ab.Name,
			Owner: ab.Owner,
			Note:  ab.Note,
			Rule:  ab.Rule,
		})
	}
	return SharedAddressBookProfileListResponse{Total: result.Total, Data: out}
}
