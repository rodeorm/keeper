package client

import (
	"context"
	"log"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

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
	log.Println("получен ответ от grpc сервера для метода Reg", resp, err)
	return err
}
