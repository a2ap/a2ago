package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/a2a4j/a2ago/internal/model"
	clientimpl "github.com/a2a4j/a2ago/internal/service/client/impl"
	"github.com/a2a4j/a2ago/internal/util"
)

func main() {
	// 创建客户端
	cardResolver := clientimpl.NewHttpCardResolver("http://localhost:8089")
	a2aClient := clientimpl.NewDefaultA2aClient(cardResolver)

	// 确保获取 agent card
	agentCard := a2aClient.RetrieveAgentCard()
	if agentCard == nil {
		log.Fatal("Failed to retrieve agent card")
	}

	// 创建消息
	textPart := model.NewTextPart("Hello, this is a test message")
	message := model.NewMessage(util.GenerateUUID(), "user", []model.Part{textPart})

	// 创建消息发送参数
	metadata := make(map[string]interface{})
	params := model.NewMessageSendParams(message, metadata)

	// 发送消息
	task, err := a2aClient.SendMessage(context.Background(), params)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	fmt.Printf("Message sent successfully, task ID: %s\n", task.ID)

	// 等待任务完成
	time.Sleep(2 * time.Second)

	// 获取任务状态
	taskQueryParams := model.NewTaskQueryParams(task.ID)
	updatedTask, err := a2aClient.GetTask(context.Background(), taskQueryParams)
	if err != nil {
		log.Fatalf("Failed to get task status: %v", err)
	}

	if updatedTask != nil && updatedTask.Status != nil {
		fmt.Printf("Task status: %s\n", updatedTask.Status.State)
		if updatedTask.Status.Message != nil && len(updatedTask.Status.Message.Parts) > 0 {
			if textPart, ok := updatedTask.Status.Message.Parts[0].(*model.TextPart); ok {
				fmt.Printf("Task message: %s\n", textPart.Text)
			}
		}
	} else {
		fmt.Println("Task not found or no status available")
	}
}
