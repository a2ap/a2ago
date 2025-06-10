package client

import (
	"context"

	"github.com/a2ap/a2ago/internal/model"
)

// CardResolver is the interface for resolving agent cards.
type CardResolver interface {
	// ResolveCard resolves the agent card.
	ResolveCard(ctx context.Context) (*model.AgentCard, error)
}
