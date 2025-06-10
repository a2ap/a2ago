package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/a2ap/a2ago/internal/model"
)

// HttpCardResolver is the HTTP implementation of CardResolver.
type HttpCardResolver struct {
	serverURL string
	client    *http.Client
}

// NewHttpCardResolver creates a new HttpCardResolver.
func NewHttpCardResolver(serverURL string) *HttpCardResolver {
	return &HttpCardResolver{
		serverURL: serverURL,
		client:    &http.Client{},
	}
}

// ResolveCard resolves the agent card.
func (r *HttpCardResolver) ResolveCard(ctx context.Context) (*model.AgentCard, error) {
	// Use the standard well-known path
	url := fmt.Sprintf("%s/.well-known/agent.json", r.serverURL)
	log.Printf("Retrieve agent card to %s", r.serverURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Error creating request to %s: %v", r.serverURL, err)
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := r.client.Do(req)
	if err != nil {
		log.Printf("Error sending request to %s: %v", r.serverURL, err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Unexpected status code from %s: %d", r.serverURL, resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var card model.AgentCard
	if err := json.NewDecoder(resp.Body).Decode(&card); err != nil {
		log.Printf("Error decoding response from %s: %v", r.serverURL, err)
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	log.Printf("Retrieve agent card %s successfully. Info: %+v", r.serverURL, card)
	return &card, nil
}
