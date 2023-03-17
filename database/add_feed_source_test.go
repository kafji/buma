package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSource = struct {
	name string
	url  string
}{
	"Rust Blog",
	"https://blog.rust-lang.org/feed.xml",
}

func TestAddFeedSource(t *testing.T) {
	ctx := context.Background()

	withTestUser(ctx, t, func(db Database, userID int) {
		userSourceID, ok := db.AddFeedSource(ctx, userID, testSource.name, testSource.url)

		if !assert.True(t, ok) {
			return
		}

		assert.Equal(t, 1, userSourceID)
	})
}
