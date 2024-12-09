package sender

import (
	"testing"

	"github.com/rodeorm/keeper/internal/core"
)

func TestNewEmail(t *testing.T) {
	from := "test@example.com"
	msg := &core.Message{
		Destination: "recipient@example.com",
		Login:       "Test Subject",
		Text:        "This is a test email body.",
	}

	emailMessage := newEmail(from, msg)

	if emailMessage.gms.GetHeader("From")[0] != from {
		t.Errorf("expected From header to be %s, got %s", from, emailMessage.gms.GetHeader("From"))
	}

	if emailMessage.gms.GetHeader("To")[0] != msg.Destination {
		t.Errorf("expected To header to be %s, got %s", msg.Destination, emailMessage.gms.GetHeader("To")[0])
	}

	if emailMessage.gms.GetHeader("Subject")[0] != msg.Login {
		t.Errorf("expected Subject header to be %s, got %s", msg.Login, emailMessage.gms.GetHeader("Subject"))
	}

}
