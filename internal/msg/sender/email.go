package sender

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/rodeorm/keeper/internal/core"
	"gopkg.in/gomail.v2"
)

// EmailMessage сообщение, отправляемое по электронной почте
type EmailMessage struct {
	gms *gomail.Message
}

func (s Sender) NewEmail(from string, m *core.Message) *EmailMessage {
	text := customiseMail("mail", "otp", m)
	ems := &EmailMessage{gms: gomail.NewMessage()}
	ems.gms.SetHeader("From", s.from)
	ems.gms.SetHeader("To", m.Destination)
	ems.gms.SetHeader("Subject", "Одноразовый пароль от keeper")
	ems.gms.SetBody("text/html", text)
	attachPath, err := filepath.Abs(filepath.Join(".", "static", "img", s.fileName))
	if err == nil {
		ems.gms.Attach(attachPath)
	}
	return ems
}

func customiseMail(folder string, page string, param interface{}) string {
	templatePath, _ := filepath.Abs(fmt.Sprintf("./%s/%s.html", folder, page))
	mail, _ := template.ParseFiles(templatePath)
	var body bytes.Buffer
	mail.Execute(&body, param)
	return body.String()
}
