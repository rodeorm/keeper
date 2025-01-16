package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCouple
func (g *grpcServer) CreateCouple(ctx context.Context, cr *proto.CreateCoupleRequest) (*proto.CreateCoupleResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}
	couple := &core.Couple{
		Source:   cr.Couple.Source,
		Login:    cr.Couple.Login,
		Password: cr.Couple.Password,
		Meta:     cr.Couple.Meta,
	}

	err = g.cfg.CoupleStorager.AddCoupleByUser(ctx, couple, usr)
	if err != nil {
		return nil, status.Error(codes.Aborted, `не удалось создать`)
	}

	return &proto.CreateCoupleResponse{}, status.Error(codes.OK, `удалось сохранить пару логин-пароль в БД`)
}
