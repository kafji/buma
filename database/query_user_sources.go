package database

import (
	"context"

	"golang.org/x/exp/slog"
	"kafji.net/buma/models"
)

func (s Database) QueryUserSources(ctx context.Context, userID int) []models.UserSource {
	rows, err := s.conn.QueryContext(ctx, "SELECT name, url FROM feed_sources WHERE user_id = $1;", userID)
	if err != nil {
		slog.Error("failed to query user feed sources", err)
		panic(err)
	}

	ss := []models.UserSource{}

	for rows.Next() {
		s := models.UserSource{}
		err := rows.Scan(&s.Name, &s.URL)
		if err != nil {
			slog.Error("failed to read query result", err)
			panic(err)
		}

		ss = append(ss, s)
	}

	return ss
}
