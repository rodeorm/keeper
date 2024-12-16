package repo

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
)

func (s *postgresStorage) AddCoupleByUser(context.Context, *core.Couple) error {
	return nil
}
func (s *postgresStorage) SelectCoupleByUser(context.Context, *core.User) ([]core.Couple, error) {
	return nil, nil
}
func (s *postgresStorage) UpdateCoupleByUser(context.Context, *core.Couple, *core.User) error {
	return nil
}
func (s *postgresStorage) DeleteCoupleByUser(context.Context, *core.Couple, *core.User) error {
	return nil
}
