package server

import (
	"context"
	"log"

	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadAllCards возвращает все карты конкретного пользователя
func (g *grpcServer) ReadAllCards(ctx context.Context, cr *proto.ReadAllCardsRequest) (*proto.ReadAllCardsResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}

	cards, err := g.cfg.CardStorager.SelectAllCardsByUser(ctx, usr)
	if err != nil {
		return nil, status.Error(codes.DataLoss, `не удалось получить данные`)
	}

	resp := proto.ReadAllCardsResponse{}
	resp.Cards = make([]*proto.Card, 0)
	for _, v := range cards {
		c := proto.Card{CardNumber: v.CardNumber,
			OwnerName: v.OwnerName,
			ExpMonth:  int32(v.ExpMonth),
			ExpYear:   int32(v.ExpYear),
			CVC:       int32(v.CVC),
			Meta:      v.Meta,
			Id:        int32(v.ID),
		}
		resp.Cards = append(resp.Cards, &c)
	}

	log.Printf("вернули пользователю %s количество карт %d", usr.Login, len(resp.Cards))

	return &resp, status.Error(codes.OK, `аутентификация и авторизация пройдены`)
}
