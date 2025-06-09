# 简单客户端示例

这个示例展示了如何使用 `a2ago` 库构建一个简单的 Agent-to-Agent (A2A) 客户端。它演示了如何连接到 A2A 服务器，发送消息，并查询任务的实时状态。

## 特性

- 基本消息发送
- 任务状态查询
- WebSocket 连接实时更新
- 错误处理和重试逻辑

## 先决条件

- Go 1.21 或更高版本
- 一个正在运行的 A2A 服务器，例如 `a2ago/examples/server-hello-world` 示例，监听在 `http://localhost:8089`

## 快速开始

1.  确保你已克隆 `a2ago` 仓库并进入其根目录。

2.  进入此示例目录：

```bash
cd examples/simple-client
```

3.  安装依赖：

```bash
go mod download
```

4.  运行客户端：

```bash
go run main.go
```

客户端将连接到服务器，发送测试消息，并显示任务状态和响应。

## 代码结构

- `main.go`：客户端初始化和主要逻辑
- `client.go`：A2A 客户端实现
- `model.go`：数据模型和类型

## 许可证

此示例是 A2AGO 项目的一部分，采用 MIT 许可证。 