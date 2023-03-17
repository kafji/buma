package database

import (
	"context"

	"github.com/rs/zerolog/log"
	"kafji.net/buma/services"
)

func (s Database) PutFeed(ctx context.Context, items []services.StorableFeedItem) {
	tx, err := s.conn.Begin()
	if err != nil {
		log.Panic().Err(err).Msg("database: failed to begin transaction")
	}

	q := "WITH source AS (SELECT id FROM feed_sources WHERE url = $1) " +
		"INSERT INTO feed_items (source_id, url, title) SELECT source.id, $2, $3 FROM source;"
	for _, i := range items {
		_, err := tx.ExecContext(ctx, q, i.SourceURL, i.URL, i.Title)
		if err != nil {
			log.Panic().
				Str("source_url", i.SourceURL).
				Str("url", i.URL).
				Str("title", i.Title).
				Err(err).
				Msg("database: failed to insert feed item")
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Panic().Err(err).Msg("database: failed to commit transaction")
	}
}
