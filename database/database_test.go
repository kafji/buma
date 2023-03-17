package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testUser = struct {
	email    string
	password string
	salt     []byte
}{
	"test@example.com",
	"hunter2",
	[]byte{1, 1, 1, 1},
}

func addTestUser(t *testing.T, db Database) {
	ok := db.AddUser(context.Background(), testUser.email, []byte(testUser.password), testUser.salt)
	if !assert.True(t, ok) {
		return
	}
}

func getTestUser(ctx context.Context, t *testing.T, db Database) int {
	id, _, _, found := db.GetUserByEmail(ctx, testUser.email)
	if !assert.True(t, found) {
		t.Fail()
	}
	return id
}

func withTestUser(ctx context.Context, t *testing.T, f func(db Database, userID int)) {
	WithDatabase(ctx, t, func(db Database) {
		addTestUser(t, db)
		id := getTestUser(ctx, t, db)
		f(db, id)
	})
}
