package repo

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/crypt"
)

func (s *postgresStorage) AddBinaryByUser(ctx context.Context, b *core.Binary, u *core.User) error {
	d, err := crypt.CryptData(b, s.CryptKey)
	if err != nil {
		return err
	}
	err = s.preparedStatements["AddByte"].GetContext(ctx, b, u.ID, core.BinaryType, d, time.Now(), b.Name, b.Meta)
	if err != nil {
		return err
	}
	return nil
}
func (s *postgresStorage) SelectAllBinariesByUser(ctx context.Context, u *core.User) ([]core.Binary, error) {
	cyphBins := make([]core.Data, 0)
	bs := make([]core.Binary, 0)
	err := s.preparedStatements["SelectAllBytes"].SelectContext(ctx, &cyphBins, u.ID, core.BinaryType)
	if err != nil {
		return nil, err
	}

	for _, v := range cyphBins {
		decData, _ := crypt.Decrypt(v.ByteData, s.CryptKey)
		b := core.Binary{}
		json.Unmarshal(decData, &b)
		b.ID = v.ID
		b.Name = v.Name
		b.Meta = v.Meta
		bs = append(bs, b)
	}

	return bs, nil
}

func (s *postgresStorage) SelectBinaryByUser(ctx context.Context, b *core.Binary, u *core.User) error {
	cyphBin := core.Data{}

	//SELECT id, bytedata, createddate FROM dbo.Data WHERE userid = $1 AND id = $2;
	err := s.preparedStatements["SelectByte"].GetContext(ctx, &cyphBin, u.ID, b.ID)
	if err != nil {
		log.Println("SelectBinaryByUser HERE WE ARE", err, b.ID, u.ID)
		return err
	}
	decData, err := crypt.Decrypt(cyphBin.ByteData, s.CryptKey)
	if err != nil {
		log.Println("SelectBinaryByUser 2 HERE WE ARE")
		return err
	}
	json.Unmarshal(decData, &b)

	return nil
}

func (s *postgresStorage) UpdateBinaryByUser(ctx context.Context, b *core.Binary, u *core.User) error {
	return nil
}

func (s *postgresStorage) DeleteBinaryByUser(ctx context.Context, b *core.Binary, u *core.User) error {
	return nil
}
