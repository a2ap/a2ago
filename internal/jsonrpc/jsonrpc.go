package jsonrpc

// JSONRPCRequest represents a JSON-RPC request
type JSONRPCRequest struct {
	ID     string      `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

// JSONRPCResponse represents a JSON-RPC response
type JSONRPCResponse struct {
	ID     string        `json:"id"`
	Result interface{}   `json:"result,omitempty"`
	Error  *JSONRPCError `json:"error,omitempty"`
}

// JSONRPCError represents a JSON-RPC error
type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// JSON-RPC error codes
const (
	ParseError     = -32700
	InvalidRequest = -32600
	MethodNotFound = -32601
	InvalidParams  = -32602
	InternalError  = -32603
)

// NewJSONRPCRequest creates a new JSON-RPC request
func NewJSONRPCRequest(method string, params interface{}, id string) *JSONRPCRequest {
	return &JSONRPCRequest{
		ID:     id,
		Method: method,
		Params: params,
	}
}
