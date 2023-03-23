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
		slog.Error("failed to create source driver instance", err)
		panic(err)
	}
	defer s.Close()

	d, err := postgres.WithInstance(db.conn, &postgres.Config{})
	if err != nil {
		slog.Error("failed to create database driver instance", err)
		panic(err)
	}

	m, err := migrate.NewWithInstance("iofs", s, "postgres", d)
	if err != nil {
		slog.Error("failed to create migration instance", err)
		panic(err)
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		slog.Error("migration failed", err)
		panic(err)
	}
}
