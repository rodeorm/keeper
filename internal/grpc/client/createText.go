package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
)

func CreateText(ctxBg context.Context, token string, t core.Text, c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	ctx = meta.AddTokenToCtx(ctx, token)

	_, err := c.CreateText(ctx,
		&proto.CreateTextRequest{Text: &proto.Text{
			Text: t.Value,
			Meta: t.Meta,
		}},
		grpc.Header(&header),
		grpc.Trailer(&trailer))
	if status.Code(err) != codes.OK {
		return err
	}

	return nil
}
