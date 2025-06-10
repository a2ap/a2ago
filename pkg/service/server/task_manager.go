package server

import (
	"context"

	"github.com/a2ap/a2ago/internal/model"
)

// TaskManager defines the interface for managing tasks in the A2A system.
// The TaskManager is responsible for handling the lifecycle and state of tasks.
type TaskManager interface {
	// LoadOrCreateContext loads or creates a new task context
	LoadOrCreateContext(ctx context.Context, params *model.MessageSendParams) (*model.RequestContext, error)

	// GetTask gets a task by its ID
	GetTask(ctx context.Context, taskID string) (*model.Task, error)

	// ApplyTaskUpdate applies a list of task updates
	ApplyTaskUpdate(ctx context.Context, task *model.Task, updates []model.TaskUpdate) (*model.Task, error)

	// ApplyTaskUpdate applies a single task update
	ApplyTaskUpdateSingle(ctx context.Context, task *model.Task, update model.TaskUpdate) (*model.Task, error)

	// ApplyStatusUpdate applies a status update to a task
	ApplyStatusUpdate(ctx context.Context, task *model.Task, event *model.TaskStatusUpdateEvent) (*model.Task, error)

	// ApplyArtifactUpdate applies an artifact update to a task
	ApplyArtifactUpdate(ctx context.Context, task *model.Task, event *model.TaskArtifactUpdateEvent) (*model.Task, error)

	// RegisterTaskNotification registers a task notification config
	RegisterTaskNotification(ctx context.Context, config *model.TaskPushNotificationConfig) error

	// GetTaskNotification gets a task notification config
	GetTaskNotification(ctx context.Context, taskID string) (*model.TaskPushNotificationConfig, error)
}
