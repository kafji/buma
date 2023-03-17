package main

import (
	"context"

	"github.com/rs/zerolog/log"
	"kafji.net/buma/app"
	"kafji.net/buma/config"
	"kafji.net/buma/server"
)

func main() {
	ctx := context.Background()

	cfg := config.ReadConfig()

	db := app.Database(ctx, &cfg)
	defer db.Close()

	log.Info().Msg("app: starting server app")

	err := server.StartServer(cfg.Server.Port, db)
	if err != nil {
		log.Panic().Err(err).Msg("app: server error")
	}
}
