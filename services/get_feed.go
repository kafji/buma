package services

import "context"

type UserFeed struct {
	Items []UserFeedItem `json:"items"`
}

type UserFeedItem struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	SourceName string `json:"source_name"`
	SourceURL  string `json:"source_url"`
}

type GetUserFeed interface {
	GetUserFeed(ctx context.Context, userID int) UserFeed
}

func GetFeed(ctx context.Context, guf GetUserFeed, userID int) UserFeed {
	return guf.GetUserFeed(ctx, userID)
}
