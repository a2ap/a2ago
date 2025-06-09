package model

// TaskQueryParams represents parameters for querying tasks.
type TaskQueryParams struct {
	// TaskID is the ID of the task.
	TaskID string `json:"taskId,omitempty"`

	// SessionID is the session ID associated with the task.
	SessionID string `json:"sessionId,omitempty"`
}

// NewTaskQueryParams creates a new TaskQueryParams with the given task ID.
func NewTaskQueryParams(taskID string) *TaskQueryParams {
	return &TaskQueryParams{
		TaskID: taskID,
	}
}

// NewTaskQueryParamsWithSession creates a new TaskQueryParams with the given task ID and session ID.
func NewTaskQueryParamsWithSession(taskID, sessionID string) *TaskQueryParams {
	return &TaskQueryParams{
		TaskID:    taskID,
		SessionID: sessionID,
	}
}

// GetTaskID returns the task ID.
func (p *TaskQueryParams) GetTaskID() string {
	return p.TaskID
}

// SetTaskID sets the task ID.
func (p *TaskQueryParams) SetTaskID(taskID string) {
	p.TaskID = taskID
}

// GetSessionID returns the session ID.
func (p *TaskQueryParams) GetSessionID() string {
	return p.SessionID
}

// SetSessionID sets the session ID.
func (p *TaskQueryParams) SetSessionID(sessionID string) {
	p.SessionID = sessionID
}
