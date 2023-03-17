package database

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func withPostgres(ctx context.Context, t *testing.T, f func(port int)) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:14",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
		Env:          map[string]string{"POSTGRES_PASSWORD": "password"},
	}
	psql, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
	}
	defer func() {
		if err := psql.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	port, err := psql.MappedPort(ctx, "5432/tcp")
	if err != nil {
		t.Error(err)
	}

	f(port.Int())
}

// WithDatabase provides fixture for tests that require instace of Database.
func WithDatabase(ctx context.Context, t *testing.T, f func(db Database)) {
	withPostgres(ctx, t, func(port int) {
		cfg := PostgresConfig{"postgres", "postgres", "password", "127.0.0.1", port}
		db := PostgresConnect(ctx, cfg)
		defer db.Close()

		RunMigrations(ctx, db)

		f(db)
	})
}
