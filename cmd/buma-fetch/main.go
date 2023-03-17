package main

import (
	"context"

	"github.com/rs/zerolog/log"
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

	log.Info().Msg("app: starting fetch app")

	services.FetchFeeds(ctx, &db, services.FetchFeedFunc(feed.FetchFeeds), &db)
}
