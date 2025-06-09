package server

import (
	"context"
)

// QueueManager defines the interface for managing message queues
type QueueManager interface {
	// Create creates a new queue for a task
	Create(ctx context.Context, taskID string) (*EventQueue, error)

	// Get gets a queue for a task
	Get(ctx context.Context, taskID string) (*EventQueue, error)

	// Tap taps into an existing task's queue to create a child queue
	Tap(ctx context.Context, taskID string) (*EventQueue, error)

	// Remove removes a queue for a task
	Remove(ctx context.Context, taskID string) error
}
