package impl

import (
	"context"
	"fmt"
	"sync"

	"github.com/a2ap/a2ago/pkg/service/server"
)

// InMemoryQueueManager is an in-memory implementation of the QueueManager interface
type InMemoryQueueManager struct {
	queues map[string]*server.EventQueue
	mu     sync.RWMutex
}

// NewInMemoryQueueManager creates a new InMemoryQueueManager
func NewInMemoryQueueManager() server.QueueManager {
	return &InMemoryQueueManager{
		queues: make(map[string]*server.EventQueue),
	}
}

// Create creates a new queue for a task
func (m *InMemoryQueueManager) Create(ctx context.Context, taskID string) (*server.EventQueue, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	queue := server.NewEventQueue()
	m.queues[taskID] = queue
	//log.Printf("[QueueManager] Create: taskID=%s, queue=%p\n%s", taskID, queue, debug.Stack())
	return queue, nil
}

// Get gets a queue for a task
func (m *InMemoryQueueManager) Get(ctx context.Context, taskID string) (*server.EventQueue, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	queue, exists := m.queues[taskID]
	//log.Printf("[QueueManager] Get: taskID=%s, exists=%v, queue=%p\n%s", taskID, exists, queue, debug.Stack())
	if !exists {
		return nil, fmt.Errorf("queue not found for task ID: %s", taskID)
	}
	return queue, nil
}

// Tap taps into an existing task's queue to create a child queue
func (m *InMemoryQueueManager) Tap(ctx context.Context, taskID string) (*server.EventQueue, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	queue, exists := m.queues[taskID]
	//log.Printf("[QueueManager] Tap: taskID=%s, exists=%v, queue=%p\n%s", taskID, exists, queue, debug.Stack())
	if !exists {
		return nil, fmt.Errorf("queue not found for task ID: %s", taskID)
	}
	return queue.Tap()
}

// Remove removes a queue for a task
func (m *InMemoryQueueManager) Remove(ctx context.Context, taskID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	queue, exists := m.queues[taskID]
	//log.Printf("[QueueManager] Remove: taskID=%s, exists=%v, queue=%p\n%s", taskID, exists, queue, debug.Stack())
	if !exists {
		return nil
	}

	if err := queue.Close(); err != nil {
		return err
	}

	delete(m.queues, taskID)
	return nil
}
