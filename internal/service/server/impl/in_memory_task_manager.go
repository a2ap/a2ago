package impl

import (
	"context"
	"fmt"
	"sync"

	"github.com/a2a4j/a2ago/internal/util"

	"github.com/a2a4j/a2ago/internal/model"
	"github.com/a2a4j/a2ago/pkg/service/server"
)

// InMemoryTaskManager is an in-memory implementation of the TaskManager interface
type InMemoryTaskManager struct {
	taskStore           server.TaskStore
	notificationConfigs map[string]*model.TaskPushNotificationConfig
	contextTaskIDs      map[string]map[string]bool
	mu                  sync.RWMutex
}

// NewInMemoryTaskManager creates a new InMemoryTaskManager
func NewInMemoryTaskManager(taskStore server.TaskStore) server.TaskManager {
	return &InMemoryTaskManager{
		taskStore:           taskStore,
		notificationConfigs: make(map[string]*model.TaskPushNotificationConfig),
		contextTaskIDs:      make(map[string]map[string]bool),
	}
}

// LoadOrCreateContext loads or creates a new task context
func (m *InMemoryTaskManager) LoadOrCreateContext(ctx context.Context, params *model.MessageSendParams) (*model.RequestContext, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	taskID := params.Message.TaskID
	if taskID == "" {
		taskID = util.GenerateUUID()
	}

	contextID := params.Message.ContextID
	if contextID == "" {
		contextID = util.GenerateUUID()
	}

	// Load or create task
	task, err := m.taskStore.Load(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to load task: %w", err)
	}

	if task == nil {
		// Create new task
		task = model.NewTask(taskID)
		task.ContextID = contextID
		task.Status = model.NewTaskStatus(model.TaskStateSubmitted)
		task.Metadata = params.Metadata
		task.Artifacts = make([]*model.TaskArtifact, 0)
		task.History = []*model.Message{params.Message}

		if err := m.taskStore.Save(ctx, task); err != nil {
			return nil, fmt.Errorf("failed to save new task: %w", err)
		}
	} else {
		// Update existing task
		taskState := task.Status.State
		if taskState == model.TaskStateCompleted || taskState == model.TaskStateFailed ||
			taskState == model.TaskStateCanceled || taskState == model.TaskStateRejected {
			// Handle as new submission (keeping history)
			taskStatus := model.NewTaskStatus(model.TaskStateSubmitted)
			if _, err := m.ApplyTaskUpdateSingle(ctx, task, taskStatus); err != nil {
				return nil, fmt.Errorf("failed to update task status: %w", err)
			}
		} else if taskState == model.TaskStateSubmitted {
			// Change state to working
			taskStatus := model.NewTaskStatus(model.TaskStateWorking)
			if _, err := m.ApplyTaskUpdateSingle(ctx, task, taskStatus); err != nil {
				return nil, fmt.Errorf("failed to update task status: %w", err)
			}
		}
	}

	// Create request context
	requestCtx := model.NewRequestContext(taskID, contextID, task)

	// Update context-task mapping
	if _, exists := m.contextTaskIDs[contextID]; !exists {
		m.contextTaskIDs[contextID] = make(map[string]bool)
	}
	m.contextTaskIDs[contextID][taskID] = true

	// Add related tasks
	relatedTasks := make([]*model.Task, 0)
	for relatedTaskID := range m.contextTaskIDs[contextID] {
		if relatedTaskID != taskID {
			if relatedTask, err := m.taskStore.Load(ctx, relatedTaskID); err == nil && relatedTask != nil {
				relatedTasks = append(relatedTasks, relatedTask)
			}
		}
	}
	requestCtx.RelatedTasks = relatedTasks

	return requestCtx, nil
}

// GetTask gets a task by its ID
func (m *InMemoryTaskManager) GetTask(ctx context.Context, taskID string) (*model.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.taskStore.Load(ctx, taskID)
}

