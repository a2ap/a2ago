package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/a2ap/a2ago/internal/util"

	model2 "github.com/a2ap/a2ago/internal/model"
	"github.com/a2ap/a2ago/pkg/service/client"

	"github.com/a2ap/a2ago/internal/jsonrpc"
)

// DefaultA2aClient is a default implementation of A2aClient.
type DefaultA2aClient struct {
	agentCard    *model2.AgentCard
	cardResolver client.CardResolver
	client       *http.Client
}

// NewDefaultA2aClient creates a new DefaultA2aClient.
func NewDefaultA2aClient(cardResolver client.CardResolver) *DefaultA2aClient {
	return &DefaultA2aClient{
		cardResolver: cardResolver,
		client:       &http.Client{Timeout: 30 * time.Second},
	}
}

// NewDefaultA2aClientWithCard creates a new DefaultA2aClient with the given AgentCard.
func NewDefaultA2aClientWithCard(agentCard *model2.AgentCard, cardResolver client.CardResolver) *DefaultA2aClient {
	return &DefaultA2aClient{
		agentCard:    agentCard,
		cardResolver: cardResolver,
		client:       &http.Client{Timeout: 30 * time.Second},
	}
}

// AgentCard returns the AgentCard info currently in client.
func (c *DefaultA2aClient) AgentCard() *model2.AgentCard {
	if c.agentCard != nil {
		return c.agentCard
	}
	return c.RetrieveAgentCard()
}

// RetrieveAgentCard retrieves the AgentCard for the server this client connects to.
func (c *DefaultA2aClient) RetrieveAgentCard() *model2.AgentCard {
	card, _ := c.cardResolver.ResolveCard(context.Background())
	c.agentCard = card
	return card
}

// unmarshalResult unmarshals a JSON-RPC response result into the target type.
func unmarshalResult(result interface{}, target interface{}) error {
	resultData, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("error marshaling result: %v", err)
	}
	return json.Unmarshal(resultData, target)
}

// SendMessage sends a task request to the server (non-streaming).
func (c *DefaultA2aClient) SendMessage(ctx context.Context, params *model2.MessageSendParams) (*model2.Task, error) {
	url := fmt.Sprintf("%s/a2a/server", c.agentCard.URL)
	jsonRpcRequest := jsonrpc.NewJSONRPCRequest("message/send", params, util.GenerateUUID())

	jsonData, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON-RPC request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// First try to decode as JSON-RPC response
	var jsonRpcResponse jsonrpc.JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsonRpcResponse); err != nil {
		// If that fails, try to decode as a direct number response
		var taskID string
		if err := json.NewDecoder(resp.Body).Decode(&taskID); err != nil {
			return nil, fmt.Errorf("error decoding response: %v", err)
		}
		// Create a task with the received ID
		task := model2.NewTask(taskID)
		return task, nil
	}

	if jsonRpcResponse.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %v", jsonRpcResponse.Error)
	}

	var task model2.Task
	if err := unmarshalResult(jsonRpcResponse.Result, &task); err != nil {
		return nil, fmt.Errorf("error unmarshaling task: %v", err)
	}

	return &task, nil
}

// SendMessageStream sends a task request and subscribes to streaming updates.
func (c *DefaultA2aClient) SendMessageStream(ctx context.Context, params *model2.MessageSendParams) (<-chan model2.SendStreamingMessageResponse, error) {
	url := fmt.Sprintf("%s/message/stream", c.agentCard.URL)
	jsonRpcRequest := jsonrpc.NewJSONRPCRequest("message/stream", params, util.GenerateUUID())

	jsonData, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON-RPC request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}

	responseChan := make(chan model2.SendStreamingMessageResponse)
	go func() {
		defer close(responseChan)
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		for {
			var jsonRpcResponse jsonrpc.JSONRPCResponse
			if err := decoder.Decode(&jsonRpcResponse); err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Printf("Error decoding JSON-RPC response: %v\n", err)
				break
			}

			if jsonRpcResponse.Error != nil {
				fmt.Printf("JSON-RPC error: %v\n", jsonRpcResponse.Error)
				break
			}

			if jsonRpcResponse.Result != nil {
				var response model2.SendStreamingMessageResponse
				if err := unmarshalResult(jsonRpcResponse.Result, &response); err != nil {
					fmt.Printf("Error unmarshaling response: %v\n", err)
					continue
				}
				responseChan <- response
			}
		}
	}()

	return responseChan, nil
}

// GetTask retrieves the current state of a task.
func (c *DefaultA2aClient) GetTask(ctx context.Context, params *model2.TaskQueryParams) (*model2.Task, error) {
	url := fmt.Sprintf("%s/a2a/server", c.agentCard.URL)
	jsonRpcRequest := jsonrpc.NewJSONRPCRequest("tasks/get", params, util.GenerateUUID())

	jsonData, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON-RPC request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var jsonRpcResponse jsonrpc.JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsonRpcResponse); err != nil {
		return nil, fmt.Errorf("error decoding JSON-RPC response: %v", err)
	}

	if jsonRpcResponse.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %v", jsonRpcResponse.Error)
	}

	var task model2.Task
	if err := unmarshalResult(jsonRpcResponse.Result, &task); err != nil {
		return nil, fmt.Errorf("error unmarshaling task: %v", err)
	}

	return &task, nil
}

