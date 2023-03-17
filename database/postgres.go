package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// https://www.postgresql.org/docs/current/errcodes-appendix.html

var conflictCode = pq.ErrorCode("23505")

type PostgresConfig struct {
	DBname   string
	User     string
	Password string
	Host     string
	Port     int
}

// PostgresConnect opens and verifies connection to postgres.
func PostgresConnect(ctx context.Context, cfg PostgresConfig) Database {
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		cfg.DBname,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port)

	log.Info().Str("tag", "database").Str("dsn", dsn).Msg("connecting to postgres")
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Panic().Str("tag", "database").Err(err).Msg("failed to open connection to postgres")
	}

	err = conn.PingContext(ctx)
	if err != nil {
		log.Panic().Str("tag", "database").Err(err).Msg("failed to verify connection to postgres")
	}

	return Database{conn}
}
