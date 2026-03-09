package httpdto

import "rustdesk-api-server-pro/internal/core"

type HeartbeatResponse struct {
	ModifiedAt int64 `json:"modified_at"`
}

func NewHeartbeatResponse(result core.HeartbeatResult) HeartbeatResponse {
	return HeartbeatResponse{ModifiedAt: result.ModifiedAt}
}
