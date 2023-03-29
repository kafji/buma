package authenticate

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeGetUserIDByToken struct {
	tokens []string
}

func newFakeGetUserIDByToken(tokens ...string) fakeGetUserIDByToken {
	return fakeGetUserIDByToken{tokens}
}

func (s fakeGetUserIDByToken) GetUserIDByToken(ctx context.Context, token string) (id int, found bool) {
	for i, t := range s.tokens {
		if t == token {
			return i + 1, true
		}
	}

	found = false
	return
}

func TestAuthenticate(t *testing.T) {
	guibt := newFakeGetUserIDByToken("t0k3n")

	userID, err := Authenticate(context.Background(), guibt, "t0k3n")

	assert.Nil(t, err)
	assert.Equal(t, 1, userID)
}

func TestAuthenticateFail(t *testing.T) {
	guibt := newFakeGetUserIDByToken()

	_, err := Authenticate(context.Background(), guibt, "t0k3n")

	assert.Equal(t, ErrInvalidToken, err)
}
