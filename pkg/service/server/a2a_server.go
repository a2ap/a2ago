package server

import (
	"context"

	"github.com/a2a4j/a2ago/internal/model"
)

// A2AServer defines the interface for an A2A server
type A2AServer interface {
	// HandleMessage handles a message request
	HandleMessage(ctx context.Context, params *model.MessageSendParams) (*model.SendMessageResponse, error)

	// HandleMessageStream handles a streaming message request
	HandleMessageStream(ctx context.Context, params *model.MessageSendParams) (<-chan *model.SendStreamingMessageResponse, error)

	// GetTask gets a task by ID
	GetTask(ctx context.Context, taskID string) (*model.Task, error)

	// CancelTask cancels a task
	CancelTask(ctx context.Context, taskID string) (*model.Task, error)

	// SetTaskPushNotification sets the push notification configuration for a task
	SetTaskPushNotification(ctx context.Context, taskID string, config *model.TaskPushNotificationConfig) (*model.TaskPushNotificationConfig, error)

	// GetTaskPushNotification gets the push notification configuration for a task
	GetTaskPushNotification(ctx context.Context, taskID string) (*model.TaskPushNotificationConfig, error)

	// SubscribeToTaskUpdates subscribes to task updates
	SubscribeToTaskUpdates(ctx context.Context, taskID string) (<-chan *model.SendStreamingMessageResponse, error)

	// GetSelfAgentCard retrieves the AgentCard for this server
	GetSelfAgentCard() *model.AgentCard
}
