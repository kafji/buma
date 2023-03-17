package database

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (s Database) AddToken(ctx context.Context, userID int, token string) {
	_, err := s.conn.ExecContext(ctx, "INSERT INTO user_tokens (user_id, token) VALUES ($1, $2);", userID, token)
	if err != nil {
		log.Panic().
			Str("tag", "database").
			Err(err).
			Msg("database: failed to add user token")
	}
}
