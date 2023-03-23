package database

import (
	"context"

	"github.com/lib/pq"
	"golang.org/x/exp/slog"
)

func (s Database) AddFeedSource(ctx context.Context, userID int, sourceName, sourceURL string) (sourceID int, ok bool) {
	q := "INSERT INTO feed_sources (user_id, url, name) VALUES ($1, $2, $3) RETURNING id;"
	row := s.conn.QueryRowContext(ctx, q, userID, sourceURL, sourceName)
	err := row.Scan(&sourceID)
	if err != nil {
		switch v := err.(type) {
		case *pq.Error:
			if v.Code == conflictCode {
				ok = false
				return
			}
		}

		slog.Error("failed to add user", userID, sourceName, sourceURL)
		panic(err)
	}

	ok = true
	return
}
