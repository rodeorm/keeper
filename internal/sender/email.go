package sender

import (
	"path/filepath"

	"github.com/rodeorm/keeper/internal/core"
	"gopkg.in/gomail.v2"
)

// EmailMessage сообщение, отправляемое по электронной почте
type EmailMessage struct {
	Message core.Message
	gms     *gomail.Message
}

func newEmail(from string, m *core.Message) *EmailMessage {
	ems := &EmailMessage{gms: gomail.NewMessage()}
	ems.gms.SetHeader("From", from)
	ems.gms.SetHeader("To", m.Destination)
	ems.gms.SetHeader("Subject", m.Login)
	ems.gms.SetBody("text/html", m.Text)
	attachPath, err := filepath.Abs(filepath.Join(".", "static", "img", m.Attachment))
	if err == nil {
		ems.gms.Attach(attachPath)
	}
	return ems
}
