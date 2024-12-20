package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Verify верифицирует OTP
func (g *grpcServer) Verify(ctx context.Context, po *proto.VerifyRequest) (*proto.VerifyResponse, error) {
	usr := core.User{Login: po.Login, OTP: po.OTP}
	var resp proto.VerifyResponse

	resp.Verified = g.cfg.UserStorager.VerifyUserOTP(ctx, g.cfg.OneTimePasswordLiveTime, &usr)
	if !resp.Verified {
		return nil, status.Error(codes.PermissionDenied, `в доступе отказано`)
	}

	// После верификации по одноразовому паролю помещаем в мета jwt-токен с данными пользователя,
	// если такого не будет у запросов на данные - то в доступе будет отказано
	md, err := meta.PutLoginToMD(usr.Login, g.cfg.CryptKey, usr.ID, 1, g.cfg.TokenLiveTime)
	if err != nil {
		return nil, status.Error(codes.Unavailable, `недоступно`)
	}
	grpc.SetHeader(ctx, md) // Прикрепляем мету с jwt-токеном в header

	return &resp, status.Error(codes.OK, `аутентификация и авторизация пройдены`)
}
