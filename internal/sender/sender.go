package sender

import (
	"fmt"
	"sync"
	"time"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/logger"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

// Sender - рабочий, отправляющий сообщения
type Sender struct {
	messageStorage core.MessageStorager //16 байт. Хранилище сообщений
	from           string               //16 байт. Отправитель
	queue          *core.Queue          //8 байт. Очередь сообщений
	dialer         *gomail.Dialer       //8 байт. Отправитель
	ID             int                  //8 байт. Идентификатор воркера
	period         int                  //8 байт. Периодичность отправки сообщений
}

// NewSender создает новый Sender
// Каждый Sender может рассылать сообщения через свой собственный smtp сервер
func NewSender(queue *core.Queue, storage core.MessageStorager, id, smtpPort, prd int, smtpServer, smtpLogin, smtpPassword string) *Sender {
	s := Sender{
		ID:             id,
		queue:          queue,
		messageStorage: storage,
		period:         prd,
	}

	s.dialer = gomail.NewDialer(smtpServer, smtpPort, smtpLogin, smtpPassword)

	return &s
}

func SenderStart(config *cfg.Server, wg *sync.WaitGroup, exit chan struct{}) {
	for i := range config.SenderQuantity {
		// Асинхронно запускаем email сендеры
		s := NewSender(
			config.MessageQueue,
			config.MessageStorager,
			i,
			config.SMTPPort,
			config.MessagePeriod,
			config.SMTPServer,
			config.SMTPLogin,
			config.SMTPPass,
		)

		go s.StartSending(exit, wg)
	}
}

// StartSending начинает отправку сообщений
func (s *Sender) StartSending(exit chan struct{}, wg *sync.WaitGroup) {
	logger.Info("StartSending", fmt.Sprintf("Сендер %d", s.ID), "стартовал")

	var wg_w sync.WaitGroup

	for {

		select {
		case _, ok := <-exit:
			if !ok {
				wg_w.Wait()
				logger.Log.Info("StartSending",
					zap.String(fmt.Sprintf("Сендер %d", s.ID), "изящно завершил дела"),
				)
				wg.Done()
				return
			}
		default:

			wg_w.Add(1)

			go func() {
				logger.Log.Info("StartSending",
					zap.String(fmt.Sprintf("Сендер %d", s.ID), "делает делишки"),
				)

				ms := s.queue.PopWait()

				if ms == nil {
					logger.Log.Info("StartSending",
						zap.String(fmt.Sprintf("Сендер %d", s.ID), "сообщений нет"),
					)
					wg_w.Done()
					return
				}
				err := s.Send(ms)
				if err != nil {
					logger.Log.Error("StartSending",
						zap.String(fmt.Sprintf("Сендер %d", s.ID), err.Error()),
					)
					wg_w.Done()
					return
				}
				logger.Info("StartSending", fmt.Sprintf("Сендер %d", s.ID), fmt.Sprintf("отправил сообщение по адресу %s", ms.Destination))
				wg_w.Done()
			}()
			time.Sleep(time.Duration(s.period) * time.Second)
		}

	}
}

// Send отправляет сообщение
func (s *Sender) Send(ms *core.Message) error {

	err := core.PersonifyMessage("mail", "approve", ms)
	if err != nil {
		return nil
	}

	email := newEmail(s.from, ms)

	if err := s.dialer.DialAndSend(email.gms); err != nil {
		logger.Log.Info("Send",
			zap.String(fmt.Sprintf("Сендер %d отправил сообщение", s.ID), ms.Text),
		)
		return err
	}

	return nil
}
