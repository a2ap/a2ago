package model

// AgentAuthentication represents authentication information for an agent
type AgentAuthentication struct {
	// Schemes are the authentication schemes supported by the agent
	Schemes []string `json:"schemes"`

	// Credentials are optional credentials for authentication
	Credentials string `json:"credentials,omitempty"`
}

// NewAgentAuthentication creates a new agent authentication
func NewAgentAuthentication(schemes []string, credentials string) *AgentAuthentication {
	return &AgentAuthentication{
		Schemes:     schemes,
		Credentials: credentials,
	}
}

// GetSchemes returns the authentication schemes
func (a *AgentAuthentication) GetSchemes() []string {
	return a.Schemes
}

// SetSchemes sets the authentication schemes
func (a *AgentAuthentication) SetSchemes(schemes []string) {
	a.Schemes = schemes
}

// GetCredentials returns the authentication credentials
func (a *AgentAuthentication) GetCredentials() string {
	return a.Credentials
}

// SetCredentials sets the authentication credentials
func (a *AgentAuthentication) SetCredentials(credentials string) {
	a.Credentials = credentials
}
