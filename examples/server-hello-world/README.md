# Server Hello World Example

This example demonstrates how to build a simple Agent-to-Agent (A2A) server using the `a2ago` library. It implements core A2A server components that can receive messages from clients, process tasks, and return responses.

## Features

- Basic message handling
- Task processing
- Response generation
- WebSocket support for real-time communication

## Prerequisites

- Go 1.21 or higher
- A basic understanding of Go programming
- Familiarity with the A2A protocol

## Getting Started

1.  Ensure you have cloned the `a2ago` repository and navigated to its root directory.

2.  Navigate to this example directory:

```bash
cd examples/server-hello-world
```

3.  Install dependencies:

```bash
go mod download
```

4.  Run the server:

```bash
go run main.go
```

The server will start on `http://localhost:8089`.

## Testing the Server

You can test the server using the `simple-client` example or by sending HTTP requests directly. The server implements the following endpoints:

- `POST /message`: Send a message to the server
- `GET /task/{taskId}`: Query task status
- `GET /ws`: WebSocket endpoint for real-time communication

## Code Structure

- `main.go`: Server initialization and configuration
- `handler.go`: Request handlers and business logic
- `model.go`: Data models and types

## License

This example is part of the A2AGO project and is licensed under the MIT License. 