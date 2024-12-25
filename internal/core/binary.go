package core

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rodeorm/keeper/internal/crypt"
)

// BinaryStorager - хранилище бинарных файлов
type BinaryStorager interface {
	AddBinaryByUser(context.Context, *Binary, *User) error
	SelectAllBinariesByUser(context.Context, *User) ([]Binary, error)
	SelectBinaryByUser(context.Context, *Binary, *User) error
	UpdateBinaryByUser(context.Context, *Binary, *User) error
	DeleteBinaryByUser(context.Context, *Binary, *User) error
}

// Binary - произвольный бинарный файл
type Binary struct {
	Value []byte //24 байта. Значение бинарное
	Name  string //16 байт. Наименование
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

// GetBinaryFromFile возвращает Binary на основании пути к файлу
func GetBinaryFromFile(fp string) (*Binary, error) {
	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, err
	}
	return &Binary{Value: data, Name: filepath.Base(fp)}, nil
}

// SaveBinaryToFile сохраняет файл из Binary
func SaveBinaryToFile(b Binary, filePath string) error {
	return os.WriteFile(filePath+"/"+b.Name, b.Value, 0666)
}
