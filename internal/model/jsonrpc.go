package model

// JSONRPCResponse represents a JSON-RPC response
type JSONRPCResponse struct {
	ID     string        `json:"id"`
	Result interface{}   `json:"result,omitempty"`
	Error  *JSONRPCError `json:"error,omitempty"`
}

// JSONRPCError represents a JSON-RPC error
type JSONRPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
