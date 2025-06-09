package model

import (
	"encoding/json"
)

// DataPart represents a data part of a message, containing structured data
type DataPart struct {
	BasePart
	// Data is the structured data content
	Data interface{} `json:"data"`
}

// NewDataPart creates a new DataPart with the given data
func NewDataPart(data interface{}) *DataPart {
	return &DataPart{
		BasePart: BasePart{
			Kind: "data",
			Type: PartTypeData,
		},
		Data: data,
	}
}

// GetType returns the type of the part
func (p *DataPart) GetType() PartType {
	return PartTypeData
}

// WithMetadata sets the metadata for the data part
func (p *DataPart) WithMetadata(metadata map[string]interface{}) *DataPart {
	p.BasePart.Metadata = metadata
	return p
}

// SetMetadata sets a metadata key-value pair
func (p *DataPart) SetMetadata(key string, value interface{}) *DataPart {
	p.BasePart.SetMetadata(key, value)
	return p
}

// MarshalJSON implements the json.Marshaler interface
func (p *DataPart) MarshalJSON() ([]byte, error) {
	type Alias DataPart
	return json.Marshal(&struct {
		*Alias
		Type PartType `json:"type"`
	}{
		Alias: (*Alias)(p),
		Type:  p.GetType(),
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (p *DataPart) UnmarshalJSON(data []byte) error {
	type Alias DataPart
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
