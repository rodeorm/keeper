package core

// UserStorager абстрагирует хранилище Пользователей
type UserStorager interface {
	RegUser(*User) error
	UpdateUser(*User) error
	DeleteUser(*User) error
	AuthUser(*User) error
}

//User - пользователи
type User struct {
	Login    string
	Password string
	Name     string
}
