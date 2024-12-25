package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBinary создает бинарный файл
func (g *grpcServer) CreateBinary(ctx context.Context, cr *proto.CreateBinaryRequest) (*proto.CreateBinaryResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}

	b := &core.Binary{
		Value: cr.Binary.Value,
		Name:  cr.Binary.Name,
		Meta:  cr.Binary.Meta,
	}

	err = g.cfg.BinaryStorager.AddBinaryByUser(ctx, b, usr)
	if err != nil {
		return nil, status.Error(codes.Aborted, `не удалось сохранить файл в БД`)
	}

	return &proto.CreateBinaryResponse{}, status.Error(codes.OK, `файл сохранен в БД`)
}
