package server

import (
	"context"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateText
func (g *grpcServer) CreateText(ctx context.Context, cr *proto.CreateTextRequest) (*proto.CreateTextResponse, error) {
	//Сначала смотрим, кто к нам обращается по мете
	usr, err := meta.GetUserIdentity(ctx, g.cfg.CryptKey)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, `отказано в доступе`)
	}
	text := &core.Text{Value: cr.Text.Text, Meta: cr.Text.Meta}

	err = g.cfg.TextStorager.AddTextByUser(ctx, text, usr)
	if err != nil {
		return nil, status.Error(codes.Aborted, `не удалось создать`)
	}

	return &proto.CreateTextResponse{}, status.Error(codes.OK, `удалось сохранить текст в БД`)
}
