package database

import (
	"context"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"golang.org/x/exp/slog"
)

//go:embed migrations/*.sql
var fs embed.FS

func RunMigrations(ctx context.Context, db Database) {
	s, err := iofs.New(fs, "migrations")
	if err != nil {
		slog.Error("failed to create source driver instance", "err", err)
		panic(err)
	}
	defer s.Close()

	d, err := postgres.WithInstance(db.conn, &postgres.Config{})
	if err != nil {
		slog.Error("failed to create database driver instance", "err", err)
		panic(err)
	}

	m, err := migrate.NewWithInstance("iofs", s, "postgres", d)
	if err != nil {
		slog.Error("failed to create migration instance", "err", err)
		panic(err)
	}

	err = m.Up()
	if err == migrate.ErrNoChange {
		return
	}
	if err != nil {
		slog.Error("migration failed", "err", err)
		panic(err)
	}
}
