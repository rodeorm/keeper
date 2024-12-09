package sender

import (
	"sync"

	"github.com/rodeorm/keeper/internal/core"
)

// Queue - очередь на отправку сообщений
type Queue struct {
	ch chan *core.Message // Канал для отправки сообщений
}

// Push помещает сообщение в очередь
func (q *Queue) Push(ms *core.Message) error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		q.ch <- ms
		wg.Done()
	}()

	wg.Wait()
	return nil
}

// NewQueue создает новую очередь сообщений размером n
func NewQueue(n int) *Queue {
	return &Queue{
		ch: make(chan *core.Message, n),
	}
}

// PopWait извлекает сообщение из очереди на отправку
func (q *Queue) PopWait() *core.Message {
	select {
	case val := <-q.ch:
		return val
	default:
		return nil
	}
}
