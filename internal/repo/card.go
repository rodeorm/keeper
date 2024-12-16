package repo

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
)

func (s *postgresStorage) AddCardByUser(context.Context, *core.Card) error {
	return nil
}
func (s *postgresStorage) SelectCardByUser(context.Context, *core.User) ([]core.Card, error) {
	return nil, nil
}
func (s *postgresStorage) UpdateCardByUser(context.Context, *core.Card, *core.User) error {
	return nil
}
func (s *postgresStorage) DeleteCardByUser(context.Context, *core.Card, *core.User) error {
	return nil
}
