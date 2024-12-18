package repo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/crypt"
)

func (s *postgresStorage) AddCardByUser(ctx context.Context, c *core.Card, u *core.User) error {
	d, err := crypt.CryptData(c, s.CryptKey)
	if err != nil {
		return err
	}
	err = s.preparedStatements["AddByte"].GetContext(ctx, c, u.ID, core.CardType, d, time.Now())
	if err != nil {
		return err
	}
	return nil
}
func (s *postgresStorage) SelectCardByUser(ctx context.Context, u *core.User) ([]core.Card, error) {
	cyphCards := make([]core.Data, 0)
	cards := make([]core.Card, 0)
	err := s.preparedStatements["SelectByte"].SelectContext(ctx, &cyphCards, u.ID, core.CardType)
	if err != nil {
		return nil, err
	}

	for _, v := range cyphCards {
		decData, _ := crypt.Decrypt(v.ByteData, s.CryptKey)
		card := core.Card{}
		json.Unmarshal(decData, &card)
		cards = append(cards, card)
	}

	return cards, nil
}
func (s *postgresStorage) UpdateCardByUser(context.Context, *core.Card, *core.User) error {
	return nil
}
func (s *postgresStorage) DeleteCardByUser(context.Context, *core.Card, *core.User) error {
	return nil
}
