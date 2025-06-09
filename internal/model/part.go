package model

import (
	"encoding/json"
)

// PartType represents the type of a part
type PartType string

const (
	// PartTypeText represents a text part
	PartTypeText PartType = "text"
	// PartTypeFile represents a file part
	PartTypeFile PartType = "file"
	// PartTypeData represents a data part
	PartTypeData PartType = "data"
)

// Part represents a part of a message
type Part interface {
	// GetType returns the type of the part
	GetType() PartType
	// GetContent returns the content of the part
	GetContent() interface{}
}

// BasePart provides common functionality for all part types
type BasePart struct {
	// Kind is the kind type of the part
	Kind string `json:"kind"`
	// Metadata is optional metadata associated with the part
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Type is the type of the part
	Type PartType `json:"type"`
}

// GetType returns the type of the part
func (p *BasePart) GetType() PartType {
	return p.Type
}

// GetKind returns the kind of the part
func (p *BasePart) GetKind() string {
	return p.Kind
}

// GetMetadata returns the metadata of the part
func (p *BasePart) GetMetadata() map[string]interface{} {
	return p.Metadata
}

// SetMetadata sets a metadata key-value pair
func (p *BasePart) SetMetadata(key string, value interface{}) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]interface{})
	}
	p.Metadata[key] = value
}

// MarshalJSON implements the json.Marshaler interface
func (p *BasePart) MarshalJSON() ([]byte, error) {
	type Alias BasePart
	return json.Marshal(&struct {
		*Alias
		Type PartType `json:"type"`
	}{
		Alias: (*Alias)(p),
		Type:  p.GetType(),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (p *BasePart) UnmarshalJSON(data []byte) error {
	type Alias BasePart
	aux := &struct {
		*Alias
		Type PartType `json:"type"`
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
