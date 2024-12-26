package cfg

import (
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/repo"
)

// ConfigurateServer создает сервер на основании данных конфиг файла
func ConfigurateServer(configFile string) (*Server, error) {
	queue := core.NewQueue(3)

	builder := &ServerBuilder{}
	srv := builder.SetConfig(configFile).
		Build()
	srv.MessageQueue = queue

	storage, err := repo.GetPostgresStorage(srv.DBProd, srv.CryptKey)
	if err != nil {
		return nil, err
	}

	// Определяем реализацию хранилищ
	srv.BinaryStorager = storage
	srv.CardStorager = storage
	srv.CoupleStorager = storage
	srv.TextStorager = storage
	srv.UserStorager = storage
	srv.SessionStorager = storage
	srv.MessageStorager = storage

	return &srv, nil
}
