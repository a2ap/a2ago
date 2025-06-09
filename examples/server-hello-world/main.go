package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/a2a4j/a2ago/examples/server-hello-world/agent"
	"github.com/a2a4j/a2ago/internal/jsonrpc"
	"github.com/a2a4j/a2ago/internal/model"
	"github.com/a2a4j/a2ago/internal/service/server/impl"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 创建 agent card（Agent 元信息）
	agentCard := &model.AgentCard{
		Name:        "A2A Go Server",
		Description: "A sample A2A agent implemented in Go",
		Version:     "1.0.0",
		URL:         "http://localhost:8089",
		Capabilities: &model.AgentCapabilities{
			Streaming:              true,
			PushNotifications:      false,
			StateTransitionHistory: true,
		},
		Skills:             []*model.AgentSkill{},
		DefaultInputModes:  []string{"text"},
		DefaultOutputModes: []string{"text"},
	}

	// 2. 创建任务管理器（TaskManager）
	taskManager := impl.NewInMemoryTaskManager(impl.NewInMemoryTaskStore())

	// 3. 创建事件队列管理器（QueueManager）
	queueManager := impl.NewInMemoryQueueManager()

	// 4. 手动注入 DemoAgentExecutor，并传入 queueManager
	//    这是 Go 端等价于 Java/Spring 自动装配的关键步骤
	agentExecutor := agent.NewDemoAgentExecutor(queueManager)

	// 5. 创建 A2A Server，并注入 taskManager、queueManager、agentExecutor、agentCard
	a2aServer := impl.NewDefaultA2AServer(taskManager, queueManager, agentExecutor, agentCard)

	// 6. 创建 Dispatcher，并注入 A2A Server
	dispatcher := impl.NewDefaultDispatcher(a2aServer)

	// 7. 创建 Gin 路由
	router := gin.Default()

	// 8. 配置 CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 9. 注册路由
	router.GET("/.well-known/agent.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, a2aServer.GetSelfAgentCard())
	})

	router.POST("/a2a/server", func(c *gin.Context) {
		// 检查是否为 SSE 流式请求
		if c.GetHeader("Accept") == "text/event-stream" {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			var request jsonrpc.JSONRPCRequest
			if err := c.ShouldBindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			responses, err := dispatcher.DispatchStream(&request)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.Stream(func(w io.Writer) bool {
				if response, ok := <-responses; ok {
					c.SSEvent("task-update", response)
					return true
				}
				return false
			})
		} else {
			var request jsonrpc.JSONRPCRequest
			if err := c.ShouldBindJSON(&request); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			response := dispatcher.Dispatch(&request)
			c.JSON(http.StatusOK, response)
		}
	})

	// 10. 启动服务
	log.Println("Starting A2A server on http://localhost:8089")
	if err := router.Run(":8089"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
