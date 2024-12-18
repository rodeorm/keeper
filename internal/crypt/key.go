package crypt

import "fmt"

//PadString приводит строку к нужной длинне для шифрования данных
func PadString(input string, length int) ([]byte, error) {
	if length != 16 && length != 24 && length != 32 {
		return nil, fmt.Errorf("допустимые длины: 16, 24 или 32 байта")
	}

	// Преобразуем строку в байты
	data := []byte(input)

	// Если длина строки уже равна нужной, возвращаем ее
	if len(data) == length {
		return data, nil
	}

	// Если длина строки больше нужной, обрезаем ее
	if len(data) > length {
		return data[:length], nil
	}

	// Если длина строки меньше нужной, дополняем нулями
	paddedData := make([]byte, length)
	copy(paddedData, data)
	return paddedData, nil
}
