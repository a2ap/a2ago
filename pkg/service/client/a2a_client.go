package client

import (
	"context"

	model2 "github.com/a2a4j/a2ago/internal/model"
)

// A2aClient defines the core functionality of an A2A client.
// The A2A client is responsible for interacting with an A2A server.
type A2aClient interface {
	// AgentCard returns the AgentCard info currently in client.
	AgentCard() *model2.AgentCard

	// RetrieveAgentCard retrieves the AgentCard for the server this client connects to.
	// This is typically fetched from a well-known endpoint.
	RetrieveAgentCard() *model2.AgentCard

	// SendMessage sends a task request to the server (non-streaming).
	SendMessage(ctx context.Context, params *model2.MessageSendParams) (*model2.Task, error)

	// SendMessageStream sends a task request and subscribes to streaming updates.
	// Returns a channel that emits task update events.
	SendMessageStream(ctx context.Context, params *model2.MessageSendParams) (<-chan model2.SendStreamingMessageResponse, error)

	// GetTask retrieves the current state of a task.
	GetTask(ctx context.Context, params *model2.TaskQueryParams) (*model2.Task, error)

	// CancelTask cancels a currently running task.
	CancelTask(ctx context.Context, params *model2.TaskIdParams) (*model2.Task, error)

	// SetTaskPushNotification sets or updates the push notification config for a task.
	SetTaskPushNotification(ctx context.Context, params *model2.TaskPushNotificationConfig) (*model2.TaskPushNotificationConfig, error)

	// GetTaskPushNotification retrieves the currently configured push notification config for a task.
	GetTaskPushNotification(ctx context.Context, params *model2.TaskIdParams) (*model2.TaskPushNotificationConfig, error)

	// ResubscribeTask resubscribes to updates for a task after a potential connection interruption.
	// Returns a channel that emits task update events.
	ResubscribeTask(ctx context.Context, params *model2.TaskQueryParams) (<-chan model2.SendStreamingMessageResponse, error)

	// Supports checks if the server likely supports optional methods based on agent card.
	// This is a client-side heuristic and might not be perfectly accurate.
	Supports(capability string) bool
}

// DefaultA2aClient is a default implementation of A2aClient.
// ... existing code ...
