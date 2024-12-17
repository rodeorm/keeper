package repo

import (
	"context"
	"fmt"
	"log"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/logger"
	"go.uber.org/zap"
)

func (s *postgresStorage) AddMessage(ctx context.Context, m *core.Message) error {
	//`INSERT INTO cmn.Emails (UserID, OTP, Email) SELECT $1, $2, $3`
	err := s.preparedStatements["AddEmail"].GetContext(ctx, &m.ID, m.UserID, m.OTP, m.Destination)
	if err != nil {
		return err
	}
	logger.Log.Info("Добавлено сообщение",
		zap.String(m.Destination, fmt.Sprintf("С идентификатором: %d", m.ID)),
	)
	return nil
}
func (s *postgresStorage) SelectUnsendedMessages(ctx context.Context) ([]core.Message, error) {
	ms := make([]core.Message, 0)

	err := s.preparedStatements["SelectEmailForSending"].SelectContext(ctx, &ms)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

func (s *postgresStorage) UpdateMessage(ctx context.Context, c *core.Message) error {
	//UPDATE cmn.Emails SET OTP = $2, Email = $3, SendedDate = $4, Used = $5, Queued = $6 WHERE id = $1;
	_, err := s.preparedStatements["UpdateEmail"].QueryContext(ctx, c.ID, c.OTP, c.Destination, c.SendedDate, c.Used, c.Queued)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
