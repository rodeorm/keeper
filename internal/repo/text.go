package repo

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
)

func (s *postgresStorage) AddTextByUser(context.Context, *core.Text) error {
	return nil
}
func (s *postgresStorage) SelectTextByUser(context.Context, *core.User) ([]core.Text, error) {
	return nil, nil
}
func (s *postgresStorage) UpdateTextByUser(context.Context, *core.Text, *core.User) error {
	return nil
}
func (s *postgresStorage) DeleteTextByUser(context.Context, *core.Text, *core.User) error {
	return nil
}
