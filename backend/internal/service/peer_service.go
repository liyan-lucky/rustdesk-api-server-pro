package service

import (
	"rustdesk-api-server-pro/internal/core"
	"rustdesk-api-server-pro/internal/repository"
)

// PeerService contains peer-related use cases and keeps controller logic thin.
type PeerService struct {
	peers repository.PeerRepository
}

func NewPeerService(peers repository.PeerRepository) *PeerService {
	return &PeerService{peers: peers}
}

func (s *PeerService) ListPeers(query core.PeerListQuery) (core.PeerListResult, error) {
	return s.peers.ListByUser(query)
}
