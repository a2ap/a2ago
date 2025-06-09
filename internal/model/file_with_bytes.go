package model

import (
	"encoding/base64"
	"encoding/json"
)

// FileWithBytes represents file content in bytes
type FileWithBytes struct {
	BaseFileContent
	// Content is the base64-encoded content
	Content string `json:"content"`
}

// NewFileWithBytes creates a new FileWithBytes with the given content
func NewFileWithBytes(content []byte, name, mimeType string) *FileWithBytes {
	return &FileWithBytes{
		BaseFileContent: BaseFileContent{
			Name:     name,
			MimeType: mimeType,
		},
		Content: base64.StdEncoding.EncodeToString(content),
	}
}

// GetBytes returns the decoded content
func (f *FileWithBytes) GetBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(f.Content)
}

// WithName sets the name of the file
func (f *FileWithBytes) WithName(name string) *FileWithBytes {
	f.Name = name
	return f
}

// WithMimeType sets the mime type of the file
func (f *FileWithBytes) WithMimeType(mimeType string) *FileWithBytes {
	f.MimeType = mimeType
	return f
}

// MarshalJSON implements the json.Marshaler interface
func (f *FileWithBytes) MarshalJSON() ([]byte, error) {
	type Alias FileWithBytes
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (f *FileWithBytes) UnmarshalJSON(data []byte) error {
	type Alias FileWithBytes
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
