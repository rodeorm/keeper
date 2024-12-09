package core

import (
	"context"
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt"
)

// SessionStorager абстрагирует хранилище данных сессии
type SessionStorager interface {
	StartSession(context.Context, *User) (*Session, error) //Начинает новую сессию
	UpdateSession(context.Context, *Session) error         //Обновляет данные сессии
	EndSession(context.Context, *Session) error            //Закрывает сессию
}

// Session - сессия пользователя, которая хранится централизованно и позволяет завершить любую клиентскую сессию через сервер
type Session struct {
	LastActionTime sql.NullTime //32 байта. Время последнего действия в сессии
	EndTime        sql.NullTime //32 байта. Время окончания сессии
	StartTime      time.Time    //24 байта. Время начала сессии
	Login          string       //16 байт — 8 байт для указателя и 8 байт для длины. Имя пользователя
	Client         string       //16 байт — 8 байт для указателя и 8 байт для длины. Имя клиентского приложения, полученное при авторизации
}

// Claims - это данные сессии, которые использует клиент для подписи своих запросов
type Claims struct {
	jwt.StandardClaims        //88 байт
	Login              string `json:"login"`     //16 байт — 8 байт для указателя и 8 байт для длины. Логин
	Client             string `json:"clientid"`  //16 байт — 8 байт для указателя и 8 байт для длины. Клиент
	SessionID          int    `json:"sessionid"` //8 байт. Идентификатор сессии

}
