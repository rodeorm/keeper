package logger

import "go.uber.org/zap"

// Info - логирует информационные сообщения как в консоль, так и во внешнее хранилище (TODO)
func Info(place, message, additional string) {
	Log.Info(place, zap.String(message, additional))
}
