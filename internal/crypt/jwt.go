package crypt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/logger"
)

// CodeSession кодирует сессию в строку c использованием JWT
// Для этого этой функции надо передать  данные логина, идентификатор сессии, ключ для кодирования, время жизни токена
func CodeSession(login string, sessionId int, jwtKey string, tokenLiveTime int) (string, error) {
	key := []byte(jwtKey)
	c := createClaims(login, sessionId, tokenLiveTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// DecodeSession декодирует данные сессии из строки
// Для этого этой функции надо передать саму строку и ключ, использованный для кодирования
func DecodeSession(tokenStr, jwtKey string) (*core.Claims, error) {
	key := []byte(jwtKey)
	claims := &core.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		logger.Error("DecodeSession", "ошибка при попытке распарсить строку с данными токена", err.Error())
		return nil, err
	}
	if !tkn.Valid {
		logger.Info("DecodeSession", "попытке предоставить невалидный токен", "токен не прошел проверку")
		return nil, fmt.Errorf("невалидный токен")
	}
	return claims, nil
}

// CreateClaims создает claims для jwt на основании логина, ид сессии и времени жизни
func createClaims(login string, sessionID, tokenLiveTime int) *core.Claims {
	// Срок истечения валидности токена
	expirationTime := time.Now().Add(time.Duration(tokenLiveTime) * time.Minute)

	return &core.Claims{
		SessionID: sessionID,
		Login:     login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // в JWT время жизни в Unix миллисекундах
		},
	}
}
