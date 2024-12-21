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

// RegUser регистрирует пользователя
func RegUser(u *core.User, ctx context.Context, c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	req := proto.RegRequest{
		User: &proto.User{
			Login:    u.Login,
			Password: u.Password,
			Name:     u.Name,
			Email:    u.Email,
			Phone:    u.Phone,
		},
	}

	resp, err := c.Reg(ctx, &req, grpc.Header(&header), grpc.Trailer(&trailer))

	if status.Code(err) != codes.OK {

		return err
	}
	u.ID = int(resp.Id)
	return nil
}
