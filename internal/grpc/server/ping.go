package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/grpc/proto"
)

// Ping осуществляет авторизацию и аутентификацию пользователя
func (g *grpcServer) Ping(ctx context.Context, po *proto.PingRequest) (*proto.PingResponse, error) {
	var resp proto.PingResponse

	return &resp, nil
}
