package model

// PushNotificationConfig represents configuration for push notifications.
type PushNotificationConfig struct {
	// URL is the URL to send push notifications to. Required field.
	URL string `json:"url"`

	// AuthToken is the authentication token for push notifications.
	AuthToken string `json:"auth_token,omitempty"`
}

// NewPushNotificationConfig creates a new PushNotificationConfig
func NewPushNotificationConfig(url string, authToken string) *PushNotificationConfig {
	return &PushNotificationConfig{
		URL:       url,
		AuthToken: authToken,
	}
}

// GetURL returns the URL
func (c *PushNotificationConfig) GetURL() string {
	return c.URL
}

// SetURL sets the URL
func (c *PushNotificationConfig) SetURL(url string) {
	c.URL = url
}

// GetAuthToken returns the authentication token
func (c *PushNotificationConfig) GetAuthToken() string {
	return c.AuthToken
}

// SetAuthToken sets the authentication token
func (c *PushNotificationConfig) SetAuthToken(authToken string) {
	c.AuthToken = authToken
}
