package impl

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/a2ap/a2ago/internal/model"
	"github.com/a2ap/a2ago/pkg/service/server"
)

// DefaultA2AServer implements the A2AServer interface
type DefaultA2AServer struct {
	taskManager   server.TaskManager
	queueManager  server.QueueManager
	agentExecutor server.AgentExecutor
	agentCard     *model.AgentCard
}

// NewDefaultA2AServer creates a new instance of DefaultA2AServer
func NewDefaultA2AServer(taskManager server.TaskManager, queueManager server.QueueManager, agentExecutor server.AgentExecutor, agentCard *model.AgentCard) server.A2AServer {
	return &DefaultA2AServer{
		taskManager:   taskManager,
		queueManager:  queueManager,
		agentExecutor: agentExecutor,
		agentCard:     agentCard,
	}
}

// HandleMessage handles a message request
func (s *DefaultA2AServer) HandleMessage(ctx context.Context, params *model.MessageSendParams) (*model.SendMessageResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Message == nil {
		return nil, fmt.Errorf("message cannot be nil")
	}

	// Load or create task context
	taskCtx, err := s.taskManager.LoadOrCreateContext(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to load or create task context: %w", err)
	}

	// Create queue
	queue, err := s.queueManager.Create(ctx, taskCtx.TaskID)
	if err != nil {
		return nil, fmt.Errorf("failed to create queue: %w", err)
	}

	// Send message
	message := &model.Message{
		TaskID:    taskCtx.TaskID,
		ContextID: taskCtx.ContextID,
		Parts:     params.Message.Parts,
		Role:      params.Message.Role,
		Kind:      params.Message.Kind,
		Metadata:  params.Message.Metadata,
	}
	if err := queue.EnqueueEvent(message); err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	// Execute task and handle events
	eventChan := queue.AsFlux()

	// 聚合响应内容
	var (
		finalStatus *model.TaskStatus
		artifactMap = make(map[string]*model.Artifact)
		history     []*model.Message
	)

	// Start goroutine to handle events
	done := make(chan struct{})
	go func() {
		for event := range eventChan {
			switch e := event.(type) {
			case *model.TaskStatusUpdateEvent:
				updatedTask, err := s.taskManager.ApplyStatusUpdate(ctx, taskCtx.Task, e)
				if err != nil {
					log.Printf("Error applying status update for task %s: %v", taskCtx.TaskID, err)
					continue
				}
				taskCtx.Task = updatedTask
				finalStatus = e.Status
				if e.Status != nil && e.Status.Message != nil {
					history = append(history, e.Status.Message)
				}
			case *model.TaskArtifactUpdateEvent:
				updatedTask, err := s.taskManager.ApplyArtifactUpdate(ctx, taskCtx.Task, e)
				if err != nil {
					log.Printf("Error applying artifact update for task %s: %v", taskCtx.TaskID, err)
					continue
				}
				taskCtx.Task = updatedTask
				if e.Artifact != nil {
					id := e.Artifact.ArtifactID
					if existing, ok := artifactMap[id]; ok {
						existing.Parts = append(existing.Parts, e.Artifact.Parts...)
					} else {
						artifactMap[id] = e.Artifact
					}
				}
			case *model.Task:
				continue
			case *model.Message:
				history = append(history, e)
			default:
				log.Printf("Unknown event type for task %s: %T", taskCtx.TaskID, e)
			}
		}
		done <- struct{}{}
	}()

	// Execute task
	if err := s.agentExecutor.Execute(ctx, taskCtx.Task, queue); err != nil {
		log.Printf("Error executing task %s: %v", taskCtx.TaskID, err)
		s.queueManager.Remove(ctx, taskCtx.TaskID)
		queue.Close()
		return nil, fmt.Errorf("failed to execute task: %w", err)
	}

	queue.Close()
	s.queueManager.Remove(ctx, taskCtx.TaskID)
	<-done // 等待事件聚合完成

	// artifacts 合并后输出
	artifacts := make([]*model.Artifact, 0, len(artifactMap))
	for _, a := range artifactMap {
		artifacts = append(artifacts, a)
	}

	resp := &model.StandardSendMessageResponse{
		TaskID:    taskCtx.TaskID,
		ContextID: taskCtx.Task.ContextID,
		Status:    finalStatus,
		Artifacts: artifacts,
		History:   history,
	}

	log.Printf("Handle message success: %+v", resp)
	var result model.SendMessageResponse = resp
	return &result, nil
}

