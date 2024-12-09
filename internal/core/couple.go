package core

// CoupleStorager  - хранилище логинов-паролей
type CoupleStorager interface {
	AddCoupleByUser(*Couple) error
	SelectCoupleByUser(*User) ([]Couple, error)
	UpdateCoupleByUser(*Couple, *User) error
	DeleteCoupleByUser(*Couple, *User) error
}

// Couple - пара логин-пароль
type Couple struct {
	Source   string
	Login    string
	Password string
	Meta     string
	ID       int
}
