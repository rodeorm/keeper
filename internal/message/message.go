package message

import (
	"context"
	"database/sql"
	"path/filepath"

	"gopkg.in/gomail.v2"
)

// EmailMessage сообщение, отправляемое по электронной почте
type EmailMessage struct {
	Message
}

// Message сообщение с с OneTimePassword
type Message struct {
	SendedTime sql.NullTime //32 байта. Время, когда сообщение было отправлено

	Login           string //16 байт. Логин, для которого было отправлено сообщение
	Destination     string //16 байт. Адрес назначения (например, адрес электронной почты)
	OneTimePassword string //16 байт. Одноразовый пароль
	Attachment      string //16 байт. Путь к вложению

	ID int //8 байт. Идентификатор
}

type MessageStorager interface {
	CreateMessage(context.Context, *Message) error
	ReadUnsendedMessages(context.Context) ([]Message, error)
	UpdateMessage(context.Context, *Message) error
}

// personifyMessage персонализирует email сообщения
func personifyMessage(folder string, p string, m *Message) (string, error) {
	return "", nil
}

func newEmail(from, destination, subject, text, attachment string) *gomail.Message {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", destination)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", text)
	attachPath, err := filepath.Abs(filepath.Join(".", "static", "img", attachment))
	if err == nil {
		m.Attach(attachPath)
	}
	return m
}
