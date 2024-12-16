package repo

import (
	"context"
	"fmt"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/logger"
	"go.uber.org/zap"
)

func (s *postgresStorage) AddMessage(ctx context.Context, m *core.Message) error {
	//`INSERT INTO cmn.Emails (UserID, OTP, Email) SELECT $1, $2, $3`
	err := s.preparedStatements["AddEmail"].QueryRowContext(ctx, m.UserID, m.OneTimePassword, m.Destination).Scan(m.ID)
	if err != nil {
		return err
	}
	logger.Log.Info("Добавлено сообщение",
		zap.String(m.Destination, fmt.Sprintf("С идентификатором: %d", m.ID)),
	)
	return nil
}
func (s *postgresStorage) SelectUnsendedMessages(context.Context) ([]core.Message, error) {
	return nil, nil
}

func (s *postgresStorage) UpdateMessage(context.Context, *core.Message) error {
	return nil
}
