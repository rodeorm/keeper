package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadAllTexts
func (g *grpcServer) ReadAllTexts(ctx context.Context, cr *proto.ReadAllTextsRequest) (*proto.ReadAllTextsResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}

	cpls, err := g.cfg.TextStorager.SelectAllTextsByUser(ctx, usr)
	if err != nil {
		return nil, status.Error(codes.DataLoss, `не удалось получить данные`)
	}

	resp := proto.ReadAllTextsResponse{}
	resp.Texts = make([]*proto.Text, 0)
	for _, v := range cpls {
		c := proto.Text{
			Text: v.Value,
			Meta: v.Meta,
			Id:   int32(v.ID),
		}

		resp.Texts = append(resp.Texts, &c)
	}

	return &resp, status.Error(codes.OK, `данные текстов`)
}
