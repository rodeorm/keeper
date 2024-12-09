package core

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// SessionStorager абстрагирует хранилище данных сессии
type SessionStorager interface {
	StartSession(*User) (*Session, error) //Начинает новую сессию
	UpdateSession(*Session) error         //Обновляет данные сессии
	EndSession(*Session) error            //Закрывает сессию
}

// Session - сессия пользователя, которая хранится централизованно и позволяет завершить любую клиентскую сессию через сервер
type Session struct {
	Login          string    //Имя пользователя
	Client         string    //Имя клиентского приложения, полученное при авторизации
	StartTime      time.Time //Время начала сессии
	LastActionTime time.Time //Время последнего действия в сессии
	EndTime        time.Time //Время окончания сессии
}

// Claims - это данные сессии, которые использует клиент для подписи своих запросов
type Claims struct {
	Login     string `json:"login"`
	Client    string `json:"clientid"`
	SessionID int    `json:"sessionid"`
	jwt.StandardClaims
}
