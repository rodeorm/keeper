package logger

import "go.uber.org/zap"

// Error - логирует ошибки как в консоль, так и во внешнее хранилище (TODO)
func Error(place, message, err string) {
	Log.Error(place, zap.String(message, err))
}
