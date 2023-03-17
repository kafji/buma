package database

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog/log"
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
		} else {
			log.Panic().Str("tag", "database").Err(err).Str("email", email).Msg("failed to get user by email")
		}
	}

	found = true
	return
}
