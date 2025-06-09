package model

// AgentCapabilities represents the capabilities of an agent.
type AgentCapabilities struct {
	// Streaming indicates if the agent supports SSE streaming methods
	Streaming bool `json:"streaming,omitempty"`

	// PushNotifications indicates if the agent supports push notification methods
	PushNotifications bool `json:"pushNotifications,omitempty"`

	// StateTransitionHistory indicates if the agent exposes status change history for tasks
	StateTransitionHistory bool `json:"stateTransitionHistory,omitempty"`
}

// NewAgentCapabilities creates a new AgentCapabilities.
func NewAgentCapabilities(streaming, pushNotifications, stateTransitionHistory bool) *AgentCapabilities {
	return &AgentCapabilities{
		Streaming:              streaming,
		PushNotifications:      pushNotifications,
		StateTransitionHistory: stateTransitionHistory,
	}
}

// GetStreaming returns whether streaming is supported
func (c *AgentCapabilities) GetStreaming() bool {
	return c.Streaming
}

// SetStreaming sets whether streaming is supported
func (c *AgentCapabilities) SetStreaming(streaming bool) {
	c.Streaming = streaming
}

// GetPushNotifications returns whether push notifications are supported
func (c *AgentCapabilities) GetPushNotifications() bool {
	return c.PushNotifications
}

// SetPushNotifications sets whether push notifications are supported
func (c *AgentCapabilities) SetPushNotifications(pushNotifications bool) {
	c.PushNotifications = pushNotifications
}

// GetStateTransitionHistory returns whether state transition history is supported
func (c *AgentCapabilities) GetStateTransitionHistory() bool {
	return c.StateTransitionHistory
}

// SetStateTransitionHistory sets whether state transition history is supported
func (c *AgentCapabilities) SetStateTransitionHistory(stateTransitionHistory bool) {
	c.StateTransitionHistory = stateTransitionHistory
}
