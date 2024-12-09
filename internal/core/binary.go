package core

import "context"

// BinaryStorager - хранилище бинарных файлов
type BinaryStorager interface {
	AddBinaryByUser(context.Context, *Binary) error
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
