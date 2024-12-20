package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Auth осуществляет авторизацию и аутентификацию пользователя
func (g *grpcServer) Auth(ctx context.Context, po *proto.AuthRequest) (*proto.AuthResponse, error) {
	usr := core.User{Login: po.Login, Password: po.Password}
	var resp proto.AuthResponse

	authenticated := g.cfg.UserStorager.AuthUser(ctx, &usr)
	if !authenticated {
		return nil, status.Error(codes.Unauthenticated, `пользователь не прошел базовую аутентификацию`)
	}

	return &resp, nil
}
