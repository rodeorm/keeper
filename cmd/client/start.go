package main

import (
	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
)

// Теоретически у клиентского приложения могла бы быть своя БД, которая будет эффективно синхронизироваться с сервером,
// Для упрощения - не будет локального хранилища на клиенте
// Аналогично - не будет локально пароля для шифрования данных. Вместо этого - шифрование на сервере общее
// Валидация данных осуществляется на клиенте
func newGRPCClient() proto.KeeperServiceClient {
	config, err := cfg.GetClientConfigFromFile()
	if err != nil {
		panic(err)
	}
	// устанавливаем соединение с сервером
	conn, err := grpc.NewClient(config.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	c := proto.NewKeeperServiceClient(conn)

	/*
		ctx := context.Background()

		u := core.User{Login: "firstUser", Password: "Thisispassword", Name: "Alexander", Phone: "792592504011", Email: "ilyin-a-l@ya.ru"}
		client.RegUser(&u, ctx, c)
	*/
	return c
}
