package database

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (s Database) GetFeedSources(ctx context.Context) []string {
	rows, err := s.conn.QueryContext(ctx, "SELECT DISTINCT url FROM feed_sources;")
	if err != nil {
		log.Panic().
			Str("tag", "database").
			Err(err).
			Msg("failed to query feed sources")
	}
	defer rows.Close()

	urls := []string{}

	for rows.Next() {
		var url string

		err := rows.Scan(&url)
		if err != nil {
			log.Panic().
				Str("tag", "database").
				Err(err).
				Msg("failed reading row")
		}

		urls = append(urls, url)
	}

	return urls
}
