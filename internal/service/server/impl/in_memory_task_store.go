package impl

import (
	"context"
	"sync"

	"github.com/a2ap/a2ago/pkg/service/server"

	"github.com/a2ap/a2ago/internal/model"
)

// InMemoryTaskStore 是 TaskStore 接口的内存实现
type InMemoryTaskStore struct {
	tasks map[string]*model.Task
	mu    sync.RWMutex
}

// NewInMemoryTaskStore 创建一个新的 InMemoryTaskStore
func NewInMemoryTaskStore() server.TaskStore {
	return &InMemoryTaskStore{
		tasks: make(map[string]*model.Task),
	}
}

// Save 保存任务
func (s *InMemoryTaskStore) Save(ctx context.Context, task *model.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.ID] = task
	return nil
}

// Load 加载任务
func (s *InMemoryTaskStore) Load(ctx context.Context, taskID string) (*model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[taskID]
	if !exists {
		return nil, nil
	}

	return task, nil
}

// DeleteTask 删除任务
func (s *InMemoryTaskStore) Delete(ctx context.Context, taskID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tasks, taskID)
	return nil
}
