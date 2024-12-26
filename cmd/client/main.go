package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/cli"
	"github.com/rodeorm/keeper/internal/grpc/client"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
)

// Теоретически у клиентского приложения могла бы быть своя БД, которая могла бы эффективно синхронизироваться с сервером,
// Но для упрощения - не будет локального хранилища на клиенте.
// Аналогично - не будет локально пароля для шифрования данных. Вместо этого - шифрование на сервере общее
// Валидация данных осуществляется на клиенте
func main() {

	config, err := cfg.GetClientConfigFromFile()
	if err != nil {
		log.Println("Ошибка при попытке получить конфигурационный файл")
		os.Exit(1)
		//	panic(err)
	}
	maxMessageSize := 40 * 1024 * 1024

	// устанавливаем соединение с сервером
	conn, err := grpc.NewClient(config.ServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMessageSize)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(maxMessageSize)),
	)
	if err != nil {
		log.Println("Ошибка при попытке установить соединение с сервером")
		os.Exit(1)
	}

	grpc := proto.NewKeeperServiceClient(conn)
	err = client.Ping(grpc)
	if err != nil {
		log.Println("Ошибка при проверке соединения с сервером")
		os.Exit(1)
	}
	defer conn.Close()

	initModel := cli.InitialModel(grpc, config.FilePath)
	p := tea.NewProgram(initModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Println("Ошибка при попытке запустить программу")
		os.Exit(1)
	}
}
