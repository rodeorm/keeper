package cfg

import (
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/repo"
)

func ConfigurateServer(configFile string) (*Server, error) {
	queue := core.NewQueue(3)

	builder := &ServerBuilder{}
	srv := builder.SetConfig(configFile).
		Build()
	srv.MessageQueue = queue

	//Получаем хранилище. Пока одно.
	storage, err := repo.GetPostgresStorage(srv.DBProd)
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
