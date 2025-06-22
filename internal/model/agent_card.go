package model

// AgentCard represents an agent's capabilities and metadata
type AgentCard struct {
	// ID is the unique identifier of the agent
	ID string `json:"id,omitempty"`

	// Name is the name of the agent
	Name string `json:"name"`

	// Description is an optional description of the agent
	Description string `json:"description,omitempty"`

	// URL is the base URL endpoint for interacting with the agent
	URL string `json:"url"`

	// Provider is information about the provider of the agent
	Provider *AgentProvider `json:"provider,omitempty"`

	// Version is the version identifier for the agent or its API
	Version string `json:"version"`

	// DocumentationURL is an optional URL pointing to the agent's documentation
	DocumentationURL string `json:"documentationUrl,omitempty"`

	// Capabilities are the capabilities supported by the agent
	Capabilities *AgentCapabilities `json:"capabilities,omitempty"`

	// Authentication are authentication details required to interact with the agent
	Authentication *AgentAuthentication `json:"authentication,omitempty"`

	// SecuritySchemes are security scheme details used for authenticating with this agent
	SecuritySchemes map[string]*SecurityScheme `json:"securitySchemes,omitempty"`

	// Security are security requirements for contacting the agent
	Security []map[string][]string `json:"security,omitempty"`

	// DefaultInputModes are default input modes supported by the agent (e.g., 'text', 'file', 'json')
	DefaultInputModes []string `json:"defaultInputModes,omitempty"`

	// DefaultOutputModes are default output modes supported by the agent (e.g., 'text', 'file', 'json')
	DefaultOutputModes []string `json:"defaultOutputModes,omitempty"`

	// Skills are specific skills offered by the agent
	Skills []*AgentSkill `json:"skills,omitempty"`

	// SupportsAuthenticatedExtendedCard is whether the agent supports authenticated extended card
	SupportsAuthenticatedExtendedCard bool `json:"supportsAuthenticatedExtendedCard,omitempty"`
}

// NewAgentCard creates a new agent card
func NewAgentCard(id string, name string, description string, url string, provider *AgentProvider, version string, documentationURL string, capabilities *AgentCapabilities, authentication *AgentAuthentication, securitySchemes map[string]*SecurityScheme, security []map[string][]string, defaultInputModes []string, defaultOutputModes []string, skills []*AgentSkill, supportsAuthenticatedExtendedCard bool) *AgentCard {
	return &AgentCard{
		ID:                                id,
		Name:                              name,
		Description:                       description,
		URL:                               url,
		Provider:                          provider,
		Version:                           version,
		DocumentationURL:                  documentationURL,
		Capabilities:                      capabilities,
		Authentication:                    authentication,
		SecuritySchemes:                   securitySchemes,
		Security:                          security,
		DefaultInputModes:                 defaultInputModes,
		DefaultOutputModes:                defaultOutputModes,
		Skills:                            skills,
		SupportsAuthenticatedExtendedCard: supportsAuthenticatedExtendedCard,
	}
}

// GetID returns the agent ID
func (c *AgentCard) GetID() string {
	return c.ID
}

// SetID sets the agent ID
func (c *AgentCard) SetID(id string) {
	c.ID = id
}

// GetName returns the agent name
func (c *AgentCard) GetName() string {
	return c.Name
}

// SetName sets the agent name
func (c *AgentCard) SetName(name string) {
	c.Name = name
}

// GetDescription returns the agent description
func (c *AgentCard) GetDescription() string {
	return c.Description
}

// SetDescription sets the agent description
func (c *AgentCard) SetDescription(description string) {
	c.Description = description
}

// GetURL returns the agent URL
func (c *AgentCard) GetURL() string {
	return c.URL
}

// SetURL sets the agent URL
func (c *AgentCard) SetURL(url string) {
	c.URL = url
}

// GetProvider returns the agent provider
func (c *AgentCard) GetProvider() *AgentProvider {
	return c.Provider
}

// SetProvider sets the agent provider
func (c *AgentCard) SetProvider(provider *AgentProvider) {
	c.Provider = provider
}

// GetVersion returns the agent version
func (c *AgentCard) GetVersion() string {
	return c.Version
}

// SetVersion sets the agent version
func (c *AgentCard) SetVersion(version string) {
	c.Version = version
}

// GetDocumentationURL returns the agent documentation URL
func (c *AgentCard) GetDocumentationURL() string {
	return c.DocumentationURL
}

// SetDocumentationURL sets the agent documentation URL
func (c *AgentCard) SetDocumentationURL(documentationURL string) {
	c.DocumentationURL = documentationURL
}

// GetCapabilities returns the agent capabilities
func (c *AgentCard) GetCapabilities() *AgentCapabilities {
	return c.Capabilities
}

// SetCapabilities sets the agent capabilities
func (c *AgentCard) SetCapabilities(capabilities *AgentCapabilities) {
	c.Capabilities = capabilities
}

// GetAuthentication returns the agent authentication
func (c *AgentCard) GetAuthentication() *AgentAuthentication {
	return c.Authentication
}

// SetAuthentication sets the agent authentication
func (c *AgentCard) SetAuthentication(authentication *AgentAuthentication) {
	c.Authentication = authentication
}

// GetSecuritySchemes returns the agent security schemes
func (c *AgentCard) GetSecuritySchemes() map[string]*SecurityScheme {
	return c.SecuritySchemes
}

// SetSecuritySchemes sets the agent security schemes
func (c *AgentCard) SetSecuritySchemes(securitySchemes map[string]*SecurityScheme) {
	c.SecuritySchemes = securitySchemes
}

// GetSecurity returns the agent security requirements
func (c *AgentCard) GetSecurity() []map[string][]string {
	return c.Security
}

// SetSecurity sets the agent security requirements
func (c *AgentCard) SetSecurity(security []map[string][]string) {
	c.Security = security
}

// GetDefaultInputModes returns the agent default input modes
func (c *AgentCard) GetDefaultInputModes() []string {
	return c.DefaultInputModes
}

// SetDefaultInputModes sets the agent default input modes
func (c *AgentCard) SetDefaultInputModes(defaultInputModes []string) {
	c.DefaultInputModes = defaultInputModes
}

// GetDefaultOutputModes returns the agent default output modes
func (c *AgentCard) GetDefaultOutputModes() []string {
	return c.DefaultOutputModes
}

// SetDefaultOutputModes sets the agent default output modes
func (c *AgentCard) SetDefaultOutputModes(defaultOutputModes []string) {
	c.DefaultOutputModes = defaultOutputModes
}

// GetSkills returns the agent skills
func (c *AgentCard) GetSkills() []*AgentSkill {
	return c.Skills
}

// SetSkills sets the agent skills
func (c *AgentCard) SetSkills(skills []*AgentSkill) {
	c.Skills = skills
}

// GetSupportsAuthenticatedExtendedCard returns whether the agent supports authenticated extended card
func (c *AgentCard) GetSupportsAuthenticatedExtendedCard() bool {
	return c.SupportsAuthenticatedExtendedCard
}

// SetSupportsAuthenticatedExtendedCard sets whether the agent supports authenticated extended card
func (c *AgentCard) SetSupportsAuthenticatedExtendedCard(supports bool) {
	c.SupportsAuthenticatedExtendedCard = supports
}
