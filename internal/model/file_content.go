package model

import (
	"encoding/json"
)

// FileContent represents the content of a file
type FileContent struct {
	// ID is the unique identifier of the file
	ID string `json:"id"`
	// Name is the name of the file
	Name string `json:"name"`
	// MimeType is the MIME type of the file
	MimeType string `json:"mimeType"`
	// Size is the size of the file in bytes
	Size int64 `json:"size"`
	// URI is the URI of the file
	URI string `json:"uri,omitempty"`
	// Bytes is the raw bytes of the file
	Bytes []byte `json:"bytes,omitempty"`
}

// NewFileContent creates a new FileContent
func NewFileContent(id, name, mimeType string, size int64, uri string, bytes []byte) *FileContent {
	return &FileContent{
		ID:       id,
		Name:     name,
		MimeType: mimeType,
		Size:     size,
		URI:      uri,
		Bytes:    bytes,
	}
}

// GetID returns the ID of the file
func (f *FileContent) GetID() string {
	return f.ID
}

// GetName returns the name of the file
func (f *FileContent) GetName() string {
	return f.Name
}

// GetMimeType returns the MIME type of the file
func (f *FileContent) GetMimeType() string {
	return f.MimeType
}

// GetSize returns the size of the file
func (f *FileContent) GetSize() int64 {
	return f.Size
}

// GetURI returns the URI of the file
func (f *FileContent) GetURI() string {
	return f.URI
}

// GetBytes returns the raw bytes of the file
func (f *FileContent) GetBytes() []byte {
	return f.Bytes
}

// BaseFileContent provides common functionality for all file content types
type BaseFileContent struct {
	// Name is the name of the file
	Name string `json:"name"`
	// MimeType is the MIME type of the file
	MimeType string `json:"mimeType"`
}

// GetName returns the name of the file
func (f *BaseFileContent) GetName() string {
	return f.Name
}

// GetMimeType returns the MIME type of the file
func (f *BaseFileContent) GetMimeType() string {
	return f.MimeType
}

// MarshalJSON implements the json.Marshaler interface
func (f *BaseFileContent) MarshalJSON() ([]byte, error) {
	type Alias BaseFileContent
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (f *BaseFileContent) UnmarshalJSON(data []byte) error {
	type Alias BaseFileContent
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
