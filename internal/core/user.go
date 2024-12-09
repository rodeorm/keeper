package core

import "context"

// UserStorager абстрагирует хранилище Пользователей
type UserStorager interface {
	RegUser(context.Context, *User) error
	UpdateUser(context.Context, *User) error
	DeleteUser(context.Context, *User) error
	AuthUser(context.Context, *User) error
}

//User - пользователи
type User struct {
	Login    string // Логин
	Password string // Пароль
	Name     string // Имя
	Email    string // Адрес электронной почты
}
