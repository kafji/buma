package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByEmail(t *testing.T) {
	ctx := context.Background()

	WithDatabase(ctx, t, func(db Database) {
		addTestUser(t, db)

		_, pw, salt, ok := db.GetUserByEmail(ctx, testUser.email)
		if !assert.True(t, ok) {
			return
		}

		assert.Equal(t, testUser.password, string(pw))
		assert.Equal(t, testUser.salt, salt)
	})
}
