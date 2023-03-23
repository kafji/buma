package database

import (
	"context"
	"database/sql"

	"golang.org/x/exp/slog"
)

func (s Database) GetUserByEmail(
	ctx context.Context,
	email string,
) (id int, password []byte, salt []byte, found bool) {
	row := s.conn.QueryRowContext(ctx, "SELECT id, password, salt FROM users WHERE email = $1;", email)

	err := row.Scan(&id, &password, &salt)
	if err != nil {
		if err == sql.ErrNoRows {
			found = false
			return
		}
		slog.Error("failed to get user by email", err, email)
		panic(err)
	}

	found = true
	return
}
