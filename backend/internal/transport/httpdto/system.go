package httpdto

import "rustdesk-api-server-pro/internal/core"

type HeartbeatResponse struct {
	ModifiedAt int64             `json:"modified_at"`
	Strategy   map[string]string `json:"strategy,omitempty"`
	Sysinfo    bool              `json:"sysinfo,omitempty"`
	Disconnect []int             `json:"disconnect,omitempty"`
}

func NewHeartbeatResponse(result core.HeartbeatResult) HeartbeatResponse {
	return HeartbeatResponse{
		ModifiedAt: result.ModifiedAt,
		Strategy:   result.Strategy,
		Sysinfo:    result.Sysinfo,
		Disconnect: result.Disconnect,
	}
}
