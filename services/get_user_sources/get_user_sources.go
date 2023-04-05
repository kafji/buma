package getusersources

import (
	"context"

	"kafji.net/buma/models"
)

type QueryUserSources interface {
	QueryUserSources(ctx context.Context, userID int) []models.UserSource
}

func GetUserSources(ctx context.Context, gufs QueryUserSources, userID int) []models.UserSource {
	return gufs.QueryUserSources(ctx, userID)
}
