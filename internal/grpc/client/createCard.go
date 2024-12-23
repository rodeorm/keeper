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

func CreateCard(ctxBg context.Context, token string, crd core.Card, c proto.KeeperServiceClient) error {
	var header, trailer metadata.MD
	ctx, cancel := context.WithTimeout(ctxBg, 10*time.Second)
	defer cancel()
	ctx = meta.AddTokenToCtx(ctx, token)

	_, err := c.CreateCard(ctx,
		&proto.CreateCardRequest{Card: &proto.Card{CardNumber: crd.CardNumber,
			OwnerName: crd.OwnerName,
			ExpMonth:  int32(crd.ExpMonth),
			ExpYear:   int32(crd.ExpYear),
			Meta:      crd.Meta}},
		grpc.Header(&header),
		grpc.Trailer(&trailer))
	if status.Code(err) != codes.OK {
		return err
	}

	return nil
}
