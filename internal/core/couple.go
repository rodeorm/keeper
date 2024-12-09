package core

import "context"

// CoupleStorager  - хранилище логинов-паролей
type CoupleStorager interface {
	AddCoupleByUser(context.Context, *Couple) error
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
