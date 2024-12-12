package main

import (
	"sync"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/grpc/server"
	"github.com/rodeorm/keeper/internal/sender"
)

func main() {

	cfgSrv, err := cfg.ConfigurateServer(*c) // Остальные параметры забираем из конфигурационного файла
	if err != nil {
		panic(err) // Если не получилось настроить, то паникуем в ужасе
	}

	// через этот канал горутины узнают, что надо закрываться для изящного завершения работы
	exit := make(chan struct{})
	// через эту группу мы синхронизируем горутины
	var wg sync.WaitGroup

	wg.Add(cfgSrv.SenderQuantity + 1) // Количество отправителей + одна горутина с gRPC сервером

	go server.ServerStart(cfgSrv, &wg, exit) // Асинхронно запускаем grpc сервер
	go sender.SenderStart(cfgSrv, &wg, exit) // Асинхронно запускаем сендеры
	wg.Wait()                                // Ждем
}
