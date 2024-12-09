package core

//CardStorager абстрагирует хранилище данных банковских карт
type CardStorager interface {
	AddCardByUser(*Card) error
	SelectCardByUser(*User) ([]Card, error)
	UpdateCardByUser(*Card, *User) error
	DeleteCardByUser(*Card, *User) error
}

// Card - Данные банковских карт
type Card struct {
	CardNumber string //16-тизначный номер (в некоторых случаях 18-тизначный — включает зашифрованную информацию о банке-эмитенте)
	OwnerName  string //Имя и фамилия владельца на латинице
	Meta       string
	ID         int //Уникальный идентификатор
	ExpMonth   int //Срок действия: месяц
	ExpYear    int //Срок действия: год
	CVC        int //CVC или CVV2 — код из 3 или 4 цифр для совершения интернет-платежей, расположенный на обратной стороне
}
