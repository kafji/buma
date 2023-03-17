package database

import (
	"context"

	"github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func (s Database) AddUser(ctx context.Context, email string, password []byte, salt []byte) bool {
	q := "INSERT INTO users (email, password, salt) VALUES ($1, $2, $3);"
	_, err := s.conn.ExecContext(ctx, q, email, password, salt)
	if err != nil {
		switch v := err.(type) {
		case *pq.Error:
			if v.Code == pq.ErrorCode("23505") {
				return false
			}
		}

		log.Panic().
			Str("tag", "database").
			Err(err).
			Msg("failed to add user")
	}

	return true
}