// HandleMessageStream handles a streaming message request
func (s *DefaultA2AServer) HandleMessageStream(ctx context.Context, params *model.MessageSendParams) (<-chan *model.SendStreamingMessageResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Message == nil {
		return nil, fmt.Errorf("message cannot be nil")
	}

	// Load or create task context
	taskCtx, err := s.taskManager.LoadOrCreateContext(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to load or create task context: %w", err)
	}

	// Create queue
	queue, err := s.queueManager.Create(ctx, taskCtx.TaskID)
	if err != nil {
		return nil, fmt.Errorf("failed to create queue: %w", err)
	}

	// Create response channel
	responseChan := make(chan *model.SendStreamingMessageResponse)

	// Start goroutine to handle streaming
	go func() {
		defer func() {
			// Clean up queue when done
			if err := s.queueManager.Remove(ctx, taskCtx.TaskID); err != nil {
				log.Printf("Error removing queue for task %s: %v", taskCtx.TaskID, err)
			}
			close(responseChan)
		}()

		// Send initial response
		initialMessage := &model.Message{
			TaskID: taskCtx.TaskID,
			Parts:  params.Message.Parts,
		}
		var initialResponse model.SendStreamingMessageResponse = initialMessage
		responseChan <- &initialResponse

		// Send message
		if err := queue.EnqueueEvent(initialMessage); err != nil {
			errorMessage := &model.Message{
				TaskID: taskCtx.TaskID,
				Parts: []model.Part{
					model.NewTextPart(fmt.Sprintf("Error: %v", err)),
				},
			}
			var errorResponse model.SendStreamingMessageResponse = errorMessage
			responseChan <- &errorResponse
			return
		}

		// Execute task and handle events
		if err := s.agentExecutor.Execute(ctx, taskCtx.Task, queue); err != nil {
			errorMessage := &model.Message{
				TaskID: taskCtx.TaskID,
				Parts: []model.Part{
					model.NewTextPart(fmt.Sprintf("Error executing task: %v", err)),
				},
			}
			var errorResponse model.SendStreamingMessageResponse = errorMessage
			responseChan <- &errorResponse
			return
		}

		// *** IMPORTANT: Close the queue after agent execution to signal end of stream ***
		queue.Close()

		// Handle events from queue
		for event := range queue.AsFlux() {
			switch e := event.(type) {
			case *model.TaskStatusUpdateEvent:
				// Handle status update
				updatedTask, err := s.taskManager.ApplyStatusUpdate(ctx, taskCtx.Task, e)
				if err != nil {
					log.Printf("Error applying status update for task %s: %v", taskCtx.TaskID, err)
					continue
				}
				taskCtx.Task = updatedTask
				var response model.SendStreamingMessageResponse = updatedTask
				responseChan <- &response

			case *model.TaskArtifactUpdateEvent:
				// Handle artifact update
				updatedTask, err := s.taskManager.ApplyArtifactUpdate(ctx, taskCtx.Task, e)
				if err != nil {
					log.Printf("Error applying artifact update for task %s: %v", taskCtx.TaskID, err)
					continue
				}
				taskCtx.Task = updatedTask
				var response model.SendStreamingMessageResponse = updatedTask
				responseChan <- &response

			case *model.Message:
				// Handle message events
				var response model.SendStreamingMessageResponse = e
				responseChan <- &response

			case *model.Task:
				// Skip task events
				continue

			default:
				// Skip unknown events
				log.Printf("Unknown event type for task %s: %T", taskCtx.TaskID, e)
			}
		}

		log.Printf("Task %s updates stream completed via handleMessageStream", taskCtx.TaskID)
	}()

	return responseChan, nil
}

// GetTask gets a task by ID
func (s *DefaultA2AServer) GetTask(ctx context.Context, taskID string) (*model.Task, error) {
	return s.taskManager.GetTask(ctx, taskID)
}

