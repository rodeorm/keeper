package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

// Функция для шифрования данных
func Encrypt(plainDatum []byte, key []byte) ([]byte, error) {
	plainDatum = pad(plainDatum, aes.BlockSize) // Добавляем дополнение
	// Проверка длины ключа
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("допустимые длины ключа: 16, 24 или 32 байта")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Подготовка к шифрованию
	ciphertext := make([]byte, aes.BlockSize+len(plainDatum))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Шифрование
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plainDatum)

	return ciphertext, nil
}

// pad добавляет дополнение PKCS#7 к данным
func pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// unpad удаляет дополнение PKCS#7 из данных
func unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("данные пусты")
	}
	padding := data[length-1]
	if int(padding) > length {
		return nil, errors.New("некорректное дополнение")
	}
	return data[:length-int(padding)], nil
}

// Decrypt расшифровывает данные с использованием AES в режиме CBC
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	// Проверка длины ключа
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, errors.New("допустимые длины ключа: 16, 24 или 32 байта")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Проверка длины зашифрованных данных
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("недостаточная длина зашифрованных данных")
	}

	// Извлечение вектора инициализации (IV)
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Расшифрование
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Удаление дополнения
	return unpad(ciphertext)
}
