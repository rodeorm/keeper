package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rodeorm/keeper/internal/crypt"
)

// TextStorager абстрагирует хранилище текстовых данных
type TextStorager interface {
	AddTextByUser(context.Context, *Text, *User) error
	SelectTextByUser(context.Context, *User) ([]Text, error)
	UpdateTextByUser(context.Context, *Text, *User) error
	DeleteTextByUser(context.Context, *Text, *User) error
}

// Text - Произвольные текстовые данные
type Text struct {
	Value string //16 байт. Значение
	Meta  string //16 байт. Мета
	ID    int    //8 байт. Уникальный идентификатор
}

// DecryptText расшифровывает текст
func DecryptText(data, password []byte) (*Text, error) {
	// Расшифровываем данные
	decrypted, err := crypt.Decrypt(data, password)
	if err != nil {
		return nil, fmt.Errorf("не получилось")
	}

	// Преобразуем JSON обратно в структуру
	var decryptedText *Text
	if err := json.Unmarshal(decrypted, decryptedText); err != nil {
		return nil, fmt.Errorf("не получилось")
	}
	return decryptedText, nil
}
