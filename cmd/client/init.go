package main

import (
	"fmt"
	"runtime"
)

func init() {
	// Получаем информацию о системе
	os := runtime.GOOS
	arch := runtime.GOARCH
	version := runtime.Version()
	appVersion := "beta 1.0"

	// Выводим информацию в командную строку
	fmt.Printf("Операционная система: %s\n", os)
	fmt.Printf("Архитектура: %s\n", arch)
	fmt.Printf("Версия клиента: %s\n", appVersion)
	fmt.Printf("Версия Go: %s\n\n\n", version)
}
