package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateCard создает карту
func (g *grpcServer) CreateCard(ctx context.Context, cr *proto.CreateCardRequest) (*proto.CreateCardResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}
	card := &core.Card{CardNumber: cr.Card.CardNumber,
		OwnerName: cr.Card.OwnerName,
		ExpMonth:  int(cr.Card.ExpMonth),
		ExpYear:   int(cr.Card.ExpYear),
		Meta:      cr.Card.Meta.Value}

	err = g.cfg.CardStorager.AddCardByUser(ctx, card, usr)
	if err != nil {
		return nil, status.Error(codes.Aborted, `не удалось создать`)
	}

	resp := proto.CreateCardResponse{}

	return &resp, status.Error(codes.OK, `аутентификация и авторизация пройдены`)
}
