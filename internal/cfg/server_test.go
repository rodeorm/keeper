package cfg

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Тест для функции ConfigurateServer
func TestConfigurateServer(t *testing.T) {
	// Создаем временный файл конфигурации
	configData := []byte(`RUN_ADDRESS: localhost:8080
DB_PRODUCTIVE:  postgres://login:password@localhost:5432/keeper?sslmode=disable
SSL_SERTIFICATE_RELATIVE_PATH: certFile
SSL_SERTIFICATE_KEY_RELATIVE_PATH: certKey
DB_TEST: postgres://app:qqqQQQ123@localhost:5433/keeper?sslmode=disable
SMTP_SERVER: smtp.yandex.ru
SMTP_PORT: 465
SMTP_LOGIN: login@ya.ru
SMTP_PASSWORD: xwuunnlcoezfeimq
MALE_TEMPLATE: template.html
FROM: email@email
FILE_NAME: logo.jpg
PASSWORD_KEY: verydifficultandstrong
OTP_LIVE_TIME: 1
TOKEN_LIVE_TIME: 100
SENDER_QUANTITY: 2
MESSAGE_SEND_PERIOD: 1
QUEUE_FILL_PERIOD: 1
CRYPT_KEY: оченьнадежныйключ
 	`)
	configFile := "test_config.yml"
	if err := os.WriteFile(configFile, configData, 0644); err != nil {
		t.Fatalf("не удалось создать файл %s: %v", configFile, err)
	}
	defer os.Remove(configFile) // Удаляем файл после завершения теста

	// Вызываем тестируемую функцию
	srv, err := ConfigurateServer(configFile)
	assert.NoError(t, err) // Проверяем, что ошибка не возникла
	assert.NotNil(t, srv)  // Проверяем, что сервер был создан

}
