package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *grpcServer) ReadBinary(ctx context.Context, cr *proto.ReadBinaryRequest) (*proto.ReadBinaryResponse, error) {
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}

	binary := &core.Binary{ID: int(cr.Binary.Id)}

	err = g.cfg.BinaryStorager.SelectBinaryByUser(ctx, binary, usr)
	if err != nil {
		return nil, status.Error(codes.DataLoss, `не удалось получить данные`)
	}

	resp := proto.ReadBinaryResponse{Binary: &proto.Binary{Value: binary.Value, Name: binary.Name}}

	return &resp, status.Error(codes.OK, `Файл отправлен`)
}
