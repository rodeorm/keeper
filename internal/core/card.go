package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rodeorm/keeper/internal/crypt"
)

// CardStorager абстрагирует хранилище данных банковских карт
type CardStorager interface {
	AddCardByUser(context.Context, *Card) error
	SelectCardByUser(context.Context, *User) ([]Card, error)
	UpdateCardByUser(context.Context, *Card, *User) error
	DeleteCardByUser(context.Context, *Card, *User) error
}

// Card - Данные банковских карт
type Card struct {
	CardNumber string //16 байт. 16-тизначный номер (в некоторых случаях 18-тизначный — включает зашифрованную информацию о банке-эмитенте)
	OwnerName  string //16 байт. Имя и фамилия владельца на латинице
	Meta       string //16 байт.
	ID         int    //8 байт. Уникальный идентификатор
	ExpMonth   int    //8 байт. Срок действия: месяц
	ExpYear    int    //8 байт. Срок действия: год
	CVC        int    //8 байт. CVC или CVV2 — код из 3 или 4 цифр для совершения интернет-платежей, расположенный на обратной стороне
}

// DecryptCard расшифровывает данные карт
func DecryptCard(data, password []byte) (*Card, error) {
	// Расшифровываем данные
	decrypted, err := crypt.Decrypt(data, password)
	if err != nil {
		return nil, fmt.Errorf("не получилось")
	}

	// Преобразуем JSON обратно в структуру
	var decryptedCard *Card
	if err := json.Unmarshal(decrypted, decryptedCard); err != nil {
		return nil, fmt.Errorf("не получилось")
	}
	return decryptedCard, nil
}
