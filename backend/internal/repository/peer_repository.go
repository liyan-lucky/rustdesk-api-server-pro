package repository

import "rustdesk-api-server-pro/internal/core"

// PeerRepository defines persistence operations for peer reads/writes.
// The goal is to keep controllers/services independent from the ORM details.
type PeerRepository interface {
	ListByUser(query core.PeerListQuery) (core.PeerListResult, error)
}
