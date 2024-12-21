package client

import (
	"context"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthUser первично авторизует пользователя (проверяет соответствие логину и паролю, отправляет OTP)
func AuthUser(u *core.User, ctx context.Context, c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	req := proto.AuthRequest{
		Login:    u.Login,
		Password: u.Password,
	}

	_, err := c.Auth(ctx, &req, grpc.Header(&header), grpc.Trailer(&trailer))

	if status.Code(err) != codes.OK {
		return err
	}

	return nil
}
