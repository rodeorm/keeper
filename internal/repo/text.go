package repo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/crypt"
)

func (s *postgresStorage) AddTextByUser(ctx context.Context, t *core.Text, u *core.User) error {
	d, err := crypt.CryptData(t, s.CryptKey)
	if err != nil {
		return err
	}
	err = s.preparedStatements["AddByte"].GetContext(ctx, t, u.ID, core.TextType, d, time.Now(), "", t.Meta)
	if err != nil {
		return err
	}
	return nil
}
func (s *postgresStorage) SelectAllTextsByUser(ctx context.Context, u *core.User) ([]core.Text, error) {
	cyphTexts := make([]core.Data, 0)
	ts := make([]core.Text, 0)
	err := s.preparedStatements["SelectAllBytes"].SelectContext(ctx, &cyphTexts, u.ID, core.TextType)
	if err != nil {
		return nil, err
	}

	for _, v := range cyphTexts {
		decData, _ := crypt.Decrypt(v.ByteData, s.CryptKey)
		c := core.Text{}
		json.Unmarshal(decData, &c)
		ts = append(ts, c)
	}

	return ts, nil
}
func (s *postgresStorage) UpdateTextByUser(context.Context, *core.Text, *core.User) error {
	return nil
}
func (s *postgresStorage) DeleteTextByUser(context.Context, *core.Text, *core.User) error {
	return nil
}
