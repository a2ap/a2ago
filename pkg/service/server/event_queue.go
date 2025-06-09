package server

import (
	"fmt"
	"sync"
)

// EventQueue represents a queue for managing events
// 支持热流：EnqueueEvent 实时推送到 channel，AsFlux 返回 channel
// 关闭时关闭 channel
// 兼容历史事件回放

type EventQueue struct {
	events   []interface{}
	mu       sync.RWMutex
	closed   bool
	children []*EventQueue

	eventsCh chan interface{} // 新增：事件热流 channel
	once     sync.Once        // 保证只关闭一次
}

// NewEventQueue creates a new EventQueue
func NewEventQueue() *EventQueue {
	return &EventQueue{
		events:   make([]interface{}, 0),
		children: make([]*EventQueue, 0),
		eventsCh: make(chan interface{}, 32), // 带缓冲，防止阻塞
	}
}

// EnqueueEvent enqueues an event to this queue and all its children
func (q *EventQueue) EnqueueEvent(event interface{}) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return fmt.Errorf("queue is closed")
	}

	q.events = append(q.events, event)

	// 推送到热流 channel
	select {
	case q.eventsCh <- event:
	default:
		// 如果满了就丢弃，或可扩展为阻塞/扩容
	}

	// Propagate to children
	for _, child := range q.children {
		child.EnqueueEvent(event)
	}

	return nil
}

// AsFlux returns a channel that emits events from this queue（热流）
func (q *EventQueue) AsFlux() <-chan interface{} {
	return q.eventsCh
}

// Tap taps the event queue to create a new child queue that receives all future events
func (q *EventQueue) Tap() (*EventQueue, error) {
	q.mu.Lock()
	defer q.mu.Unlock()

	childQueue := NewEventQueue()
	q.children = append(q.children, childQueue)
	return childQueue, nil
}

// Close closes the queue for future push events
func (q *EventQueue) Close() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.closed {
		return nil
	}

	q.closed = true
	q.once.Do(func() {
		close(q.eventsCh)
	})

	// Close all child queues
	for _, child := range q.children {
		child.Close()
	}

	return nil
}

// IsClosed checks if the queue is closed
func (q *EventQueue) IsClosed() bool {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.closed
}
