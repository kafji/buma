package database

import (
	"context"

	"golang.org/x/exp/slog"
)

func (s Database) AddToken(ctx context.Context, userID int, token string) {
	_, err := s.conn.ExecContext(ctx, "INSERT INTO user_tokens (user_id, token) VALUES ($1, $2);", userID, token)
	if err != nil {
		slog.Error("failed to add user token", err)
		panic(err)
	}
}
