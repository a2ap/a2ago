# 服务器 Hello World 示例

这个示例展示了如何使用 `a2ago` 库构建一个简单的 Agent-to-Agent (A2A) 服务器。它实现了核心的 A2A 服务端组件，能够接收客户端发送的消息，处理任务，并返回响应。

## 特性

- 基本消息处理
- 任务处理
- 响应生成
- WebSocket 支持实时通信

## 先决条件

- Go 1.21 或更高版本
- 基本的 Go 编程知识
- 熟悉 A2A 协议

## 快速开始

1.  确保你已克隆 `a2ago` 仓库并进入其根目录。

2.  进入此示例目录：

```bash
cd examples/server-hello-world
```

3.  安装依赖：

```bash
go mod download
```

4.  运行服务器：

```bash
go run main.go
```

服务器将在 `http://localhost:8089` 启动。

## 测试服务器

你可以使用 `simple-client` 示例或直接发送 HTTP 请求来测试服务器。服务器实现了以下端点：

- `POST /message`：向服务器发送消息
- `GET /task/{taskId}`：查询任务状态
- `GET /ws`：WebSocket 端点，用于实时通信

## 代码结构

- `main.go`：服务器初始化和配置
- `handler.go`：请求处理器和业务逻辑
- `model.go`：数据模型和类型

## 许可证

此示例是 A2AGO 项目的一部分，采用 MIT 许可证。 