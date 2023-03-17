package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"kafji.net/buma/services"
)

func TestPutFeed(t *testing.T) {
	ctx := context.Background()

	withTestUser(ctx, t, func(db Database, userID int) {
		_, ok := db.AddFeedSource(ctx, userID, "Test Source", "http://example.com")
		if !assert.True(t, ok) {
			return
		}

		db.PutFeed(
			ctx,
			[]services.StorableFeedItem{
				{
					FetchedFeedItem: services.FetchedFeedItem{
						Title: "Test Item",
						URL:   "http://example.com/item",
					},
					SourceURL: "http://example.com",
				},
			})

		feed := db.GetUserFeed(ctx, userID)
		assert.Equal(t, services.UserFeed{
			Items: []services.UserFeedItem{
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
