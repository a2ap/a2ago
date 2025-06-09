package model

import (
	"encoding/json"
)

// TextPart represents a text part of a message
type TextPart struct {
	BasePart
	// Text is the text content
	Text string `json:"text"`
}

// NewTextPart creates a new TextPart with the given text
func NewTextPart(text string) *TextPart {
	return &TextPart{
		BasePart: BasePart{
			Kind: "text",
			Type: PartTypeText,
		},
		Text: text,
	}
}

// GetType returns the type of the part
func (p *TextPart) GetType() PartType {
	return PartTypeText
}

// GetContent returns the content of the part
func (p *TextPart) GetContent() interface{} {
	return p.Text
}

// WithMetadata sets the metadata for the text part
func (p *TextPart) WithMetadata(metadata map[string]interface{}) *TextPart {
	p.BasePart.Metadata = metadata
	return p
}

// SetMetadata sets a metadata key-value pair
func (p *TextPart) SetMetadata(key string, value interface{}) {
	p.BasePart.SetMetadata(key, value)
}

// MarshalJSON implements the json.Marshaler interface
func (p *TextPart) MarshalJSON() ([]byte, error) {
	// 显式序列化所有字段，确保 Text 字段和 BasePart 字段都能被正确输出
	aux := struct {
		Kind     string                 `json:"kind"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		Type     PartType               `json:"type"`
		Text     string                 `json:"text"`
	}{
		Kind:     p.Kind,
		Metadata: p.Metadata,
		Type:     p.Type,
		Text:     p.Text,
	}
	return json.Marshal(aux)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (p *TextPart) UnmarshalJSON(data []byte) error {
	var aux struct {
		Kind     string                 `json:"kind"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
		Type     PartType               `json:"type"`
		Text     string                 `json:"text"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	p.BasePart.Kind = aux.Kind
	p.BasePart.Metadata = aux.Metadata
	p.BasePart.Type = aux.Type
	p.Text = aux.Text
	return nil
}
