package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rodeorm/keeper/internal/crypt"
)

// BinaryStorager - хранилище бинарных файлов
type BinaryStorager interface {
	AddBinaryByUser(context.Context, *Binary, *User) error
	SelectBinaryByUser(context.Context, *User) ([]Binary, error)
	UpdateBinaryByUser(context.Context, *Binary, *User) error
	DeleteBinaryByUser(context.Context, *Binary, *User) error
}

// Binary - произвольный бинарный файл
type Binary struct {
	Value []byte //24 байта. Значение бинарное
	Meta  string //16 байт. Мета информация
	ID    int    //8 байт. Идентификатор
}

// DecryptBinary расшфировывает бинарники
func DecryptBinary(data, password []byte) (*Binary, error) {
	// Расшифровываем данные
	decrypted, err := crypt.Decrypt(data, password)
	if err != nil {
		return nil, fmt.Errorf("не получилось")
	}

	// Преобразуем JSON обратно в структуру
	var decryptedBin *Binary
	if err := json.Unmarshal(decrypted, decryptedBin); err != nil {
		return nil, fmt.Errorf("не получилось")
	}
	return decryptedBin, nil
}
