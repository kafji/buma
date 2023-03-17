package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	ctx := context.Background()

	WithDatabase(ctx, t, func(db Database) {
		ok := db.AddUser(ctx, testUser.email, []byte(testUser.password), testUser.salt)

		if !assert.True(t, ok) {
			return
		}

		id := getTestUser(ctx, t, db)
		assert.Greater(t, id, 0)
	})
}

func TestAddUserNonUniqueEmail(t *testing.T) {
	ctx := context.Background()

	WithDatabase(ctx, t, func(db Database) {
		addUser := func() bool {
			return db.AddUser(ctx, testUser.email, []byte(testUser.password), testUser.salt)
		}

		addUser()
		ok := addUser()

		assert.False(t, ok)
	})
}
