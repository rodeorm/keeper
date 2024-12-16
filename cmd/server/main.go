package main

import (
	"context"
	"log"
	"sync"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/server"
)

func main() {

	srv, err := cfg.ConfigurateServer(*c) // Остальные параметры забираем из конфигурационного файла
	if err != nil {
		panic(err) // Если не получилось получить настройки, то поникуем в ужасе
	}

	u := core.User{Login: "user", Password: "12345", Email: "it-serv@it.ru"}
	err = srv.UserStorager.RegUser(context.TODO(), &u)
	log.Println(err)
	//srv.UserStorager.AuthUser(context.TODO(), &u)

	// Через этот канал горутины узнают, что надо закрываться для изящного завершения работы
	exit := make(chan struct{})
	// Через эту группу мы синхронизируем горутины
	var wg sync.WaitGroup

	//wg.Add(srv.SenderQuantity + 1) // Количество отправителей + одна горутина с gRPC сервером
	wg.Add(1)
	go server.ServerStart(srv, &wg, exit) // Асинхронно запускаем grpc сервер
	//go sender.SenderStart(srv, &wg, exit) // Асинхронно запускаем сендеры
	wg.Wait()
}
