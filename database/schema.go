package database

import (
	"context"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var fs embed.FS

func RunMigrations(ctx context.Context, db Database) {
	s, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Panic().
			Str("tag", "database:migration").
			Err(err).
			Msg("failed to create source driver instance")
	}
	defer s.Close()

	d, err := postgres.WithInstance(db.conn, &postgres.Config{})
	if err != nil {
		log.Panic().
			Str("tag", "database:migration").
			Err(err).
			Msg("failed to create database driver instance")
	}

	m, err := migrate.NewWithInstance("iofs", s, "postgres", d)
	if err != nil {
		log.Panic().
			Str("tag", "database:migration").
			Err(err).
			Msg("failed to create migration instance")
	}

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			return
		}

		log.Panic().
			Str("tag", "database:migration").
			Err(err).
			Msg("failed to migrate")
	}
}
