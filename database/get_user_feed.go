package database

import (
	"context"

	"github.com/rs/zerolog/log"
	"kafji.net/buma/services"
)

func (s Database) GetUserFeed(ctx context.Context, userID int) services.UserFeed {
	q := "SELECT feed_items.id, feed_items.title, feed_items.url, feed_sources.name, feed_sources.url " +
		"FROM feed_items, feed_sources " +
		"WHERE feed_items.source_id = feed_sources.id " +
		"AND feed_sources.user_id = $1;"
	rows, err := s.conn.QueryContext(ctx, q, userID)
	if err != nil {
		log.Panic().
			Str("tag", "database").
			Err(err).
			Msg("failed to query feed")
	}
	defer rows.Close()

	uf := services.UserFeed{}

	for rows.Next() {
		item := services.UserFeedItem{}
		rows.Scan(&item.ID, &item.Title, &item.URL, &item.SourceName, &item.SourceURL)
		uf.Items = append(uf.Items, item)
	}

	return uf
}
