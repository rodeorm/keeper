package meta

import (
	"context"

	"github.com/rodeorm/keeper/internal/crypt"
	"google.golang.org/grpc/metadata"
)

// GetLoginFromCtx получает идентфикатор пользователя из мета в контексте
func GetLoginFromCtx(ctx *context.Context, jwtKey string) (string, error) {

	var token string

	md, ok := metadata.FromIncomingContext(*ctx)
	if ok {
		values := md.Get("token")
		if len(values) > 0 {
			// ключ содержит слайс строк, получаем первую строку
			token = values[0]
		}
	}

	claims, err := crypt.DecodeSession(token, jwtKey)
	if err != nil {
		return "", err
	}

	return claims.Login, nil
}

// PutUserKeyToMD помещает идентификатор пользователя в мету
func PutLoginToMD(login, jwtKey string, sessionID, tokenLiveTime int) (metadata.MD, error) {
	val, err := crypt.CodeSession(login, sessionID, jwtKey, tokenLiveTime)
	if err != nil {
		return nil, err
	}

	return metadata.Pairs("token", val), nil
}