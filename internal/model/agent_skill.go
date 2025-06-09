package model

// AgentSkill represents a skill that an agent possesses
type AgentSkill struct {
	// ID is the unique identifier of the skill
	ID string `json:"id"`

	// Name is the name of the skill
	Name string `json:"name"`

	// Description is an optional description of the skill
	Description string `json:"description,omitempty"`

	// Tags are optional tags associated with the skill
	Tags []string `json:"tags,omitempty"`

	// Examples are optional examples of how to use the skill
	Examples []string `json:"examples,omitempty"`

	// InputModes are optional input modes supported by the skill
	InputModes []string `json:"inputModes,omitempty"`

	// OutputModes are optional output modes supported by the skill
	OutputModes []string `json:"outputModes,omitempty"`
}

// NewAgentSkill creates a new agent skill
func NewAgentSkill(id string, name string, description string, tags []string, examples []string, inputModes []string, outputModes []string) *AgentSkill {
	return &AgentSkill{
		ID:          id,
		Name:        name,
		Description: description,
		Tags:        tags,
		Examples:    examples,
		InputModes:  inputModes,
		OutputModes: outputModes,
	}
}

// GetID returns the skill ID
func (s *AgentSkill) GetID() string {
	return s.ID
}

// SetID sets the skill ID
func (s *AgentSkill) SetID(id string) {
	s.ID = id
}

// GetName returns the skill name
func (s *AgentSkill) GetName() string {
	return s.Name
}

// SetName sets the skill name
func (s *AgentSkill) SetName(name string) {
	s.Name = name
}

// GetDescription returns the skill description
func (s *AgentSkill) GetDescription() string {
	return s.Description
}

// SetDescription sets the skill description
func (s *AgentSkill) SetDescription(description string) {
	s.Description = description
}

// GetTags returns the skill tags
func (s *AgentSkill) GetTags() []string {
	return s.Tags
}

// SetTags sets the skill tags
func (s *AgentSkill) SetTags(tags []string) {
	s.Tags = tags
}

// GetExamples returns the skill examples
func (s *AgentSkill) GetExamples() []string {
	return s.Examples
}

// SetExamples sets the skill examples
func (s *AgentSkill) SetExamples(examples []string) {
	s.Examples = examples
}

// GetInputModes returns the skill input modes
func (s *AgentSkill) GetInputModes() []string {
	return s.InputModes
}

// SetInputModes sets the skill input modes
func (s *AgentSkill) SetInputModes(inputModes []string) {
	s.InputModes = inputModes
}

// GetOutputModes returns the skill output modes
func (s *AgentSkill) GetOutputModes() []string {
	return s.OutputModes
}

// SetOutputModes sets the skill output modes
func (s *AgentSkill) SetOutputModes(outputModes []string) {
	s.OutputModes = outputModes
}
