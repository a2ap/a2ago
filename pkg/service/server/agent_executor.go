package server

import (
	"context"

	"github.com/a2ap/a2ago/internal/model"
)

// AgentExecutor defines the interface for executing tasks on agents.
type AgentExecutor interface {
	// Execute executes a task on an agent.
	Execute(ctx context.Context, task *model.Task, queue *EventQueue) error

	// Cancel cancels a task.
	Cancel(ctx context.Context, taskID string) error

	// GetTaskStatus gets the status of a task.
	GetTaskStatus(ctx context.Context, taskID string) (*model.TaskStatus, error)

	// GetTaskArtifact gets an artifact from a task.
	GetTaskArtifact(ctx context.Context, taskID string, artifactID string) (*model.Artifact, error)

	// ListTaskArtifacts lists all artifacts for a task.
	ListTaskArtifacts(ctx context.Context, taskID string) ([]*model.Artifact, error)

	// RegisterTaskNotification registers a task notification.
	RegisterTaskNotification(ctx context.Context, config *model.TaskPushNotificationConfig) error

	// GetTaskNotification gets a task notification.
	GetTaskNotification(ctx context.Context, taskID string) (*model.TaskPushNotificationConfig, error)
}
