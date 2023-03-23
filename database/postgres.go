package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"golang.org/x/exp/slog"
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

	slog.Info("connecting to postgres", dsn)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("failed to open connection to postgres", err)
		panic(err)
	}

	err = conn.PingContext(ctx)
	if err != nil {
		slog.Error("failed to verify connection to postgres", err)
		panic(err)
	}

	return Database{conn}
}
