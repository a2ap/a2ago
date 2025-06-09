package model

// SendStreamingMessageResponse is a marker interface for responses emitted by streaming message operations.
// This interface serves as a common type for all possible response objects that can be
// emitted in a streaming context when sending messages to an agent. The streaming
// response can include various types of objects:
//
// - Message objects: Immediate responses or intermediate messages
// - Task objects: Task creation or final task state
// - TaskStatusUpdateEvent objects: Status changes during task execution
// - TaskArtifactUpdateEvent objects: Artifact updates during task processing
//
// This design enables type-safe handling of heterogeneous streaming responses while
// providing flexibility for the A2A protocol's asynchronous communication patterns.
// Clients can use pattern matching or type checking to handle different response
// types appropriately in the streaming context.
type SendStreamingMessageResponse interface {
	// IsSendStreamingMessageResponse is a marker method to ensure type safety
	IsSendStreamingMessageResponse()
}
