package core

import "time"

//Data - структура для отражения данных, полученных из БД
type Data struct {
	CreatedDate time.Time //Дата создания
	ByteData    []byte    //Зашифрованные бинарные данные в БД
	Name        string
	Meta        string
	ID          int //Идентификатор
	UserID      int //Идентификатор пользователя
	TypeID      int //Идентификатор типа
}
