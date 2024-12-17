/*
		Package core отражает предметную область.
		Для минимизации размера padding bytes, все fields определены от highest allocation to lowest allocation.
	    Это положит любые обязательные padding bytes на "дно" структур и уменьшит общий размер обязательных padding bytes
*/
package core

import "context"

// UserStorager абстрагирует хранилище Пользователей
type UserStorager interface {
	//RegUser регистрирует пользователя
	RegUser(context.Context, *User) error
	//UpdateUser обновляет данные пользователя
	UpdateUser(context.Context, *User) error
	//DeleteUser удаляет пользователя
	DeleteUser(context.Context, *User) error
	// AuthUser аутентифицирует пользователя по логину-паролю для дальнейшей авторизации
	AuthUser(context.Context, *User) bool
	// VerifyUser подтверждение одноразовый пароль для пользователя
	VerifyUserOTP(ctx context.Context, otpLiveTime int, u *User) bool
}

//User - пользователь
type User struct {
	Login    string // Логин
	Password string // Пароль
	OTP      string // Одноразовый пароль
	Name     string // Имя
	Phone    string // Номер телефона
	Email    string // Адрес электронной почты
	ID       int    // Уникальный идентификатор
	Verified bool   // Контактные данные подтверждены
}
