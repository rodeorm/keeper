package main

import (
	"context"
	"sync"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/server"
	"github.com/rodeorm/keeper/internal/msg/filler"
	"github.com/rodeorm/keeper/internal/msg/sender"
)

func main() {

	srv, err := cfg.ConfigurateServer(*c) // Остальные параметры забираем из конфигурационного файла
	if err != nil {
		panic(err) // Если не получилось получить настройки, то в ужасе паникуем
	}

	ctx := context.TODO()

	usr1 := &core.User{Login: "user1", Password: "12345", Email: "ilyin-a-l@yandex.ru"}
	usr2 := &core.User{Login: "user2", Password: "12345", Email: "ilyin-a-l@ya.ru"}
	usr3 := &core.User{Login: "user3", Password: "12345", Email: "ilyin-a-l@ya.ru"}
	usr4 := &core.User{Login: "user4", Password: "12345", Email: "ilyin-a-l@ya.ru"}

	srv.UserStorager.RegUser(ctx, usr1)
	srv.UserStorager.RegUser(ctx, usr2)
	srv.UserStorager.RegUser(ctx, usr3)
	srv.UserStorager.RegUser(ctx, usr4)

	// Через этот канал горутины узнают, что надо закрываться для изящного завершения работы
	exit := make(chan struct{})
	// Через эту группу мы синхронизируем горутины
	var wg sync.WaitGroup

	wg.Add(srv.SenderQuantity + 2)  // Количество отправителей + одна горутина наполнителя очереди + одна горутина с gRPC сервером
	go server.Start(srv, &wg, exit) // Асинхронно запускаем grpc сервер
	go filler.Start(srv, &wg, exit) // Асинхронно наполняем очередь сообщений с OTP к отправке
	go sender.Start(srv, &wg, exit) // Асинхронно запускаем сендеры
	wg.Wait()
}
