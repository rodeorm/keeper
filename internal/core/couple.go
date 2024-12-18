package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rodeorm/keeper/internal/crypt"
)

// CoupleStorager  - хранилище логинов-паролей
type CoupleStorager interface {
	AddCoupleByUser(context.Context, *Couple, *User) error
	SelectCoupleByUser(context.Context, *User) ([]Couple, error)
	UpdateCoupleByUser(context.Context, *Couple, *User) error
	DeleteCoupleByUser(context.Context, *Couple, *User) error
}

// Couple - пара логин-пароль
type Couple struct {
	Source   string //16 байт. Источник. Например, сайт, для которого сохраняется логин-пароль
	Login    string //16 байт. Логин
	Password string //16 байт. Пароль
	Meta     string //16 байт. Мета
	ID       int    //8 байт. Идентификатор
}

// DecryptCouple расшифровывает пары логин-пароль
func DecryptCouple(data, password []byte) (*Couple, error) {
	// Расшифровываем данные
	decrypted, err := crypt.Decrypt(data, password)
	if err != nil {
		return nil, fmt.Errorf("не получилось")
	}

	// Преобразуем JSON обратно в структуру
	var decryptedCouple *Couple
	if err := json.Unmarshal(decrypted, decryptedCouple); err != nil {
		return nil, fmt.Errorf("не получилось")
	}
	return decryptedCouple, nil
}
