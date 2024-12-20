package client

import (
	"context"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthUser первично авторизует пользователя
func AuthUser(u *core.User, ctx context.Context, c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	req := proto.AuthRequest{
		Login:    u.Login,
		Password: u.Password,
	}

	resp, err := c.Auth(ctx, &req, grpc.Header(&header), grpc.Trailer(&trailer))
	// Если вернулся nil, то значит ошибка
	if resp == nil {
		return err
	}

	return nil
}
