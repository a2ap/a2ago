package util

import (
	"encoding/json"
	"strings"
)

// JsonUtil provides utility functions for JSON serialization and deserialization operations.
// This package provides a centralized JSON processing facility optimized for the A2A protocol.
// It handles common JSON operations including object serialization, deserialization with type safety,
// and JSON validation.
//
// Key features:
//   - Provides null-safe operations with proper error handling
//   - Includes JSON string validation utilities
//   - All functions are thread-safe, making this package suitable for concurrent usage
//     across the A2A framework.
type JsonUtil struct{}

// ToJson converts an object to a JSON string
func (j *JsonUtil) ToJson(source interface{}) (string, error) {
	if source == nil {
		return "", nil
	}
	bytes, err := json.Marshal(source)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// FromJson converts a JSON string to an object of the specified type
func (j *JsonUtil) FromJson(jsonStr string, target interface{}) error {
	if jsonStr == "" {
		return nil
	}
	return json.Unmarshal([]byte(jsonStr), target)
}

// IsJsonStr checks if the string is a valid JSON string
func (j *JsonUtil) IsJsonStr(jsonStr string) bool {
	if jsonStr == "" {
		return false
	}
	jsonStr = strings.TrimSpace(jsonStr)
	if !(strings.HasPrefix(jsonStr, "{") && strings.HasSuffix(jsonStr, "}")) &&
		!(strings.HasPrefix(jsonStr, "[") && strings.HasSuffix(jsonStr, "]")) {
		return false
	}
	var v interface{}
	return json.Unmarshal([]byte(jsonStr), &v) == nil
}

// NewJsonUtil creates a new JsonUtil instance
func NewJsonUtil() *JsonUtil {
	return &JsonUtil{}
}
