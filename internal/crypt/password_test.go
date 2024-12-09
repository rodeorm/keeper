package crypt

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// Тест для функции HashPassword
func TestHashPassword(t *testing.T) {
	password := "MySecurePassword123!"

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err, "Ошибка при хэшировании пароля")
	require.NotEmpty(t, hashedPassword, "Хэш пароля не должен быть пустым")

	// Дополнительно проверяем, что хеш действительно является хешем
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	require.NoError(t, err, "Пароль не совпадает с хэшем")
}

// Тест для функции CheckPasswordByHash
func TestCheckPasswordByHash(t *testing.T) {
	password := "MySecurePassword123!"
	hash, err := HashPassword(password)
	require.NoError(t, err, "Ошибка при хэшировании пароля")

	// Проверяем, что CheckPasswordByHash возвращает true для правильного пароля
	isValid := CheckPasswordByHash(password, hash)
	require.True(t, isValid, "Пароль должен совпадать с хэшем")

	// Проверяем, что CheckPasswordByHash возвращает false для неправильного пароля
	isValid = CheckPasswordByHash("WrongPassword", hash)
	require.False(t, isValid, "Пароль не должен совпадать с хэшем")
}
