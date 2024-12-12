package cfg

import (
	"fmt"
	"os"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/logger"
	"github.com/rodeorm/keeper/internal/sender"
	"gopkg.in/yaml.v2"
)

// Server - сервер
type Server struct {
	core.UserStorager
	core.BinaryStorager
	core.CardStorager
	core.MessageStorager
	core.CoupleStorager
	core.SessionStorager
	core.TextStorager

	ServerConfig
	MessageQueue *sender.Queue
}

// ServerBuilder абстракция для создания сервера
type ServerBuilder struct {
	server Server
}

// Build возвращает сконфигурированный сервер
func (s ServerBuilder) Build() Server {
	return s.server
}

// SetSessionStorage определяет реализацию хранилища бинарных файлов для сервера
func (s ServerBuilder) SetTextStorage(m core.TextStorager) ServerBuilder {
	s.server.TextStorager = m
	return s
}

// SetSessionStorage определяет реализацию хранилища бинарных файлов для сервера
func (s ServerBuilder) SetSessionStorage(m core.SessionStorager) ServerBuilder {
	s.server.SessionStorager = m
	return s
}

// SetCoupleStorage определяет реализацию хранилища бинарных файлов для сервера
func (s ServerBuilder) SetCoupleStorage(m core.CoupleStorager) ServerBuilder {
	s.server.CoupleStorager = m
	return s
}

// SetMessageStorage определяет реализацию хранилища бинарных файлов для сервера
func (s ServerBuilder) SetMessageStorage(m core.MessageStorager) ServerBuilder {
	s.server.MessageStorager = m
	return s
}

// SetCardStorage определяет реализацию хранилища бинарных файлов для сервера
func (s ServerBuilder) SetCardStorage(c core.CardStorager) ServerBuilder {
	s.server.CardStorager = c
	return s
}

// SetBinaryStorage определяет реализацию хранилища бинарных файлов для сервера
func (s ServerBuilder) SetBinaryStorage(b core.BinaryStorager) ServerBuilder {
	s.server.BinaryStorager = b
	return s
}

// SetConfig заполняет конфигурацию данными из переменных окружения и флагов
func (s ServerBuilder) SetConfig(name string) ServerBuilder {
	cfg, err := GetConfigFromFile(name)
	if err != nil {
		return s
	}
	s.server.ServerConfig = *cfg
	return s
}

// SetUserStorager определяет реализацию хранилища пользователей для сервера
func (s ServerBuilder) SetUserStorage(u core.UserStorager) ServerBuilder {
	s.server.UserStorager = u
	return s
}

// GetConfigFromFile возвращает конфигурацию из конфигурационного файла
func GetConfigFromFile(name string) (*ServerConfig, error) {
	if name == "" {
		name = "config.yml"
	}
	file, err := os.Open(name)

	if err != nil {
		logger.Error("SetConfigFromFile", fmt.Sprintf("открытие файла %s", name), err.Error())
		return nil, err
	}
	defer file.Close()

	cfg := &ServerConfig{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		logger.Error("SetConfigFromFile", fmt.Sprintf("открытие файла %s", name), err.Error())
		return nil, err
	}

	return cfg, nil
}
