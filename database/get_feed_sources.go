package database

import (
	"context"

	"golang.org/x/exp/slog"
)

func (s Database) GetFeedSources(ctx context.Context) []string {
	rows, err := s.conn.QueryContext(ctx, "SELECT DISTINCT url FROM feed_sources;")
	if err != nil {
		slog.Error("failed to query feed sources", err)
		panic(err)
	}
	defer rows.Close()

	urls := []string{}

	for rows.Next() {
		var url string

		err := rows.Scan(&url)
		if err != nil {
			slog.Error("failed reading row", err)
			panic(err)
		}

		urls = append(urls, url)
	}

	return urls
}
