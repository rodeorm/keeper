package crypt

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

// ReturnOTP возвращает OTP
func ReturnOTP(n int) (string, error) {
	if n <= 0 {
		return "", fmt.Errorf("некорректное значение ключа %v", n)
	}

	n = 3 //!!! лень вводить больше при тестах

	b := make([]rune, n)
	for i := range b {
		// Генерируем случайное число
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		b[i] = letters[index.Int64()]
	}
	return string(b), nil
	// return "a", nil // Временно для тестов
}
