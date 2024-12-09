package core

import "context"

// TextStorager абстрагирует хранилище текстовых данных
type TextStorager interface {
	AddTextByUser(context.Context, *Text) error
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
