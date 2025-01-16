package cfg

import (
	"os"
	"testing"
)

// TestGetClientConfigFromFile тестирует функцию GetClientConfigFromFile
func TestGetClientConfigFromFile(t *testing.T) {
	// Определяем тестовый файл конфигурации
	configData := []byte(`
SERVER_ADDRESS: 8080
FILE_PATH: localhost
`)
	name := "config.yml"

	// Создаем временный файл с конфигурацией
	if err := os.WriteFile(name, configData, 0644); err != nil {
		t.Fatalf("не удалось создать файл %s: %v", name, err)
	}
	defer os.Remove(name) // Удаляем файл после завершения теста

	// Вызываем функцию для получения конфигурации
	cfg, err := GetClientConfigFromFile()
	if err != nil {
		t.Fatalf("ошибка при получении конфигурации: %v", err)
	}

	// Проверяем значения полей конфигурации
	expected := &ClientConfig{
		ServerAddress: "8080",
		FilePath:      "localhost",
	}

	if *cfg != *expected {
		t.Errorf("ожидалось %#v, но получено %#v", expected, cfg)
	}
}
