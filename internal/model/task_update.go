package model

// TaskUpdate is a marker interface for task update entities in the A2A protocol.
// This interface serves as a common type for all task update events that can be
// applied to modify the state of a task during its lifecycle. Implementations
// of this interface represent different types of updates that can occur:
// - Status updates (via TaskStatusUpdateEvent)
// - Artifact updates (via TaskArtifactUpdateEvent)
// - Other custom update types as needed
type TaskUpdate interface {
	// IsTaskUpdate is a marker method to ensure type safety
	IsTaskUpdate()
}
