package model

import (
	"encoding/json"
)

// FileWithUri represents a file with a URI reference
type FileWithUri struct {
	BaseFileContent
	// URI is the URI of the file
	URI string `json:"uri"`
}

// NewFileWithUri creates a new FileWithUri with the given URI
func NewFileWithUri(uri string) *FileWithUri {
	return &FileWithUri{
		URI: uri,
	}
}

// NewFileWithUriWithMetadata creates a new FileWithUri with the given name, MIME type, and URI
func NewFileWithUriWithMetadata(name, mimeType, uri string) *FileWithUri {
	return &FileWithUri{
		BaseFileContent: BaseFileContent{
			Name:     name,
			MimeType: mimeType,
		},
		URI: uri,
	}
}

// WithName sets the name for the file
func (f *FileWithUri) WithName(name string) *FileWithUri {
	f.Name = name
	return f
}

// WithMimeType sets the MIME type for the file
func (f *FileWithUri) WithMimeType(mimeType string) *FileWithUri {
	f.MimeType = mimeType
	return f
}

// MarshalJSON implements the json.Marshaler interface
func (f *FileWithUri) MarshalJSON() ([]byte, error) {
	type Alias FileWithUri
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (f *FileWithUri) UnmarshalJSON(data []byte) error {
	type Alias FileWithUri
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
