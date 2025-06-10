package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/a2ap/a2ago/pkg/service/server"

	"github.com/a2ap/a2ago/internal/jsonrpc"
	"github.com/a2ap/a2ago/internal/model"
)

// DefaultDispatcher implements the Dispatcher interface for handling JSON-RPC requests
type DefaultDispatcher struct {
	a2aServer server.A2AServer
}

// NewDefaultDispatcher creates a new instance of DefaultDispatcher
func NewDefaultDispatcher(a2aServer server.A2AServer) *DefaultDispatcher {
	return &DefaultDispatcher{
		a2aServer: a2aServer,
	}
}

// Dispatch handles synchronous JSON-RPC requests
func (d *DefaultDispatcher) Dispatch(request *jsonrpc.JSONRPCRequest) *jsonrpc.JSONRPCResponse {
	response := &jsonrpc.JSONRPCResponse{
		ID: request.ID,
	}

	ctx := context.Background()
	switch request.Method {
	case "message/send":
		var params model.MessageSendParams
		paramsBytes, err := json.Marshal(request.Params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		if err = json.Unmarshal(paramsBytes, &params); err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		messageResponse, err := d.a2aServer.HandleMessage(ctx, &params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InternalError,
				Message: "Internal error",
				Data:    err.Error(),
			}
			return response
		}
		response.Result = messageResponse

	case "tasks/get":
		var params model.TaskIdParams
		paramsBytes, err := json.Marshal(request.Params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		if err = json.Unmarshal(paramsBytes, &params); err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		task, err := d.a2aServer.GetTask(ctx, params.ID)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InternalError,
				Message: "Internal error",
				Data:    err.Error(),
			}
			return response
		}
		response.Result = task

	case "tasks/cancel":
		var params model.TaskIdParams
		paramsBytes, err := json.Marshal(request.Params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		if err = json.Unmarshal(paramsBytes, &params); err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		task, err := d.a2aServer.CancelTask(ctx, params.ID)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InternalError,
				Message: "Internal error",
				Data:    err.Error(),
			}
			return response
		}
		response.Result = task

	case "tasks/pushNotificationConfig/set":
		var config model.TaskPushNotificationConfig
		paramsBytes, err := json.Marshal(request.Params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		if err = json.Unmarshal(paramsBytes, &config); err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		setResult, err := d.a2aServer.SetTaskPushNotification(ctx, config.GetTaskID(), &config)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InternalError,
				Message: "Internal error",
				Data:    err.Error(),
			}
			return response
		}
		response.Result = setResult

	case "tasks/pushNotificationConfig/get":
		var params model.TaskIdParams
		paramsBytes, err := json.Marshal(request.Params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		if err = json.Unmarshal(paramsBytes, &params); err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			return response
		}
		getConfigResult, err := d.a2aServer.GetTaskPushNotification(ctx, params.ID)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InternalError,
				Message: "Internal error",
				Data:    err.Error(),
			}
			return response
		}
		response.Result = getConfigResult

	default:
		log.Printf("Unsupported method: %s", request.Method)
		response.Error = &jsonrpc.JSONRPCError{
			Code:    jsonrpc.MethodNotFound,
			Message: "Method not found",
			Data:    fmt.Sprintf("Method '%s' not supported", request.Method),
		}
	}

	return response
}

// DispatchStream handles streaming JSON-RPC requests
func (d *DefaultDispatcher) DispatchStream(request *jsonrpc.JSONRPCRequest) (<-chan *jsonrpc.JSONRPCResponse, error) {
	response := &jsonrpc.JSONRPCResponse{
		ID: request.ID,
	}

	ctx := context.Background()
	switch request.Method {
	case "message/stream":
		var params model.MessageSendParams
		paramsBytes, err := json.Marshal(request.Params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			ch := make(chan *jsonrpc.JSONRPCResponse, 1)
			ch <- response
			close(ch)
			return ch, nil
		}
		if err = json.Unmarshal(paramsBytes, &params); err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			ch := make(chan *jsonrpc.JSONRPCResponse, 1)
			ch <- response
			close(ch)
			return ch, nil
		}
		modelResponses, err := d.a2aServer.HandleMessageStream(ctx, &params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InternalError,
				Message: "Internal error",
				Data:    err.Error(),
			}
			ch := make(chan *jsonrpc.JSONRPCResponse, 1)
			ch <- response
			close(ch)
			return ch, nil
		}
		responses := make(chan *jsonrpc.JSONRPCResponse)
		go func() {
			for resp := range modelResponses {
				responses <- &jsonrpc.JSONRPCResponse{
					ID:     request.ID,
					Result: resp,
				}
			}
			close(responses)
		}()
		return responses, nil

	case "tasks/resubscribe":
		var params model.TaskIdParams
		paramsBytes, err := json.Marshal(request.Params)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			ch := make(chan *jsonrpc.JSONRPCResponse, 1)
			ch <- response
			close(ch)
			return ch, nil
		}
		if err = json.Unmarshal(paramsBytes, &params); err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InvalidParams,
				Message: "Invalid params",
				Data:    err.Error(),
			}
			ch := make(chan *jsonrpc.JSONRPCResponse, 1)
			ch <- response
			close(ch)
			return ch, nil
		}
		modelResponses, err := d.a2aServer.SubscribeToTaskUpdates(ctx, params.ID)
		if err != nil {
			response.Error = &jsonrpc.JSONRPCError{
				Code:    jsonrpc.InternalError,
				Message: "Internal error",
				Data:    err.Error(),
			}
			ch := make(chan *jsonrpc.JSONRPCResponse, 1)
			ch <- response
			close(ch)
			return ch, nil
		}
		responses := make(chan *jsonrpc.JSONRPCResponse)
		go func() {
			for resp := range modelResponses {
				responses <- &jsonrpc.JSONRPCResponse{
					ID:     request.ID,
					Result: resp,
				}
			}
			close(responses)
		}()
		return responses, nil

	default:
		log.Printf("Unsupported method: %s", request.Method)
		response.Error = &jsonrpc.JSONRPCError{
			Code:    jsonrpc.MethodNotFound,
			Message: "Method not found",
			Data:    fmt.Sprintf("Method '%s' not supported", request.Method),
		}
		ch := make(chan *jsonrpc.JSONRPCResponse, 1)
		ch <- response
		close(ch)
		return ch, nil
	}
}
