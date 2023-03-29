package main

import (
	"context"

	"golang.org/x/exp/slog"
	"kafji.net/buma/app"
	"kafji.net/buma/config"
	"kafji.net/buma/httpserver"
)

func main() {
	ctx := context.Background()

	cfg := config.ReadConfig()

	db := app.Database(ctx, &cfg)
	defer func() {
		err := db.Close()
		if err != nil {
			slog.Error("error while closing database", err)
			panic(err)
		}
	}()

	slog.Info("starting http server app")

	httpserver.StartServer(cfg.HTTPServer.Port, db)
}
