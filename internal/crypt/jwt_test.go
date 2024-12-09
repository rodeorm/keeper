package crypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Тест для функции CodeSession
func TestCodeSession(t *testing.T) {
	login := "testUser"
	sessionID := 12345
	jwtKey := "superSecretKey" // Преобразуем ключ в []byte
	tokenLiveTime := 30        // минут

	token, err := CodeSession(login, sessionID, jwtKey, tokenLiveTime)
	require.NoError(t, err, "Ошибка при кодировании сессии")
	require.NotEmpty(t, token, "Токен не должен быть пустым")
}

// Тест для функции DecodeSession
func TestDecodeSession(t *testing.T) {
	login := "testUser"
	sessionID := 12345
	jwtKey := "superSecretKey" // Преобразуем ключ в []byte
	tokenLiveTime := 30        // минут

	// Кодируем сессию для теста
	token, err := CodeSession(login, sessionID, jwtKey, tokenLiveTime)
	require.NoError(t, err, "Ошибка при кодировании сессии")

	// Декодируем сессию
	claims, err := DecodeSession(token, jwtKey)
	require.NoError(t, err, "Ошибка при декодировании сессии")
	require.NotNil(t, claims, "Claims не должны быть nil")
	require.Equal(t, login, claims.Login, "Логин не совпадает")
	require.Equal(t, sessionID, claims.SessionID, "ID сессии не совпадает")

	// Проверяем на невалидный токен
	invalidToken := "invalid.token.string"
	claims, err = DecodeSession(invalidToken, jwtKey)
	require.Error(t, err, "Ошибка ожидалась при декодировании невалидного токена")
	require.Nil(t, claims, "Claims должны быть nil для невалидного токена")
}
