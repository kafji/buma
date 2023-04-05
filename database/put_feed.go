package database

import (
	"context"

	"golang.org/x/exp/slog"
	fetchfeeds "kafji.net/buma/services/fetch_feeds"
)

func (s Database) PutFeed(ctx context.Context, items []fetchfeeds.StorableFeedItem) {
	tx, err := s.conn.Begin()
	if err != nil {
		slog.Error("failed to begin transaction", "err", err)
		panic(err)
	}

	q := "WITH source AS (SELECT id FROM feed_sources WHERE url = $1) " +
		"INSERT INTO feed_items (source_id, url, title) SELECT source.id, $2, $3 FROM source " +
		"ON CONFLICT ON CONSTRAINT feed_items_url_key DO UPDATE SET title = $3;"
	stmt, err := tx.PrepareContext(ctx, q)
	if err != nil {
		slog.Error("failed to prepare statement", "err", err)
		panic(err)
	}

	for _, i := range items {
		_, err := stmt.ExecContext(ctx, i.SourceURL, i.URL, i.Title)
		if err != nil {
			slog.Error("failed to insert feed item", "err", err)
			panic(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		slog.Error("failed to commit transaction", "err", err)
		panic(err)
	}
}
