package model

// RequestContext represents a request context in the system
type RequestContext struct {
	// TaskID is the ID of the task
	TaskID string `json:"taskId"`

	// ContextID is the ID of the context
	ContextID string `json:"contextId"`

	// Task is the task associated with this context
	Task *Task `json:"task"`

	// RelatedTasks are tasks related to this context
	RelatedTasks []*Task `json:"relatedTasks,omitempty"`
}

// NewRequestContext creates a new RequestContext
func NewRequestContext(taskID, contextID string, task *Task) *RequestContext {
	return &RequestContext{
		TaskID:       taskID,
		ContextID:    contextID,
		Task:         task,
		RelatedTasks: make([]*Task, 0),
	}
}
