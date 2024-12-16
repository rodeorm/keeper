package repo

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
)

func (s *postgresStorage) AddBinaryByUser(context.Context, *core.Binary) error {
	return nil
}
func (s *postgresStorage) SelectBinaryByUser(context.Context, *core.User) ([]core.Binary, error) {
	return nil, nil
}
func (s *postgresStorage) UpdateBinaryByUser(context.Context, *core.Binary, *core.User) error {
	return nil
}
func (s *postgresStorage) DeleteBinaryByUser(context.Context, *core.Binary, *core.User) error {
	return nil
}
