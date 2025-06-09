package model

// TaskIdParams represents parameters for task ID related operations.
type TaskIdParams struct {
	// ID is the ID of the task.
	ID string `json:"id"`

	// Metadata is the message metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewTaskIdParams creates a new TaskIdParams with the given ID.
func NewTaskIdParams(id string) *TaskIdParams {
	return &TaskIdParams{
		ID: id,
	}
}

// NewTaskIdParamsWithMetadata creates a new TaskIdParams with the given ID and metadata.
func NewTaskIdParamsWithMetadata(id string, metadata map[string]interface{}) *TaskIdParams {
	return &TaskIdParams{
		ID:       id,
		Metadata: metadata,
	}
}

// GetID returns the task ID.
func (p *TaskIdParams) GetID() string {
	return p.ID
}

// SetID sets the task ID.
func (p *TaskIdParams) SetID(id string) {
	p.ID = id
}

// GetMetadata returns the metadata.
func (p *TaskIdParams) GetMetadata() map[string]interface{} {
	return p.Metadata
}

// SetMetadata sets the metadata.
func (p *TaskIdParams) SetMetadata(metadata map[string]interface{}) {
	p.Metadata = metadata
}
