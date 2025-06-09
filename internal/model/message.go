package model

import (
	"encoding/json"
	"log"
)

// MessageResponse represents the response from sending a message
type MessageResponse struct {
	TaskID string `json:"taskId"`
	Status string `json:"status"`
}

// Message represents a message in the A2A system
type Message struct {
	// TaskID is the ID of the task this message belongs to
	TaskID string `json:"taskId,omitempty"`
	// ContextID is the ID of the context this message belongs to
	ContextID string `json:"contextId,omitempty"`
	// Parts is the list of message parts
	Parts []Part `json:"parts"`
	// Metadata is the metadata associated with the message
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Kind is the kind of message
	Kind string `json:"kind,omitempty"`
	// Role is the role of the message (e.g., "agent")
	Role string `json:"role,omitempty"`
}

// MessagePart represents a part of a message
type MessagePart struct {
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

// MessageSendParams represents the parameters for sending a message
type MessageSendParams struct {
	// Message is the message to send
	Message *Message `json:"message"`
	// Metadata is the metadata associated with the message
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// StandardSendMessageResponse is the standard response for message/send
// 对齐 Java 端结构，包含任务ID、上下文ID、状态、产物、历史
// 用于 HandleMessage 返回完整响应
type StandardSendMessageResponse struct {
	TaskID    string      `json:"taskId"`
	ContextID string      `json:"contextId,omitempty"`
	Status    *TaskStatus `json:"status,omitempty"`
	Artifacts []*Artifact `json:"artifacts,omitempty"`
	History   []*Message  `json:"history,omitempty"`
}

// IsSendMessageResponse implements the SendMessageResponse interface
func (r *StandardSendMessageResponse) IsSendMessageResponse() {}

// NewMessage creates a new Message
func NewMessage(taskID, contextID string, parts []Part) *Message {
	return &Message{
		TaskID:    taskID,
		ContextID: contextID,
		Parts:     parts,
		Metadata:  make(map[string]interface{}),
	}
}

// AddPart adds a part to the message
func (m *Message) AddPart(part Part) {
	m.Parts = append(m.Parts, part)
}

// GetParts returns the message parts
func (m *Message) GetParts() []Part {
	return m.Parts
}

// SetParts sets the message parts
func (m *Message) SetParts(parts []Part) {
	m.Parts = parts
}

// GetTaskID returns the task ID
func (m *Message) GetTaskID() string {
	return m.TaskID
}

// SetTaskID sets the task ID
func (m *Message) SetTaskID(taskID string) {
	m.TaskID = taskID
}

// GetContextID returns the context ID
func (m *Message) GetContextID() string {
	return m.ContextID
}

// SetContextID sets the context ID
func (m *Message) SetContextID(contextID string) {
	m.ContextID = contextID
}

// WithTaskID sets the task ID for the message
func (m *Message) WithTaskID(taskID string) *Message {
	m.TaskID = taskID
	return m
}

// WithContextID sets the context ID for the message
func (m *Message) WithContextID(contextID string) *Message {
	m.ContextID = contextID
	return m
}

// WithParts sets the parts for the message
func (m *Message) WithParts(parts []Part) *Message {
	m.Parts = parts
	return m
}

// WithMetadata sets the metadata for the message
func (m *Message) WithMetadata(metadata map[string]interface{}) *Message {
	m.Metadata = metadata
	return m
}

// WithKind sets the kind for the message
func (m *Message) WithKind(kind string) *Message {
	m.Kind = kind
	return m
}

// SetMetadata sets metadata for the message
func (m *Message) SetMetadata(key string, value interface{}) {
	if m.Metadata == nil {
		m.Metadata = make(map[string]interface{})
	}
	m.Metadata[key] = value
}

// GetMetadata gets metadata from the message
func (m *Message) GetMetadata(key string) interface{} {
	if m.Metadata == nil {
		return nil
	}
	return m.Metadata[key]
}

// IsSendStreamingMessageResponse implements the SendStreamingMessageResponse interface
func (m *Message) IsSendStreamingMessageResponse() {}

// MarshalJSON implements the json.Marshaler interface
func (m *Message) MarshalJSON() ([]byte, error) {
	type Alias Message
	// 手动序列化 Parts，确保每个 Part 调用自己的 MarshalJSON
	parts := make([]json.RawMessage, len(m.Parts))
	for i, part := range m.Parts {
		if part == nil {
			parts[i] = []byte("null")
			continue
		}
		log.Printf("[MarshalJSON] part[%d] type: %T, value: %+v", i, part, part)
		b, err := json.Marshal(part)
		if err != nil {
			log.Printf("[MarshalJSON] part[%d] failed to marshal: %v", i, err)
			// 尝试类型断言为 *TextPart
			if tp, ok := part.(*TextPart); ok {
				log.Printf("[MarshalJSON] part[%d] fallback marshal as *TextPart: %+v", i, tp)
				b, err = json.Marshal(tp)
				if err == nil {
					parts[i] = b
					continue
				}
			}
			return nil, err
		}
		parts[i] = b
	}
	return json.Marshal(&struct {
		*Alias
		Parts []json.RawMessage `json:"parts"`
	}{
		Alias: (*Alias)(m),
		Parts: parts,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	aux := &struct {
		*Alias
		Parts []json.RawMessage `json:"parts"`
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse each part based on its type
	m.Parts = make([]Part, 0, len(aux.Parts))
	for i, partData := range aux.Parts {
		var partType struct {
			Type PartType `json:"type"`
		}
		if err := json.Unmarshal(partData, &partType); err != nil {
			log.Printf("[UnmarshalJSON] part[%d] failed to get type: %v, data: %s", i, err, string(partData))
			return err
		}
		//log.Printf("[UnmarshalJSON] part[%d] type: %q, data: %s", i, partType.Type, string(partData))

		var part Part
		switch partType.Type {
		case PartTypeText:
			var textPart TextPart
			if err := json.Unmarshal(partData, &textPart); err != nil {
				log.Printf("[UnmarshalJSON] part[%d] failed to unmarshal TextPart: %v, data: %s", i, err, string(partData))
				return err
			}
			//log.Printf("[UnmarshalJSON] part[%d] TextPart: %+v", i, textPart)
			part = &textPart
		default:
			//log.Printf("[UnmarshalJSON] part[%d] unknown type: %q, data: %s", i, partType.Type, string(partData))
			return json.Unmarshal(partData, &part)
		}

		m.Parts = append(m.Parts, part)
	}

	return nil
}

// NewMessageSendParams creates a new MessageSendParams
func NewMessageSendParams(message *Message, metadata map[string]interface{}) *MessageSendParams {
	if metadata == nil {
		metadata = make(map[string]interface{})
	}
	return &MessageSendParams{
		Message:  message,
		Metadata: metadata,
	}
}

// GetMessage returns the message
func (p *MessageSendParams) GetMessage() *Message {
	return p.Message
}

// SetMessage sets the message
func (p *MessageSendParams) SetMessage(message *Message) {
	p.Message = message
}

// GetMetadata returns the metadata
func (p *MessageSendParams) GetMetadata() map[string]interface{} {
	return p.Metadata
}

// SetMetadata sets the metadata
func (p *MessageSendParams) SetMetadata(metadata map[string]interface{}) {
	p.Metadata = metadata
}
