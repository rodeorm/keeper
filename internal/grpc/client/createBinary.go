package client

import (
	"context"
	"time"

	"github.com/rodeorm/keeper/internal/core"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func CreateBinary(ctxBg context.Context, token string, b core.Binary, c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	ctx = meta.AddTokenToCtx(ctx, token)

	req := &proto.CreateBinaryRequest{
		Binary: &proto.Binary{
			Name:  b.Name,
			Meta:  b.Meta,
			Value: b.Value,
		},
	}

	_, err := c.CreateBinary(ctx, req, grpc.Header(&header), grpc.Header(&trailer))

	if status.Code(err) != codes.OK {
		return err
	}

	return nil
}
