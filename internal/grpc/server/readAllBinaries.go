package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadAllBinaries
func (g *grpcServer) ReadAllBinaries(ctx context.Context, cr *proto.ReadAllBinariesRequest) (*proto.ReadAllBinariesResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}

	binaries, err := g.cfg.BinaryStorager.SelectAllBinariesByUser(ctx, usr)
	if err != nil {
		return nil, status.Error(codes.DataLoss, `не удалось получить данные`)
	}
	resp := proto.ReadAllBinariesResponse{}
	resp.Binaries = make([]*proto.Binary, 0)
	for _, v := range binaries {
		b := proto.Binary{Name: v.Name,
			Meta: v.Meta,
			Id:   int32(v.ID),
			// Value: v.Value, // TODO: По уму сам бинарник не надо возвращать. Его лучше получить отдельно только по запросу
		}
		resp.Binaries = append(resp.Binaries, &b)
	}

	return &resp, status.Error(codes.OK, `аутентификация и авторизация пройдены`)
}
