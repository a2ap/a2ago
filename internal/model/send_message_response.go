package model

// SendMessageResponse is a marker interface for responses returned by synchronous message sending operations.
// This interface serves as a common type for all possible response objects that can be
// returned when sending a message to an agent in a non-streaming context. The response
// can be either:
//
// - A Message object: When the agent provides an immediate response
// - A Task object: When the agent creates a task to handle the request asynchronously
//
// This design allows for flexible response handling while maintaining type safety
// in the A2A protocol implementation. Clients can check the actual type of the
// response to determine the appropriate handling strategy.
type SendMessageResponse interface {
	// IsSendMessageResponse is a marker method to ensure type safety
	IsSendMessageResponse()
}
