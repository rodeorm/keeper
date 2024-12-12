package main

import (
	"sync"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/grpc/server"
	"github.com/rodeorm/keeper/internal/sender"
)

func main() {
	queue := sender.NewQueue(3)
	cfgSrv, err := cfg.ConfigurateServer(*c, queue) // Остальные параметры забираем из конфигурационного файла, если каких-то не хватает, то из переменных окружения
	if err != nil {
		panic(err) // Если не получилось настроить, то паникуем в ужасе
	}
	// через этот канал потоки узнают, что надо закрываться
	exit := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(cfgSrv.SenderQuantity + 1)

	go server.ServerStart(cfgSrv, &wg, exit) // Асинхронно запускаем grpc сервер
	for i := range cfgSrv.SenderQuantity {
		// Асинхронно запускаем email сендеры

		s := sender.NewSender(
			queue,
			cfgSrv.MessageStorager,
			i,
			cfgSrv.SMTPPort,
			cfgSrv.MessagePeriod,
			cfgSrv.SMTPServer,
			cfgSrv.SMTPLogin,
			cfgSrv.SMTPPass,
		)
		go s.StartSending(exit, &wg, 15) //TODO: Period в параметры
	}
	wg.Wait()

}
