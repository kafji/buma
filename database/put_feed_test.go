package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"kafji.net/buma/models"
	fetchfeeds "kafji.net/buma/services/fetch_feeds"
)

func TestPutFeed(t *testing.T) {
	ctx := context.Background()

	withTestUser(ctx, t, func(db Database, userID int) {
		_, ok := db.AddFeedSource(ctx, userID, "Test Source", "http://example.com")
		assert.True(t, ok)

		db.PutFeed(
			ctx,
			[]fetchfeeds.StorableFeedItem{
				{
					FetchedFeedItem: fetchfeeds.FetchedFeedItem{
						Title: "Test Item",
						URL:   "http://example.com/item",
					},
					SourceURL: "http://example.com",
				},
			})

		feed := db.QueryUserFeed(ctx, userID)
		assert.Equal(t, models.UserFeed{
			Items: []models.UserFeedItem{
				{
					ID:         1,
					Title:      "Test Item",
					URL:        "http://example.com/item",
					SourceName: "Test Source",
					SourceURL:  "http://example.com",
				},
			},
		},
			feed)
	})
}
