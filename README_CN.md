# A2AGO

`a2ago` 是 [A2A4J (Agent-to-Agent for Java)](https://github.com/a2a4j/a2a4j) 框架的 Go 语言实现版本。它提供了用于构建和集成 Agent-to-Agent 交互的 Go 客户端和服务器库。该库旨在通过标准化的消息和任务流程促进不同代理之间的通信。

## 特性

- **标准化消息协议**：实现了 A2A 消息协议，确保代理之间的一致通信。
- **任务流程管理**：支持任务创建、状态跟踪和结果检索。
- **WebSocket 支持**：通过 WebSocket 连接实现实时通信功能。
- **可扩展架构**：易于扩展和定制，适用于不同的使用场景。
- **全面的测试**：包含单元测试和集成测试，确保可靠性。

## 安装

1. 克隆仓库并进入 `a2ago` 目录：

```bash
git clone https://github.com/a2a4j/a2ago.git
cd a2ago
```

2. 安装依赖：

```bash
go mod download
```

## 使用

### 基础示例

以下是一个简单的使用示例：

```go
package main

import (
    "github.com/a2a4j/a2ago"
    "github.com/rs/zerolog/log"
)

func main() {
    // 创建新的客户端
    client := a2ago.NewClient("http://localhost:8089")
    
    // 发送消息
    response, err := client.SendMessage("Hello, Agent!")
    if err != nil {
        log.Fatal().Err(err).Msg("发送消息失败")
    }
    
    log.Info().Str("response", response).Msg("收到响应")
}
```

## 示例

仓库中包含多个示例，帮助您快速上手：

### 简单客户端

`simple-client` 示例展示了如何使用 `a2ago` 库连接到 A2A 服务器、发送消息和查询任务状态。

### 服务器 Hello World

`server-hello-world` 示例展示了如何创建一个基本的 A2A 服务器，该服务器可以接收消息并响应客户端。

## 贡献

欢迎贡献代码！请随时提交 Pull Request。

## 致谢

本项目是 [A2A4J (Agent-to-Agent for Java)](https://github.com/a2a4j/a2a4j) 框架的 Go 实现版本。我们要向 A2A4J 团队表示衷心的感谢，感谢他们的出色工作和开源精神。本 Go 版本参考了他们的实现，旨在为 Go 生态系统提供相同的功能。

## 许可证

本项目采用 Apache License 2.0 协议。 