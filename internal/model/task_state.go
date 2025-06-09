package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// TaskState represents the state of a task in the A2A protocol lifecycle.
// State transitions typically follow this pattern:
// SUBMITTED → WORKING → (INPUT_REQUIRED ↔ WORKING) → COMPLETED/FAILED/CANCELED
type TaskState string

const (
	// TaskStateSubmitted indicates that the task has been submitted but not yet started processing
	TaskStateSubmitted TaskState = "submitted"

	// TaskStateWorking indicates that the task is currently being processed by the agent
	TaskStateWorking TaskState = "working"

	// TaskStateInputRequired indicates that the task execution is paused, waiting for additional input from the client
	TaskStateInputRequired TaskState = "input-required"

	// TaskStateCompleted indicates that the task has been successfully completed
	TaskStateCompleted TaskState = "completed"

	// TaskStateFailed indicates that the task execution failed due to an error
	TaskStateFailed TaskState = "failed"

	// TaskStateCanceled indicates that the task was canceled by client request or system intervention
	TaskStateCanceled TaskState = "canceled"

	// TaskStateRejected indicates that the task was rejected by the agent (e.g., invalid parameters, unsupported operation)
	TaskStateRejected TaskState = "rejected"

	// TaskStateAuthRequired indicates that the task requires authentication before it can proceed
	TaskStateAuthRequired TaskState = "auth-required"

	// TaskStateUnknown indicates that the task state is unknown or could not be determined
	TaskStateUnknown TaskState = "unknown"
)

// String returns the string representation of the task state
func (s TaskState) String() string {
	return string(s)
}

// MarshalJSON implements custom JSON marshaling
func (s TaskState) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

// UnmarshalJSON implements custom JSON unmarshaling
func (s *TaskState) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	switch TaskState(str) {
	case TaskStateSubmitted,
		TaskStateWorking,
		TaskStateInputRequired,
		TaskStateCompleted,
		TaskStateFailed,
		TaskStateCanceled,
		TaskStateRejected,
		TaskStateAuthRequired,
		TaskStateUnknown:
		*s = TaskState(str)
		return nil
	default:
		return fmt.Errorf("unknown TaskState value: %s", str)
	}
}

// TaskStatus represents the status of a task
type TaskStatus struct {
	State     TaskState              `json:"state"`
	Message   *Message               `json:"message,omitempty"`
	Timestamp string                 `json:"timestamp"`
	Error     string                 `json:"error,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// NewTaskStatus creates a new TaskStatus with the given state
func NewTaskStatus(state TaskState) *TaskStatus {
	return &TaskStatus{
		State:     state,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// IsTaskUpdate implements the TaskUpdate interface
func (s *TaskStatus) IsTaskUpdate() {}
