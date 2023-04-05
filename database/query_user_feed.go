package database

import (
	"context"

	"golang.org/x/exp/slog"
	"kafji.net/buma/models"
)

func (s Database) QueryUserFeed(ctx context.Context, userID int) models.UserFeed {
	q := "SELECT feed_items.id, feed_items.title, feed_items.url, feed_sources.name, feed_sources.url " +
		"FROM feed_items, feed_sources " +
		"WHERE feed_items.source_id = feed_sources.id " +
		"AND feed_sources.user_id = $1;"

	rows, err := s.conn.QueryContext(ctx, q, userID)
	if err != nil {
		slog.Error("failed to query user feed", "err", err)
		panic(err)
	}
	defer rows.Close()

	f := models.UserFeed{}

	for rows.Next() {
		item := models.UserFeedItem{}
		rows.Scan(&item.ID, &item.Title, &item.URL, &item.SourceName, &item.SourceURL)
		f.Items = append(f.Items, item)
	}

	return f
}
