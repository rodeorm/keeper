package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Функция для шифрования данных
func Encrypt(plainDatum []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plainDatum))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plainDatum)
	return ciphertext, nil
}

// Функция для расшифровки данных
func Decrypt(cipherDatum []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(cipherDatum) < aes.BlockSize {
		return nil, fmt.Errorf("слишком короткие зашифрованные данные")
	}
	iv := cipherDatum[:aes.BlockSize]
	cipherDatum = cipherDatum[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherDatum, cipherDatum)
	return cipherDatum, nil
}
