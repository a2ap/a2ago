package model

// SecurityScheme represents a security scheme for agent authentication
type SecurityScheme struct {
	// Type is the type of security scheme (e.g., "http", "apiKey", "oauth2")
	Type string `json:"type,omitempty"`

	// Scheme is the scheme name for HTTP authentication (e.g., "bearer", "basic")
	Scheme string `json:"scheme,omitempty"`

	// Name is the name of the header, query parameter or cookie for API key authentication
	Name string `json:"name,omitempty"`

	// In is the location of the API key (e.g., "header", "query", "cookie")
	In string `json:"in,omitempty"`

	// Description is a description of the security scheme
	Description string `json:"description,omitempty"`

	// BearerFormat is the bearer format for bearer token authentication
	BearerFormat string `json:"bearerFormat,omitempty"`
}

// NewSecurityScheme creates a new security scheme
func NewSecurityScheme(typ string, scheme string, name string, in string, description string, bearerFormat string) *SecurityScheme {
	return &SecurityScheme{
		Type:         typ,
		Scheme:       scheme,
		Name:         name,
		In:           in,
		Description:  description,
		BearerFormat: bearerFormat,
	}
}

// GetType returns the type of security scheme
func (s *SecurityScheme) GetType() string {
	return s.Type
}

// SetType sets the type of security scheme
func (s *SecurityScheme) SetType(typ string) {
	s.Type = typ
}

// GetScheme returns the scheme name
func (s *SecurityScheme) GetScheme() string {
	return s.Scheme
}

// SetScheme sets the scheme name
func (s *SecurityScheme) SetScheme(scheme string) {
	s.Scheme = scheme
}

// GetName returns the name of the header, query parameter or cookie
func (s *SecurityScheme) GetName() string {
	return s.Name
}

// SetName sets the name of the header, query parameter or cookie
func (s *SecurityScheme) SetName(name string) {
	s.Name = name
}

// GetIn returns the location of the API key
func (s *SecurityScheme) GetIn() string {
	return s.In
}

// SetIn sets the location of the API key
func (s *SecurityScheme) SetIn(in string) {
	s.In = in
}

// GetDescription returns the description of the security scheme
func (s *SecurityScheme) GetDescription() string {
	return s.Description
}

// SetDescription sets the description of the security scheme
func (s *SecurityScheme) SetDescription(description string) {
	s.Description = description
}

// GetBearerFormat returns the bearer format
func (s *SecurityScheme) GetBearerFormat() string {
	return s.BearerFormat
}

// SetBearerFormat sets the bearer format
func (s *SecurityScheme) SetBearerFormat(bearerFormat string) {
	s.BearerFormat = bearerFormat
}
