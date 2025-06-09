# Simple Client Example

This example demonstrates how to build a simple Agent-to-Agent (A2A) client using the `a2ago` library. It shows how to connect to an A2A server, send messages, and query the real-time status of tasks.

## Features

- Basic message sending
- Task status querying
- WebSocket connection for real-time updates
- Error handling and retry logic

## Prerequisites

- Go 1.21 or higher
- A running A2A server, such as the `a2ago/examples/server-hello-world` example, listening on `http://localhost:8089`

## Getting Started

1.  Ensure you have cloned the `a2ago` repository and navigated to its root directory.

2.  Navigate to this example directory:

```bash
cd examples/simple-client
```

3.  Install dependencies:

```bash
go mod download
```

4.  Run the client:

```bash
go run main.go
```

The client will connect to the server, send a test message, and display the task status and response.

## Code Structure

- `main.go`: Client initialization and main logic
- `client.go`: A2A client implementation
- `model.go`: Data models and types

## License

This example is part of the A2AGO project and is licensed under the MIT License. 