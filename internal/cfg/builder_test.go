package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerBuilder(t *testing.T) {
	builder := ServerBuilder{}
	// Тест на установку конфигурации
	t.Run("SetConfig with valid config file", func(t *testing.T) {
		// Предположим, что у вас есть файл конфигурации test_config.yml
		builder.SetConfig("test_config.yml")
		assert.NotNil(t, builder.server.ServerConfig) // Проверяем, что конфигурация установлена
	})

	t.Run("SetConfig with invalid config file", func(t *testing.T) {
		builder.SetConfig("invalid_config.yml")
		assert.Empty(t, builder.server.ServerConfig) // Проверяем, что конфигурация не установлена
	})
}
