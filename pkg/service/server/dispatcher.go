package server

import (
	"github.com/a2a4j/a2ago/internal/jsonrpc"
)

// Dispatcher defines the interface for handling JSON-RPC requests
type Dispatcher interface {
	// Dispatch handles synchronous JSON-RPC requests
	Dispatch(request *jsonrpc.JSONRPCRequest) *jsonrpc.JSONRPCResponse

	// DispatchStream handles streaming JSON-RPC requests
	DispatchStream(request *jsonrpc.JSONRPCRequest) (<-chan *jsonrpc.JSONRPCResponse, error)
}

// DefaultDispatcher implements the Dispatcher interface for handling JSON-RPC requests
