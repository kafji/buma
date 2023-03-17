package services

import "context"

type GetUserIDByToken interface {
	GetUserIDByToken(ctx context.Context, token string) (id int, found bool)
}

func Authenticate(ctx context.Context, guibt GetUserIDByToken, token string) (userID int, ok bool) {
	return guibt.GetUserIDByToken(ctx, token)
}
