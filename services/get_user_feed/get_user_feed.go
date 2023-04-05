package getuserfeed

import (
	"context"

	"kafji.net/buma/models"
)

type QueryUserFeed interface {
	QueryUserFeed(ctx context.Context, userID int) models.UserFeed
}

func GetUserFeed(ctx context.Context, quf QueryUserFeed, userID int) models.UserFeed {
	return quf.QueryUserFeed(ctx, userID)
}
