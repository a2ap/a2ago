package model

import (
	"encoding/json"
	"time"
)

// Task represents a task in the A2A system
type Task struct {
	// ID is the unique identifier of the task
	ID string `json:"id"`
	// ContextID is the ID of the context this task belongs to
	ContextID string `json:"contextId,omitempty"`
	// Status is the current status of the task
	Status *TaskStatus `json:"status"`
	// Artifacts is the list of artifacts associated with the task
	Artifacts []*TaskArtifact `json:"artifacts"`
	// History is the list of messages exchanged within this task
	History []*Message `json:"history"`
	// Metadata is the metadata associated with the task
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// CreatedAt is the creation time of the task
	CreatedAt string `json:"createdAt,omitempty"`
}

// NewTask creates a new Task
func NewTask(id string) *Task {
	return &Task{
		ID:        id,
		CreatedAt: time.Now().Format(time.RFC3339),
		Artifacts: make([]*TaskArtifact, 0),
		History:   make([]*Message, 0),
		Metadata:  make(map[string]interface{}),
	}
}

// GetID returns the ID of the task
func (t *Task) GetID() string {
	return t.ID
}

// GetContextID returns the context ID of the task
func (t *Task) GetContextID() string {
	return t.ContextID
}

// SetContextID sets the context ID of the task
func (t *Task) SetContextID(contextID string) {
	t.ContextID = contextID
}

// GetStatus returns the status of the task
func (t *Task) GetStatus() *TaskStatus {
	return t.Status
}

// SetStatus sets the status of the task
func (t *Task) SetStatus(status *TaskStatus) {
	t.Status = status
}

// GetArtifacts returns the artifacts of the task
func (t *Task) GetArtifacts() []*TaskArtifact {
	return t.Artifacts
}

// SetArtifacts sets the artifacts of the task
func (t *Task) SetArtifacts(artifacts []*TaskArtifact) {
	t.Artifacts = artifacts
}

// GetHistory returns the history of the task
func (t *Task) GetHistory() []*Message {
	return t.History
}

// SetHistory sets the history of the task
func (t *Task) SetHistory(history []*Message) {
	t.History = history
}

// GetMetadata returns the metadata of the task
func (t *Task) GetMetadata() map[string]interface{} {
	return t.Metadata
}

// SetMetadata sets the metadata of the task
func (t *Task) SetMetadata(metadata map[string]interface{}) {
	if t.Metadata == nil {
		t.Metadata = make(map[string]interface{})
	}
	t.Metadata = metadata
}

// IsSendMessageResponse implements the SendMessageResponse interface
func (t *Task) IsSendMessageResponse() {}

// IsSendStreamingMessageResponse implements the SendStreamingMessageResponse interface
func (t *Task) IsSendStreamingMessageResponse() {}

// AddArtifact adds an artifact to the task
func (t *Task) AddArtifact(artifact *TaskArtifact) {
	if t.Artifacts == nil {
		t.Artifacts = make([]*TaskArtifact, 0)
	}
	t.Artifacts = append(t.Artifacts, artifact)
}

// MarshalJSON implements the json.Marshaler interface
func (t *Task) MarshalJSON() ([]byte, error) {
	type Alias Task
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (t *Task) UnmarshalJSON(data []byte) error {
	type Alias Task
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
