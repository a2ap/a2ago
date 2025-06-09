package model

import (
	"time"

	"github.com/a2a4j/a2ago/internal/util"
)

// Event represents an event in the system.
type Event struct {
	// ID is the unique identifier of the event.
	ID string `json:"id"`
	// Type is the type of the event.
	Type string `json:"type"`
	// Data is the data of the event.
	Data interface{} `json:"data"`
	// Timestamp is the timestamp of the event.
	Timestamp string `json:"timestamp"`
}

// NewEvent creates a new Event.
func NewEvent(eventType string, data interface{}) *Event {
	return &Event{
		ID:        util.GenerateUUID(),
		Type:      eventType,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// GetID returns the ID of the event.
func (e *Event) GetID() string {
	return e.ID
}

// GetType returns the type of the event.
func (e *Event) GetType() string {
	return e.Type
}

// GetData returns the data of the event.
func (e *Event) GetData() interface{} {
	return e.Data
}

// GetTimestamp returns the timestamp of the event.
func (e *Event) GetTimestamp() string {
	return e.Timestamp
}

// TaskStatusUpdateEvent represents a status update event for a task (A2A protocol)
type TaskStatusUpdateEvent struct {
	TaskID    string                 `json:"taskId"`
	ContextID string                 `json:"contextId"`
	Kind      string                 `json:"kind"` // always "status-update"
	Status    *TaskStatus            `json:"status"`
	Final     bool                   `json:"final"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// TaskArtifactUpdateEvent represents an artifact update event for a task (A2A protocol)
type TaskArtifactUpdateEvent struct {
	TaskID    string                 `json:"taskId"`
	ContextID string                 `json:"contextId"`
	Kind      string                 `json:"kind"` // always "artifact-update"
	Artifact  *Artifact              `json:"artifact"`
	Append    bool                   `json:"append,omitempty"`
	LastChunk bool                   `json:"lastChunk,omitempty"`
	Final     bool                   `json:"final"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// IsTaskUpdate implements the TaskUpdate interface
func (e *TaskArtifactUpdateEvent) IsTaskUpdate() {}

// IsTaskUpdate implements the TaskUpdate interface
func (e *TaskStatusUpdateEvent) IsTaskUpdate() {}

// IsSendStreamingMessageResponse implements the SendStreamingMessageResponse interface
func (e *TaskStatusUpdateEvent) IsSendStreamingMessageResponse() {}

// IsSendStreamingMessageResponse implements the SendStreamingMessageResponse interface
func (e *TaskArtifactUpdateEvent) IsSendStreamingMessageResponse() {}
