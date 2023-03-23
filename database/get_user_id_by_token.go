package database

import (
	"context"
	"database/sql"

	"golang.org/x/exp/slog"
)

func (s Database) GetUserIDByToken(ctx context.Context, token string) (userID int, found bool) {
	q := "SELECT users.id FROM users " +
		"JOIN user_tokens ON users.id = user_tokens.user_id " +
		"WHERE user_tokens.token = $1;"
	row := s.conn.QueryRow(q, token)

	err := row.Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			found = false
			return
		}
		slog.Error("failed to get user id by token", err)
		panic(err)
	}

	found = true
	return
}