// CancelTask cancels a task
func (s *DefaultA2AServer) CancelTask(ctx context.Context, taskID string) (*model.Task, error) {
	// Get task
	task, err := s.taskManager.GetTask(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	if task == nil {
		log.Printf("Task with ID %s not found for cancellation", taskID)
		return nil, fmt.Errorf("task with ID %s not found for cancellation", taskID)
	}

	// Create task status with explicit timestamp
	taskStatus := &model.TaskStatus{
		State:     model.TaskStateCanceled,
		Timestamp: fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	// Get event queue if exists
	queue, err := s.queueManager.Get(ctx, taskID)
	if err != nil {
		log.Printf("Error getting queue for task %s: %v", taskID, err)
	} else if queue != nil {
		// Create status update event with timestamp
		statusUpdate := &model.TaskStatusUpdateEvent{
			TaskID: taskID,
			Status: taskStatus,
			Final:  true,
		}

		// Send event to queue
		if err := queue.EnqueueEvent(statusUpdate); err != nil {
			log.Printf("Error sending cancel event to queue for task %s: %v", taskID, err)
		}

		// Close queue
		if err := queue.Close(); err != nil {
			log.Printf("Error closing queue for task %s: %v", taskID, err)
		}
	}

	// Execute cancellation
	if err := s.agentExecutor.Cancel(ctx, taskID); err != nil {
		log.Printf("Error cancelling task %s: %v", taskID, err)
		return nil, fmt.Errorf("failed to cancel task: %w", err)
	}

	log.Printf("Task %s cancelled successfully", taskID)
	return task, nil
}

// SetTaskPushNotification sets the push notification configuration for a task
func (s *DefaultA2AServer) SetTaskPushNotification(ctx context.Context, taskID string, config *model.TaskPushNotificationConfig) (*model.TaskPushNotificationConfig, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetTaskPushNotification gets the push notification configuration for a task
func (s *DefaultA2AServer) GetTaskPushNotification(ctx context.Context, taskID string) (*model.TaskPushNotificationConfig, error) {
	return nil, fmt.Errorf("not implemented")
}

// SubscribeToTaskUpdates subscribes to task updates
func (s *DefaultA2AServer) SubscribeToTaskUpdates(ctx context.Context, taskID string) (<-chan *model.SendStreamingMessageResponse, error) {
	// Get task
	task, err := s.taskManager.GetTask(ctx, taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	if task == nil {
		log.Printf("Task with ID %s not found for subscription", taskID)
		return nil, fmt.Errorf("task with ID %s not found for subscription", taskID)
	}

	// Get event queue
	queue, err := s.queueManager.Get(ctx, taskID)
	if err != nil || queue == nil {
		log.Printf("Error getting queue for task %s: %v", taskID, err)
		return nil, fmt.Errorf("failed to get queue for task %s: %w", taskID, err)
	}

	responseChan := make(chan *model.SendStreamingMessageResponse)

	go func() {
		defer close(responseChan)

		log.Printf("Subscriber attached to task %s updates via subscribeToTaskUpdates", taskID)

		// Forward events from queue to response channel
		for event := range queue.AsFlux() {
			// Convert event to SendStreamingMessageResponse
			var response model.SendStreamingMessageResponse
			switch e := event.(type) {
			case *model.TaskStatusUpdateEvent:
				response = e
			case *model.TaskArtifactUpdateEvent:
				response = e
			case *model.Message:
				response = e
			case *model.Task:
				response = e
			default:
				log.Printf("Unknown event type for task %s: %T", taskID, e)
				continue
			}

			// Send response
			responseChan <- &response
		}
	}()

	return responseChan, nil
}

// GetSelfAgentCard retrieves the AgentCard for this server
func (s *DefaultA2AServer) GetSelfAgentCard() *model.AgentCard {
	return s.agentCard
}

// GetAuthenticatedExtendedCard retrieves the authenticated extended AgentCard for the server.
// By default, this returns the same card as getSelfAgentCard().
// Subclasses can override this method to provide more detailed information
// for authenticated clients.
func (s *DefaultA2AServer) GetAuthenticatedExtendedCard(ctx context.Context) (*model.AgentCard, error) {
	// TODO: Implement logic to return a more detailed AgentCard for authenticated clients.
	// This may involve:
	// 1. Checking the authenticated principal (if applicable through security context).
	// 2. Loading a different AgentCard configuration or modifying the existing one.
	// 3. Adding skills or details that are not available in the public AgentCard.
	// For now, it returns the same card as getSelfAgentCard().
	return s.agentCard, nil
}

// ListTasks returns all tasks
func (s *DefaultA2AServer) ListTasks(ctx context.Context) ([]*model.Task, error) {
	return s.taskManager.ListTasks(ctx)
}

