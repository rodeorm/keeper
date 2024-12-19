package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/proto"
)

// Reg регистрирует нового пользователя
func (g *grpcServer) Reg(ctx context.Context, po *proto.RegRequest) (*proto.RegResponse, error) {
	var resp proto.RegResponse

	usr := &core.User{Login: po.User.Login, Password: po.User.Password, Phone: po.User.Phone, Email: po.User.Email, Name: po.User.Name}
	err := g.cfg.UserStorager.RegUser(ctx, usr)

	//TODO: обработать разные причины, почему не получилось зарегистрировать пользователя
	if err != nil {
		return nil, status.Error(codes.Unavailable, `ошибка при попытке зарегистрировать пользователя`)
	}
	return &resp, status.Error(codes.OK, `успешно зарегистрировали пользователя`)
}
