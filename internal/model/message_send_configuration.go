package model

// MessageSendConfiguration represents configuration for sending messages in the A2A4J framework.
type MessageSendConfiguration struct {
	// AcceptedOutputModes are the accepted output modalities by the client
	AcceptedOutputModes []string `json:"acceptedOutputModes,omitempty"`

	// HistoryLength is the number of recent messages to be retrieved
	HistoryLength *int `json:"historyLength,omitempty"`

	// PushNotificationConfig is where the server should send notifications when disconnected
	PushNotificationConfig *PushNotificationConfig `json:"pushNotificationConfig,omitempty"`

	// Blocking indicates if the server should treat the client as a blocking request
	Blocking *bool `json:"blocking,omitempty"`
}

// NewMessageSendConfiguration creates a new MessageSendConfiguration
func NewMessageSendConfiguration(acceptedOutputModes []string, historyLength *int,
	pushNotificationConfig *PushNotificationConfig, blocking *bool) *MessageSendConfiguration {
	return &MessageSendConfiguration{
		AcceptedOutputModes:    acceptedOutputModes,
		HistoryLength:          historyLength,
		PushNotificationConfig: pushNotificationConfig,
		Blocking:               blocking,
	}
}

// GetAcceptedOutputModes returns the accepted output modes
func (c *MessageSendConfiguration) GetAcceptedOutputModes() []string {
	return c.AcceptedOutputModes
}

// SetAcceptedOutputModes sets the accepted output modes
func (c *MessageSendConfiguration) SetAcceptedOutputModes(acceptedOutputModes []string) {
	c.AcceptedOutputModes = acceptedOutputModes
}

// GetHistoryLength returns the history length
func (c *MessageSendConfiguration) GetHistoryLength() *int {
	return c.HistoryLength
}

// SetHistoryLength sets the history length
func (c *MessageSendConfiguration) SetHistoryLength(historyLength *int) {
	c.HistoryLength = historyLength
}

// GetPushNotificationConfig returns the push notification config
func (c *MessageSendConfiguration) GetPushNotificationConfig() *PushNotificationConfig {
	return c.PushNotificationConfig
}

// SetPushNotificationConfig sets the push notification config
func (c *MessageSendConfiguration) SetPushNotificationConfig(pushNotificationConfig *PushNotificationConfig) {
	c.PushNotificationConfig = pushNotificationConfig
}

// GetBlocking returns the blocking flag
func (c *MessageSendConfiguration) GetBlocking() *bool {
	return c.Blocking
}

// SetBlocking sets the blocking flag
func (c *MessageSendConfiguration) SetBlocking(blocking *bool) {
	c.Blocking = blocking
}
