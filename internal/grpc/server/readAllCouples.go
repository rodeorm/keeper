package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadAllCouples
func (g *grpcServer) ReadAllCouples(ctx context.Context, cr *proto.ReadAllCouplesRequest) (*proto.ReadAllCouplesResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}

	cpls, err := g.cfg.CoupleStorager.SelectAllCouplesByUser(ctx, usr)
	if err != nil {
		return nil, status.Error(codes.DataLoss, `не удалось получить данные`)
	}

	resp := proto.ReadAllCouplesResponse{}
	resp.Couples = make([]*proto.Couple, 0)
	for _, v := range cpls {
		c := proto.Couple{
			Source:   v.Source,
			Login:    v.Login,
			Password: v.Password,
			Meta:     v.Meta,
			Id:       int32(v.ID),
		}

		resp.Couples = append(resp.Couples, &c)
	}

	return &resp, status.Error(codes.OK, `данные пар логин-пароль`)
}
