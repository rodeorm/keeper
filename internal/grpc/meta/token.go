package meta

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// Добавляет jwt-токен в контекст
func AddTokenToCtx(ctxBg context.Context, token string) context.Context {
	md := metadata.Pairs("token", token)
	return metadata.NewOutgoingContext(ctxBg, md)

}

func GetTokenFromMeta(md metadata.MD) string {
	var token string
	if md != nil {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}
	return token
}
