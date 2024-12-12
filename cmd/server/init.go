package main

import "flag"

var (
	c *string // путь к конфигурационному файлу
)

func init() {
	// Парсим флаги
	flag.Parse()
	c = flag.String("c", "", "CONFIG_FILE")
}
