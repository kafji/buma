package database

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
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
		} else {
			log.Panic().Str("tag", "tag").Err(err).Msg("failed to get user id by token")
		}
	}

	found = true
	return
}
