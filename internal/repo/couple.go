package repo

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/crypt"
)

func (s *postgresStorage) AddCoupleByUser(ctx context.Context, c *core.Couple, u *core.User) error {
	d, err := crypt.CryptData(c, s.CryptKey)
	if err != nil {
		return err
	}
	err = s.preparedStatements["AddByte"].GetContext(ctx, c, u.ID, core.CoupleType, d, time.Now(), "", c.Meta)
	if err != nil {
		return err
	}
	return nil
}
func (s *postgresStorage) SelectAllCouplesByUser(ctx context.Context, u *core.User) ([]core.Couple, error) {
	cyphCouples := make([]core.Data, 0)
	cs := make([]core.Couple, 0)
	err := s.preparedStatements["SelectAllBytes"].SelectContext(ctx, &cyphCouples, u.ID, core.CoupleType)
	if err != nil {
		return nil, err
	}

	for _, v := range cyphCouples {
		decData, _ := crypt.Decrypt(v.ByteData, s.CryptKey)
		c := core.Couple{}
		json.Unmarshal(decData, &c)
		cs = append(cs, c)
	}

	log.Println("SelectAllCouplesByUser", cs)

	return cs, nil
}
func (s *postgresStorage) UpdateCoupleByUser(context.Context, *core.Couple, *core.User) error {
	return nil
}
func (s *postgresStorage) DeleteCoupleByUser(context.Context, *core.Couple, *core.User) error {
	return nil
}
