// Package crypt реализует шифрование-дешифрование данных
package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

// Пароль пользователя - это ключ для шифрования его данных
// Сам пароль хранится в БД в зашифрованном виде

// HashPassword хэширует пароль
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordByHash проверяет пароль на соответствие хэшу в БД
//
// Если возвращает значение "истина", значит пароль соответствует сохраненному в БД
func CheckPasswordByHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
