package core

// TextStorager абстрагирует хранилище текстовых данных
type TextStorager interface {
	AddTextByUser(*Text) error
	SelectTextByUser(*User) ([]Text, error)
	UpdateTextByUser(*Text, *User) error
	DeleteTextByUser(*Text, *User) error
}

// Text - Произвольные текстовые данные
type Text struct {
	Value string
	Meta  string
	ID    int //Уникальный идентификатор
}
