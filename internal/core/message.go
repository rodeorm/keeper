package core

import (
	"context"
	"database/sql"

	"github.com/rodeorm/keeper/internal/crypt"
)

// Message - сообщение с с OneTimePassword
// Теоретически сообщение может быть отправлено любым способом: sms, email и т.п. поэтому названия атрибутов максимально неспецифичные
type Message struct {
	SendedTime sql.NullTime //32 байта. Время, когда сообщение было отправлено

	Login           string //16 байт. Логин, для которого было отправлено сообщение
	Destination     string //16 байт. Адрес назначения (например, адрес электронной почты)
	OneTimePassword string //16 байт. Одноразовый пароль
	Attachment      string //16 байт. Путь к вложению
	Text            string //16 байт. Сообщение

	ID     int //8 байт. Идентификатор
	UserID int //8 байт. Идентификатор
}

type MessageStorager interface {
	AddMessage(ctx context.Context, m *Message) error
	SelectUnsendedMessages(context.Context) ([]Message, error)
	UpdateMessage(context.Context, *Message) error
}

// personifyMessage персонализирует сообщение
func PersonifyMessage(folder string, p string, m *Message) error {
	return nil
}

func NewAuthMessage(u *User) (*Message, error) {
	m := &Message{Login: u.Login, Destination: u.Email, UserID: u.ID}
	otp, err := crypt.ReturnOTP(len(u.Login))
	if err != nil {
		return nil, err
	}
	m.OneTimePassword = otp
	return m, nil
}
