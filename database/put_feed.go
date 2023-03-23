package database

import (
	"context"

	"golang.org/x/exp/slog"
	"kafji.net/buma/services"
)

func (s Database) PutFeed(ctx context.Context, items []services.StorableFeedItem) {
	tx, err := s.conn.Begin()
	if err != nil {
		slog.Error("database: failed to begin transaction", err)
		panic(err)
	}

	q := "WITH source AS (SELECT id FROM feed_sources WHERE url = $1) " +
		"INSERT INTO feed_items (source_id, url, title) SELECT source.id, $2, $3 FROM source;"
	for _, i := range items {
		_, err := tx.ExecContext(ctx, q, i.SourceURL, i.URL, i.Title)
		if err != nil {
			slog.Error("failed to insert feed item", err)
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("failed to commit transaction", err)
		panic(err)
	}
}
