package server

import (
	"context"

	"github.com/a2ap/a2ago/internal/model"
)

// TaskStore defines the interface for storing and retrieving tasks.
type TaskStore interface {
	// Save saves a task and its associated message history.
	// Overwrites existing data if the task ID exists.
	Save(ctx context.Context, task *model.Task) error

	// Load loads a task and its history by task ID.
	// Returns nil if not found.
	Load(ctx context.Context, taskID string) (*model.Task, error)

	// Delete removes a task by its ID.
	Delete(ctx context.Context, taskID string) error

	// ListTasks returns all tasks
	ListTasks(ctx context.Context) ([]*model.Task, error)

	//// List lists all tasks with pagination.
	//List(ctx context.Context, page, pageSize int) ([]*model.Task, error)
}

// InMemoryTaskStore 是 TaskStore 接口的内存实现
