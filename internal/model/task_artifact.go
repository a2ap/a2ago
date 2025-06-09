package model

// TaskArtifact represents an artifact associated with a task
type TaskArtifact struct {
	// ID is the unique identifier of the artifact
	ID string `json:"id"`
	// Content is the content of the artifact
	Content Part `json:"content"`
	// Metadata is the metadata associated with the artifact
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewTaskArtifact creates a new TaskArtifact
func NewTaskArtifact(id string, content Part, metadata map[string]interface{}) *TaskArtifact {
	return &TaskArtifact{
		ID:       id,
		Content:  content,
		Metadata: metadata,
	}
}

// GetID returns the ID of the artifact
func (a *TaskArtifact) GetID() string {
	return a.ID
}

// GetContent returns the content of the artifact
func (a *TaskArtifact) GetContent() Part {
	return a.Content
}

// SetContent sets the content of the artifact
func (a *TaskArtifact) SetContent(content Part) {
	a.Content = content
}

// GetMetadata returns the metadata of the artifact
func (a *TaskArtifact) GetMetadata() map[string]interface{} {
	return a.Metadata
}

// SetMetadata sets the metadata of the artifact
func (a *TaskArtifact) SetMetadata(metadata map[string]interface{}) {
	a.Metadata = metadata
}
