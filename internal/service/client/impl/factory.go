package client

import (
	"github.com/a2ap/a2ago/internal/model"
	client2 "github.com/a2ap/a2ago/pkg/service/client"
)

// NewClient creates a new A2A client with the given server URL.
func NewClient(serverURL string) client2.A2aClient {
	cardResolver := NewHttpCardResolver(serverURL)
	return NewDefaultA2aClient(cardResolver)
}

// NewClientWithCard creates a new A2A client with the given server URL and agent card.
func NewClientWithCard(serverURL string, agentCard *model.AgentCard) client2.A2aClient {
	cardResolver := NewHttpCardResolver(serverURL)
	return NewDefaultA2aClientWithCard(agentCard, cardResolver)
}
