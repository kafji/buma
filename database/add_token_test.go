package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToken(t *testing.T) {
	ctx := context.Background()

	withTestUser(ctx, t, func(db Database, testUserID int) {
		token := "token"

		db.AddToken(ctx, testUserID, token)

		userID, ok := db.GetUserIDByToken(ctx, token)
		if !assert.True(t, ok) {
			return
		}

		assert.Equal(t, testUserID, userID)
	})
}
