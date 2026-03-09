package httpdto

import "rustdesk-api-server-pro/internal/core"

type PeerInfo struct {
	Username   string `json:"username"`
	OS         string `json:"os"`
	DeviceName string `json:"device_name"`
}

type PeerRow struct {
	ID              string   `json:"id"`
	Info            PeerInfo `json:"info"`
	Status          int      `json:"status"`
	User            string   `json:"user"`
	UserName        string   `json:"user_name"`
	DeviceGroupName string   `json:"device_group_name"`
	Note            string   `json:"note"`
}

type PeerListResponse struct {
	Total int64     `json:"total"`
	Data  []PeerRow `json:"data"`
}

func NewPeerListResponse(result core.PeerListResult, status int, username string) PeerListResponse {
	rows := make([]PeerRow, 0, len(result.Items))
	for _, p := range result.Items {
		ownerName := p.OwnerName
		if ownerName == "" {
			ownerName = username
		}
		rows = append(rows, PeerRow{
			ID: p.RustdeskID,
			Info: PeerInfo{
				Username:   p.Username,
				OS:         p.Platform,
				DeviceName: p.Hostname,
			},
			Status:          status,
			User:            ownerName,
			UserName:        p.LoginName,
			DeviceGroupName: p.DeviceGroupName,
			Note:            p.Note,
		})
	}
	return PeerListResponse{
		Total: result.Total,
		Data:  rows,
	}
}
