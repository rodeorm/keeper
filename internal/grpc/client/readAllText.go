package client

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
)

func ReadAllTexts(ctxBg context.Context, token string, c proto.KeeperServiceClient) ([]*proto.Text, error) {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	ctx = meta.AddTokenToCtx(ctx, token)
	resp, err := c.ReadAllTexts(ctx, &proto.ReadAllTextsRequest{}, grpc.Header(&header), grpc.Trailer(&trailer))

	if status.Code(err) != codes.OK {
		return nil, err
	}

	return resp.Texts, nil
}
