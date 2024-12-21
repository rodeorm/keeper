package client

import (
	"context"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Verify направляет на проверку OTP. Возвращает прошел или не прошел авторизацию и токен авторизации
func Verify(u *core.User, ctx context.Context, c proto.KeeperServiceClient) (bool, string) {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	req := proto.VerifyRequest{
		Login: u.Login,
		OTP:   u.OTP,
	}
	resp, err := c.Verify(ctx, &req, grpc.Header(&header), grpc.Trailer(&trailer))

	if status.Code(err) != codes.OK {
		return false, err.Error()
	}

	token := meta.GetTokenFromMeta(header)

	return resp.Verified, token
}
