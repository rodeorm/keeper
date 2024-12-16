package repo

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
)

// StartSession начинает новую сессию
func (s *postgresStorage) StartSession(context.Context, *core.User) (*core.Session, error) {
	return nil, nil
}

// UpdateSession обновляет данные сессии
func (s *postgresStorage) UpdateSession(context.Context, *core.Session) error {
	return nil
}

// EndSession закрывает сессию
func (s *postgresStorage) EndSession(context.Context, *core.Session) error {
	return nil
}
