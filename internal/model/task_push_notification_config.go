package model

// TaskPushNotificationConfig represents the configuration for task push notifications
type TaskPushNotificationConfig struct {
	// TaskID is the ID of the task
	TaskID string `json:"taskId"`
	// URL is the URL to send notifications to
	URL string `json:"url"`
}

// NewTaskPushNotificationConfig creates a new TaskPushNotificationConfig
func NewTaskPushNotificationConfig(taskID, url string) *TaskPushNotificationConfig {
	return &TaskPushNotificationConfig{
		TaskID: taskID,
		URL:    url,
	}
}

// TaskPushNotificationSetResult represents the result of setting a task push notification
type TaskPushNotificationSetResult struct {
	Success bool `json:"success"`
}

// TaskPushNotificationGetResult represents the result of getting a task push notification
type TaskPushNotificationGetResult struct {
	Config *TaskPushNotificationConfig `json:"config"`
}

// GetTaskID returns the task ID
func (c *TaskPushNotificationConfig) GetTaskID() string {
	return c.TaskID
}

// SetTaskID sets the task ID
func (c *TaskPushNotificationConfig) SetTaskID(taskID string) {
	c.TaskID = taskID
}

// GetURL returns the URL
func (c *TaskPushNotificationConfig) GetURL() string {
	return c.URL
}
