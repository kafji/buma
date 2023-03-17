package services

import "context"

type UserFeedSource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type GetUserFeedSources interface {
	GetUserFeedSources(ctx context.Context, userID int) []UserFeedSource
}

func GetSources(ctx context.Context, gufs GetUserFeedSources, userID int) []UserFeedSource {
	return gufs.GetUserFeedSources(ctx, userID)
}
