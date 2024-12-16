package crypt

import (
	"encoding/json"
	"fmt"
)

// CryptData зашифровывает структуру любого типа
func CryptData(data any, password []byte) ([]byte, error) {
	// Приводим любую структуру к json
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("не получилось")
	}
	// Шифруем json, полученный из структуры
	encrypted, err := Encrypt(dataJson, password)
	if err != nil {
		return nil, fmt.Errorf("не получилось")
	}
	return encrypted, nil
}
