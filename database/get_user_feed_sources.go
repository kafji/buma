package database

import (
	"context"

	"github.com/rs/zerolog/log"
	"kafji.net/buma/services"
)

func (s Database) GetUserFeedSources(ctx context.Context, userID int) []services.UserFeedSource {
	rows, err := s.conn.QueryContext(ctx, "SELECT name, url FROM feed_sources WHERE user_id = $1;", userID)
	if err != nil {
		log.Panic().Str("tag", "database").Err(err).Msg("failed to query user feed sources")
	}

	ss := []services.UserFeedSource{}

	for rows.Next() {
		s := services.UserFeedSource{}
		err := rows.Scan(&s.Name, &s.URL)
		if err != nil {
			log.Panic().Str("tag", "database").Err(err).Msg("failed to read query result")
		}

		ss = append(ss, s)
	}

	return ss
}
