package server

import (
	"context"
	"log"

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

	log.Println("сервер получил запрос на получение бинарников пользователя", usr)

	binaries, err := g.cfg.BinaryStorager.SelectBinaryByUser(ctx, usr)
	if err != nil {
		return nil, status.Error(codes.DataLoss, `не удалось получить данные`)
	}
	resp := proto.ReadAllBinariesResponse{}
	resp.Binaries = make([]*proto.Binary, 0)
	for _, v := range binaries {

		b := proto.Binary{Name: v.Name,
			Meta:  v.Meta,
			Id:    int32(v.ID),
			Value: v.Value, // По уму сам бинарник не надо возвращать. Его лучше получить отдельно только по запросу, но времени у меня нет.
		}
		log.Println("Получили бинарники", v.Name, v.Meta)
		resp.Binaries = append(resp.Binaries, &b)
	}

	return &resp, status.Error(codes.OK, `аутентификация и авторизация пройдены`)
}