// ApplyTaskUpdate applies a list of task updates
func (m *InMemoryTaskManager) ApplyTaskUpdate(ctx context.Context, task *model.Task, updates []model.TaskUpdate) (*model.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, update := range updates {
		var err error
		switch u := update.(type) {
		case *model.TaskStatusUpdateEvent:
			task, err = m.applyStatusUpdate(task, u)
		case *model.TaskArtifactUpdateEvent:
			task, err = m.applyArtifactUpdate(task, u)
		default:
			return nil, fmt.Errorf("unsupported task update type: %T", update)
		}
		if err != nil {
			return nil, err
		}
	}

	if err := m.taskStore.Save(ctx, task); err != nil {
		return nil, fmt.Errorf("failed to save task: %w", err)
	}

	return task, nil
}

// ApplyTaskUpdateSingle applies a single task update
func (m *InMemoryTaskManager) ApplyTaskUpdateSingle(ctx context.Context, task *model.Task, update model.TaskUpdate) (*model.Task, error) {
	return m.ApplyTaskUpdate(ctx, task, []model.TaskUpdate{update})
}

// ApplyStatusUpdate applies a status update to a task
func (m *InMemoryTaskManager) ApplyStatusUpdate(ctx context.Context, task *model.Task, event *model.TaskStatusUpdateEvent) (*model.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	updatedTask, err := m.applyStatusUpdate(task, event)
	if err != nil {
		return nil, err
	}

	if err := m.taskStore.Save(ctx, updatedTask); err != nil {
		return nil, fmt.Errorf("failed to save task: %w", err)
	}

	return updatedTask, nil
}

// ApplyArtifactUpdate applies an artifact update to a task
func (m *InMemoryTaskManager) ApplyArtifactUpdate(ctx context.Context, task *model.Task, event *model.TaskArtifactUpdateEvent) (*model.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	updatedTask, err := m.applyArtifactUpdate(task, event)
	if err != nil {
		return nil, err
	}

	if err := m.taskStore.Save(ctx, updatedTask); err != nil {
		return nil, fmt.Errorf("failed to save task: %w", err)
	}

	return updatedTask, nil
}

// RegisterTaskNotification registers a task notification config
func (m *InMemoryTaskManager) RegisterTaskNotification(ctx context.Context, config *model.TaskPushNotificationConfig) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.notificationConfigs[config.TaskID] = config
	return nil
}

// GetTaskNotification gets a task notification config
func (m *InMemoryTaskManager) GetTaskNotification(ctx context.Context, taskID string) (*model.TaskPushNotificationConfig, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	config, exists := m.notificationConfigs[taskID]
	if !exists {
		return nil, nil
	}

	return config, nil
}

// applyStatusUpdate applies a status update to a task
func (m *InMemoryTaskManager) applyStatusUpdate(task *model.Task, event *model.TaskStatusUpdateEvent) (*model.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task is nil")
	}

	if event == nil || event.Status == nil {
		return nil, fmt.Errorf("invalid status update event")
	}

	task.Status = event.Status

	// Check if the status update includes an agent message and add it to history
	if event.Status.Message != nil && event.Status.Message.Role == "agent" {
		if task.History == nil {
			task.History = make([]*model.Message, 0)
		}
		task.History = append(task.History, event.Status.Message)
	}

	return task, nil
}

// applyArtifactUpdate applies an artifact update to a task
func (m *InMemoryTaskManager) applyArtifactUpdate(task *model.Task, event *model.TaskArtifactUpdateEvent) (*model.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task is nil")
	}

	if event == nil || event.Artifact == nil {
		return nil, fmt.Errorf("invalid artifact update event")
	}

	// Convert Artifact to TaskArtifact
	taskArtifact := model.NewTaskArtifact(
		event.Artifact.ArtifactID,
		event.Artifact.Parts[0], // Assuming the first part is the main content
		event.Artifact.Metadata,
	)

	if event.Append {
		// Append artifact
		task.Artifacts = append(task.Artifacts, taskArtifact)
	} else {
		// Replace artifact
		found := false
		for i, artifact := range task.Artifacts {
			if artifact.ID == taskArtifact.ID {
				task.Artifacts[i] = taskArtifact
				found = true
				break
			}
		}
		if !found {
			task.Artifacts = append(task.Artifacts, taskArtifact)
		}
	}

	return task, nil
}
