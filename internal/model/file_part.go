package model

import (
	"encoding/json"
)

// FilePart represents a file part of a message
type FilePart struct {
	BasePart
	// File is the file content
	File *FileContent `json:"file"`
}

// NewFilePart creates a new FilePart with the given file content
func NewFilePart(file *FileContent) *FilePart {
	return &FilePart{
		BasePart: BasePart{
			Kind: "file",
			Type: PartTypeFile,
		},
		File: file,
	}
}

// GetType returns the type of the part
func (p *FilePart) GetType() PartType {
	return PartTypeFile
}

// WithMetadata sets the metadata for the file part
func (p *FilePart) WithMetadata(metadata map[string]interface{}) *FilePart {
	p.BasePart.Metadata = metadata
	return p
}

// SetMetadata sets a metadata key-value pair
func (p *FilePart) SetMetadata(key string, value interface{}) *FilePart {
	p.BasePart.SetMetadata(key, value)
	return p
}

// MarshalJSON implements the json.Marshaler interface
func (p *FilePart) MarshalJSON() ([]byte, error) {
	type Alias FilePart
	return json.Marshal(&struct {
		*Alias
		Type PartType `json:"type"`
	}{
		Alias: (*Alias)(p),
		Type:  p.GetType(),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (p *FilePart) UnmarshalJSON(data []byte) error {
	type Alias FilePart
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
