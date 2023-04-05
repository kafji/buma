package main

import (
	"context"

	"golang.org/x/exp/slog"
	"kafji.net/buma/app"
	"kafji.net/buma/config"
	"kafji.net/buma/feed"
	fetchfeeds "kafji.net/buma/services/fetch_feeds"
)

func main() {
	ctx := context.Background()

	cfg := config.ReadConfig()

	db := app.Database(ctx, &cfg)
	defer db.Close()

	slog.Info("starting fetch app")

	fetchfeeds.FetchFeeds(ctx, &db, fetchfeeds.FetchFeedFunc(feed.FetchFeeds), &db)
}
