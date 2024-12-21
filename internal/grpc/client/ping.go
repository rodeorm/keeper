package client

import (
	"context"

	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Ping пинг для проверки соединения
func Ping(c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx := context.TODO()
	req := proto.PingRequest{}

	_, err := c.Ping(ctx, &req, grpc.Header(&header), grpc.Trailer(&trailer))

	if err != nil {
		return err
	}

	return nil
}
