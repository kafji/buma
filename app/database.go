package app

import (
	"context"

	"kafji.net/buma/config"
	"kafji.net/buma/database"
)

func Database(ctx context.Context, cfg *config.Config) database.Database {
	c := database.PostgresConfig{
		DBname:   cfg.Database.Name,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
	}
	db := database.PostgresConnect(ctx, c)

	database.RunMigrations(ctx, db)

	return db
}
