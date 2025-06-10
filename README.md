# A2AGO

`a2ago` is a Go language implementation of the [A2A4J (Agent-to-Agent for Java)](https://github.com/a2a4j/a2a4j) framework. It provides a Go client and server library for building and integrating Agent-to-Agent interactions. This library aims to facilitate communication between different agents based on standardized messages and task flows.

## Features

- **Standardized Message Protocol**: Implements the A2A message protocol for consistent communication between agents.
- **Task Flow Management**: Supports task creation, status tracking, and result retrieval.
- **WebSocket Support**: Real-time communication capabilities through WebSocket connections.
- **Extensible Architecture**: Easy to extend and customize for different use cases.
- **Comprehensive Testing**: Includes unit tests and integration tests to ensure reliability.

## Installation

1.  Clone the repository and navigate to the `a2ago` directory:

```bash
git clone https://github.com/a2ap/a2ago.git
cd a2ago
```

2.  Install dependencies:

```bash
go mod download
```

## Usage

### Basic Example

Here's a simple example of how to use the library:

```go
package main

import (
    "github.com/a2ap/a2ago"
    "github.com/rs/zerolog/log"
)

func main() {
    // Create a new client
    client := a2ago.NewClient("http://localhost:8089")
    
    // Send a message
    response, err := client.SendMessage("Hello, Agent!")
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to send message")
    }
    
    log.Info().Str("response", response).Msg("Received response")
}
```

## Examples

The repository includes several examples to help you get started:

### Simple Client

The `simple-client` example demonstrates how to use the `a2ago` library to connect to an A2A server, send messages, and query task status.

### Server Hello World

The `server-hello-world` example shows how to create a basic A2A server that can receive messages and respond to clients.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgments

This project is a Go implementation of the [A2A4J (Agent-to-Agent for Java)](https://github.com/a2a4j/a2a4j) framework. We would like to express our sincere gratitude to the A2A4J team for their excellent work and for making their project open source. This Go version was developed with reference to their implementation, aiming to provide the same functionality in the Go ecosystem.

## License

This project is licensed under the Apache License 2.0. 