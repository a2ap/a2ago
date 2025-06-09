package util

import (
	"github.com/google/uuid"
)

// GenerateUUID generates a new UUID.
func GenerateUUID() string {
	return uuid.New().String()
}
