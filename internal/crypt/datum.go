package crypt

import (
	"encoding/json"
	"fmt"
)

// CryptData зашифровывает структуру любого типа
func CryptData(data any, key []byte) ([]byte, error) {
	// Приводим любую структуру к json
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("CryptData: %v", err)
	}
	// Шифруем json, полученный из структуры
	encrypted, err := Encrypt(dataJson, key)
	if err != nil {
		return nil, fmt.Errorf("CryptData 2: %v", err)
	}
	return encrypted, nil
}
