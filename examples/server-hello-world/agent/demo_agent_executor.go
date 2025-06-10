package agent

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/a2ap/a2ago/internal/model"
	"github.com/a2ap/a2ago/pkg/service/server"
)

// DemoAgentExecutor implements the AgentExecutor interface
type DemoAgentExecutor struct {
	queueManager server.QueueManager
}

// NewDemoAgentExecutor creates a new instance of DemoAgentExecutor
func NewDemoAgentExecutor(queueManager server.QueueManager) server.AgentExecutor {
	return &DemoAgentExecutor{
		queueManager: queueManager,
	}
}

// Execute implements the agent's main logic for processing user requests
func (e *DemoAgentExecutor) Execute(ctx context.Context, task *model.Task, queue *server.EventQueue) error {
	taskID := task.ID
	log.Printf("Demo agent starting execution for task: %s", taskID)

	// 1. Send task start status
	if err := e.sendWorkingStatus(queue, taskID, "Starting to process user request..."); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	// 2. Send analysis phase status
	if err := e.sendWorkingStatus(queue, taskID, "Analyzing user input..."); err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	// 3. Send processing progress status
	if err := e.sendWorkingStatus(queue, taskID, "Generating response..."); err != nil {
		return err
	}
	time.Sleep(800 * time.Millisecond)

	// 4. Send first text artifact (chunk)
	if err := e.sendTextArtifact(queue, taskID, "text-response", "AI Assistant Response",
		"Here's my analysis of your question:\n\n", false, false); err != nil {
		return err
	}
	time.Sleep(300 * time.Millisecond)

	// 5. Continue sending text artifact (chunk)
	if err := e.sendTextArtifact(queue, taskID, "text-response", "AI Assistant Response",
		"Based on the information provided, I suggest the following approach:\n", true, false); err != nil {
		return err
	}
	time.Sleep(500 * time.Millisecond)

	// 6. Send code artifact
	if err := e.sendCodeArtifact(queue, taskID); err != nil {
		return err
	}
	time.Sleep(400 * time.Millisecond)

	// 7. Complete text artifact (last chunk)
	if err := e.sendTextArtifact(queue, taskID, "text-response", "AI Assistant Response",
		"\n\nIf you have any questions, please feel free to ask!", true, true); err != nil {
		return err
	}
	time.Sleep(300 * time.Millisecond)

	// 8. Send summary artifact
	if err := e.sendSummaryArtifact(queue, taskID); err != nil {
		return err
	}
	time.Sleep(200 * time.Millisecond)

	// 9. Send final completion status
	if err := e.sendCompletedStatus(queue, taskID); err != nil {
		return err
	}

	log.Printf("Demo agent completed task: %s", taskID)
	return nil
}

// Cancel handles task cancellation requests
func (e *DemoAgentExecutor) Cancel(ctx context.Context, taskID string) error {
	log.Printf("Demo agent cancelling task: %s", taskID)
	// TODO: Implement cancellation logic
	return nil
}

// GetTaskStatus gets the status of a task
func (e *DemoAgentExecutor) GetTaskStatus(ctx context.Context, taskID string) (*model.TaskStatus, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetTaskArtifact retrieves a task artifact
func (e *DemoAgentExecutor) GetTaskArtifact(ctx context.Context, taskID string, artifactID string) (*model.Artifact, error) {
	return nil, fmt.Errorf("not implemented")
}

// ListTaskArtifacts lists all artifacts for a task
func (e *DemoAgentExecutor) ListTaskArtifacts(ctx context.Context, taskID string) ([]*model.Artifact, error) {
	return nil, fmt.Errorf("not implemented")
}

// RegisterTaskNotification registers a task notification
func (e *DemoAgentExecutor) RegisterTaskNotification(ctx context.Context, config *model.TaskPushNotificationConfig) error {
	return fmt.Errorf("not implemented")
}

// GetTaskNotification retrieves task notification configuration
func (e *DemoAgentExecutor) GetTaskNotification(ctx context.Context, taskID string) (*model.TaskPushNotificationConfig, error) {
	return nil, fmt.Errorf("not implemented")
}

// sendWorkingStatus sends a working status update
func (e *DemoAgentExecutor) sendWorkingStatus(queue *server.EventQueue, taskID, statusMessage string) error {
	event := &model.TaskStatusUpdateEvent{
		TaskID: taskID,
		Status: &model.TaskStatus{
			State: model.TaskStateWorking,
			Message: &model.Message{
				Role: "agent",
				Parts: []model.Part{
					model.NewTextPart(statusMessage),
				},
			},
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}
	return queue.EnqueueEvent(event)
}

// sendTextArtifact sends a text artifact
func (e *DemoAgentExecutor) sendTextArtifact(queue *server.EventQueue, taskID, artifactID, name, content string, append, lastChunk bool) error {
	event := &model.TaskArtifactUpdateEvent{
		TaskID: taskID,
		Artifact: &model.Artifact{
			ArtifactID:  artifactID,
			Name:        name,
			Description: "AI generated text reply",
			Parts: []model.Part{
				model.NewTextPart(content),
			},
			Metadata: map[string]interface{}{
				"contentType": "text/plain",
				"encoding":    "utf-8",
				"chunkIndex":  time.Now().UnixMilli(),
			},
		},
	}
	return queue.EnqueueEvent(event)
}

// sendCodeArtifact sends a code artifact
func (e *DemoAgentExecutor) sendCodeArtifact(queue *server.EventQueue, taskID string) error {
	code := `// Example code
package main

import "fmt"

func main() {
    fmt.Println("Hello, A2A!")
}`
	event := &model.TaskArtifactUpdateEvent{
		TaskID: taskID,
		Artifact: &model.Artifact{
			ArtifactID:  "code-example",
			Name:        "Example Code",
			Description: "Example Go code generated based on requirements",
			Parts: []model.Part{
				model.NewTextPart(code),
			},
			Metadata: map[string]interface{}{
				"contentType": "text/x-go-source",
				"language":    "go",
				"filename":    "main.go",
			},
		},
	}
	return queue.EnqueueEvent(event)
}

// sendSummaryArtifact sends a summary artifact
func (e *DemoAgentExecutor) sendSummaryArtifact(queue *server.EventQueue, taskID string) error {
	summary := "## Task Execution Summary\n\n✅ User request analysis completed\n✅ Text response generated\n✅ Example code provided\n✅ Task executed successfully\n\nTotal execution time: ~3 seconds\nGenerated content: Text response + Code example"
	event := &model.TaskArtifactUpdateEvent{
		TaskID: taskID,
		Artifact: &model.Artifact{
			ArtifactID:  "task-summary",
			Name:        "Task Summary",
			Description: "Summary report of this task execution",
			Parts: []model.Part{
				model.NewTextPart(summary),
			},
			Metadata: map[string]interface{}{
				"reportType":  "summary",
				"contentType": "text/markdown",
			},
		},
	}
	return queue.EnqueueEvent(event)
}

// sendCompletedStatus sends a completed status
func (e *DemoAgentExecutor) sendCompletedStatus(queue *server.EventQueue, taskID string) error {
	event := &model.TaskStatusUpdateEvent{
		TaskID: taskID,
		Status: &model.TaskStatus{
			State: model.TaskStateCompleted,
			Message: &model.Message{
				Role: "agent",
				Parts: []model.Part{
					model.NewTextPart("Task completed successfully! I have generated a detailed response and example code for you."),
				},
			},
			Timestamp: time.Now().Format(time.RFC3339),
		},
	}
	return queue.EnqueueEvent(event)
}
