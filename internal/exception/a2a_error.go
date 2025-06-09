package exception

import (
	"fmt"
)

// A2AError represents an error in the A2A protocol.
// This error contains error codes, additional data, and task-specific information
// to provide comprehensive error context in A2A protocol communications.
type A2AError struct {
	// Message is the error message
	Message string

	// Code is the error code
	Code int

	// Data is additional data about the error
	Data interface{}

	// TaskID is the ID of the task that caused the error
	TaskID string

	// Cause is the underlying error
	Cause error
}

// Common error codes
const (
	// InvalidParams indicates that invalid method parameter(s) were provided
	InvalidParams = -32602

	// MethodNotFound indicates that the method does not exist or is not available
	MethodNotFound = -32601

	// TaskNotFound indicates that the requested task does not exist
	TaskNotFound = 1001

	// TaskCancelled indicates that the task has already been cancelled
	TaskCancelled = 1002

	// AgentExecutionError indicates that there was an error during agent execution
	AgentExecutionError = 1003

	// AuthenticationError indicates that authentication failed
	AuthenticationError = 1004

	// AuthorizationError indicates that authorization failed
	AuthorizationError = 1005
)

// NewA2AError creates a new A2A error with a message
func NewA2AError(message string) *A2AError {
	return &A2AError{
		Message: message,
	}
}

// NewA2AErrorWithCause creates a new A2A error with a message and cause
func NewA2AErrorWithCause(message string, cause error) *A2AError {
	return &A2AError{
		Message: message,
		Cause:   cause,
	}
}

// NewA2AErrorWithAll creates a new A2A error with all properties
func NewA2AErrorWithAll(message string, code int, data interface{}, taskID string) *A2AError {
	return &A2AError{
		Message: message,
		Code:    code,
		Data:    data,
		TaskID:  taskID,
	}
}

// NewA2AErrorWithAllAndCause creates a new A2A error with all properties and cause
func NewA2AErrorWithAllAndCause(message string, cause error, code int, data interface{}, taskID string) *A2AError {
	return &A2AError{
		Message: message,
		Cause:   cause,
		Code:    code,
		Data:    data,
		TaskID:  taskID,
	}
}

// Error returns the error message
func (e *A2AError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Cause)
	}
	return e.Message
}

// Unwrap returns the cause of the error
func (e *A2AError) Unwrap() error {
	return e.Cause
}

// GetCode returns the error code
func (e *A2AError) GetCode() int {
	return e.Code
}

// SetCode sets the error code
func (e *A2AError) SetCode(code int) {
	e.Code = code
}

// GetData returns the additional data
func (e *A2AError) GetData() interface{} {
	return e.Data
}

// SetData sets the additional data
func (e *A2AError) SetData(data interface{}) {
	e.Data = data
}

// GetTaskID returns the task ID
func (e *A2AError) GetTaskID() string {
	return e.TaskID
}

// SetTaskID sets the task ID
func (e *A2AError) SetTaskID(taskID string) {
	e.TaskID = taskID
}

// String returns a string representation of the error
func (e *A2AError) String() string {
	return fmt.Sprintf("A2AError{code=%d, data=%v, taskID='%s', message='%s'}", e.Code, e.Data, e.TaskID, e.Message)
}
