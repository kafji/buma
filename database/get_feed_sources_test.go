package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeedSources(t *testing.T) {
	ctx := context.Background()

	withTestUser(ctx, t, func(db Database, userID int) {
		db.AddFeedSource(ctx, userID, "Test Source", "http://example.com")

		sources := db.GetFeedSources(ctx)

		assert.Equal(t,
			[]string{
				"http://example.com",
			},
			sources)
	})
}
