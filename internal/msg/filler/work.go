package filler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/logger"
	"go.uber.org/zap"
)

// Filler - рабочий, заполняющий очередь сообщений
type Filler struct {
	messageStorage core.MessageStorager //16 байт. Хранилище сообщений
	queue          *core.Queue          //8 байт. Очередь сообщений
	ID             int                  //8 байт. Идентификатор воркера
	period         int                  //8 байт. Периодичность наполнения сообщений
}

func Start(config *cfg.Server, wg *sync.WaitGroup, exit chan struct{}) {
	// Асинхронно запускаем наполнитель очереди
	s := NewFiller(
		config.MessageQueue,
		config.MessageStorager,
		config.QueueFillPeriod,
	)

	go s.StartFilling(exit, wg)
}

// StartFilling начинает наполнение очереди
func (f *Filler) StartFilling(exit chan struct{}, wg *sync.WaitGroup) {
	logger.Log.Info("StartFilling",
		zap.String("Филлер стартовал", "Успешно"))
	ctx := context.TODO()
	for {
		select {
		case _, ok := <-exit:
			if !ok {
				//Нет смысла ждать наполнения очереди, поэтому дефолт не жду
				logger.Log.Info("StartFilling",
					zap.String("Филлер изящно завершил дела", "Успешно"))
				wg.Done()
				return
			}
		default:
			go func() {
				msgs, err := f.messageStorage.SelectUnsendedMessages(ctx)

				if err != nil {
					logger.Log.Error("StartFilling",
						zap.String("ошибка при получении сообщений к отправке", err.Error()),
					)
				}

				for i, v := range msgs {
					err = f.queue.Push(&v)
					if err != nil {
						logger.Log.Error("StartFilling",
							zap.String(fmt.Sprintf("ошибка при заполнении очереди для сообщения %d", i), err.Error()),
						)
						return
					}
					v.Queued = true
					f.messageStorage.UpdateMessage(ctx, &v)
				}
			}()
			time.Sleep(time.Duration(f.period) * time.Second)
		}
	}
}
