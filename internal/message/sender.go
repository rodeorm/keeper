package message

import (
	"fmt"
	"time"

	"github.com/rodeorm/keeper/internal/logger"
	"gopkg.in/gomail.v2"
)

// Sender - рабочий, отправляющий сообщения
type Sender struct {
	messageStorage MessageStorager //16 байт. Хранилище сообщений
	from           string          //16 байт. Отправитель
	queue          *Queue          //8 байт. Очередь сообщений к отправке
	dialer         *gomail.Dialer  //8 байт. Отправитель
	ID             int             //8 байт. Идентификатор воркера
}

// NewSender создает новый Sender
func NewSender(id int, queue *Queue, storage MessageStorager, smtpPort int, smtpServer, smtpLogin, smtpPassword string) *Sender {
	s := Sender{
		ID:             id,
		queue:          queue,
		messageStorage: storage,
	}

	s.dialer = gomail.NewDialer(smtpServer, smtpPort, smtpLogin, smtpPassword)

	return &s
}

// StartSending начинает отправку сообщений
func (s *Sender) StartSending(exit chan struct{}) {
	logger.Info("Sender. Send", fmt.Sprintf("стартовал сендер %d", s.ID), time.Now().String())

	for {
		select {
		case _, ok := <-exit:
			if !ok {
				return
			}
		default:
			ms := s.queue.PopWait()

			if ms == nil {
				continue
			}
			err := s.Send(ms)
			if err != nil {
				logger.Error("Sender. Send", fmt.Sprintf("ошибка при работе сендера %d", s.ID), err.Error())
				continue
			}
			logger.Info("Sender. Send", fmt.Sprintf("сендер %d отправил сообщения", s.ID), ms.Destination)
		}
	}
}

// Send отправляет сообщение
func (s *Sender) Send(ms *Message) error {

	text, err := personifyMessage("mail", "approve", ms)
	if err != nil {
		return nil
	}

	email := newEmail(s.from, ms.Destination, ms.Login, text, ms.Attachment)

	if err := s.dialer.DialAndSend(email); err != nil {
		logger.Info("Send", "отправлено сообщение", text)
		return err
	}

	return nil
}