// CancelTask cancels a currently running task.
func (c *DefaultA2aClient) CancelTask(ctx context.Context, params *model2.TaskIdParams) (*model2.Task, error) {
	url := fmt.Sprintf("%s/tasks/cancel", c.agentCard.URL)
	jsonRpcRequest := jsonrpc.NewJSONRPCRequest("tasks/cancel", params, util.GenerateUUID())

	jsonData, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON-RPC request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var jsonRpcResponse jsonrpc.JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsonRpcResponse); err != nil {
		return nil, fmt.Errorf("error decoding JSON-RPC response: %v", err)
	}

	if jsonRpcResponse.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %v", jsonRpcResponse.Error)
	}

	var task model2.Task
	if err := unmarshalResult(jsonRpcResponse.Result, &task); err != nil {
		return nil, fmt.Errorf("error unmarshaling task: %v", err)
	}

	return &task, nil
}

// SetTaskPushNotification sets or updates the push notification config for a task.
func (c *DefaultA2aClient) SetTaskPushNotification(ctx context.Context, params *model2.TaskPushNotificationConfig) (*model2.TaskPushNotificationConfig, error) {
	url := fmt.Sprintf("%s/tasks/pushNotificationConfig/set", c.agentCard.URL)
	jsonRpcRequest := jsonrpc.NewJSONRPCRequest("tasks/pushNotificationConfig/set", params, util.GenerateUUID())

	jsonData, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON-RPC request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var jsonRpcResponse jsonrpc.JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsonRpcResponse); err != nil {
		return nil, fmt.Errorf("error decoding JSON-RPC response: %v", err)
	}

	if jsonRpcResponse.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %v", jsonRpcResponse.Error)
	}

	var config model2.TaskPushNotificationConfig
	if err := unmarshalResult(jsonRpcResponse.Result, &config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	return &config, nil
}

// GetTaskPushNotification retrieves the currently configured push notification config for a task.
func (c *DefaultA2aClient) GetTaskPushNotification(ctx context.Context, params *model2.TaskIdParams) (*model2.TaskPushNotificationConfig, error) {
	url := fmt.Sprintf("%s/tasks/pushNotificationConfig/get", c.agentCard.URL)
	jsonRpcRequest := jsonrpc.NewJSONRPCRequest("tasks/pushNotificationConfig/get", params, util.GenerateUUID())

	jsonData, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON-RPC request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	var jsonRpcResponse jsonrpc.JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&jsonRpcResponse); err != nil {
		return nil, fmt.Errorf("error decoding JSON-RPC response: %v", err)
	}

	if jsonRpcResponse.Error != nil {
		return nil, fmt.Errorf("JSON-RPC error: %v", jsonRpcResponse.Error)
	}

	var config model2.TaskPushNotificationConfig
	if err := unmarshalResult(jsonRpcResponse.Result, &config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %v", err)
	}

	return &config, nil
}

// ResubscribeTask resubscribes to updates for a task after a potential connection interruption.
func (c *DefaultA2aClient) ResubscribeTask(ctx context.Context, params *model2.TaskQueryParams) (<-chan model2.SendStreamingMessageResponse, error) {
	url := fmt.Sprintf("%s/tasks/resubscribe", c.agentCard.URL)
	jsonRpcRequest := jsonrpc.NewJSONRPCRequest("tasks/resubscribe", params, util.GenerateUUID())

	jsonData, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON-RPC request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}

	responseChan := make(chan model2.SendStreamingMessageResponse)
	go func() {
		defer close(responseChan)
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		for {
			var jsonRpcResponse jsonrpc.JSONRPCResponse
			if err := decoder.Decode(&jsonRpcResponse); err != nil {
				if err.Error() == "EOF" {
					break
				}
				fmt.Printf("Error decoding JSON-RPC response: %v\n", err)
				break
			}

			if jsonRpcResponse.Error != nil {
				fmt.Printf("JSON-RPC error: %v\n", jsonRpcResponse.Error)
				break
			}

			if jsonRpcResponse.Result != nil {
				var response model2.SendStreamingMessageResponse
				if err := unmarshalResult(jsonRpcResponse.Result, &response); err != nil {
					fmt.Printf("Error unmarshaling response: %v\n", err)
					continue
				}
				responseChan <- response
			}
		}
	}()

	return responseChan, nil
}

// RetrieveAuthenticatedExtendedAgentCard retrieves the authenticated extended AgentCard.
func (c *DefaultA2aClient) RetrieveAuthenticatedExtendedAgentCard(ctx context.Context, authToken string) (*model2.AgentCard, error) {
	if c.agentCard == nil {
		c.RetrieveAgentCard()
	}
	if c.agentCard == nil {
		return nil, fmt.Errorf("agent card not available")
	}

	url := fmt.Sprintf("%s/a2a/agent/authenticatedExtendedCard", c.agentCard.URL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request: %v", err)
	}

	req.Header.Set("X-API-Key", authToken)
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned non-200 status: %s", resp.Status)
	}

	var card model2.AgentCard
	if err := json.NewDecoder(resp.Body).Decode(&card); err != nil {
		return nil, fmt.Errorf("error decoding agent card: %v", err)
	}

	return &card, nil
}

// Supports checks if the server likely supports optional methods based on agent card.
func (c *DefaultA2aClient) Supports(capability string) bool {
	if c.agentCard == nil {
		c.agentCard = c.RetrieveAgentCard()
	}

	if c.agentCard == nil || c.agentCard.Capabilities == nil {
		return false
	}

	// Check agent supports
	switch strings.ToLower(capability) {
	case "streaming":
		return c.agentCard.Capabilities.Streaming
	case "pushnotifications":
		return c.agentCard.Capabilities.PushNotifications
	default:
		return false
	}
}
