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

func ReadBinary(ctxBg context.Context, token string, b *core.Binary, c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	ctx = meta.AddTokenToCtx(ctx, token)

	res, err := c.ReadBinary(ctx, &proto.ReadBinaryRequest{Binary: &proto.Binary{Id: int32(b.ID)}}, grpc.Header(&header), grpc.Trailer(&trailer))
	if status.Code(err) != codes.OK {
		return err
	}

	b.Name = res.Binary.Name
	b.Value = res.Binary.Value

	return nil
}
