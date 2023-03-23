package main

import (
	"context"

	"golang.org/x/exp/slog"
	"kafji.net/buma/app"
	"kafji.net/buma/config"
	"kafji.net/buma/feed"
	"kafji.net/buma/services"
)

func main() {
	ctx := context.Background()

	cfg := config.ReadConfig()

	db := app.Database(ctx, &cfg)
	defer db.Close()

	slog.Info("starting fetch app")

	services.FetchFeeds(ctx, &db, services.FetchFeedFunc(feed.FetchFeeds), &db)
}
