package model

// AgentProvider represents information about the provider of an agent
type AgentProvider struct {
	// Organization is the name of the organization providing the agent
	Organization string `json:"organization"`

	// URL is an optional URL pointing to the provider's website or information
	URL string `json:"url,omitempty"`
}

// NewAgentProvider creates a new agent provider
func NewAgentProvider(organization string, url string) *AgentProvider {
	return &AgentProvider{
		Organization: organization,
		URL:          url,
	}
}

// GetOrganization returns the organization name
func (p *AgentProvider) GetOrganization() string {
	return p.Organization
}

// SetOrganization sets the organization name
func (p *AgentProvider) SetOrganization(organization string) {
	p.Organization = organization
}

// GetURL returns the provider's URL
func (p *AgentProvider) GetURL() string {
	return p.URL
}

// SetURL sets the provider's URL
func (p *AgentProvider) SetURL(url string) {
	p.URL = url
}
