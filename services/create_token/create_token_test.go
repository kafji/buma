package createtoken

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"kafji.net/buma/hash"
)

type fakeGetUserByEmail struct {
	users map[string]struct {
		id       int
		password []byte
		salt     []byte
	}
}

func newFakeGetUserByEmail(users ...struct {
	email    string
	password string
},
) fakeGetUserByEmail {
	m := map[string]struct {
		id       int
		password []byte
		salt     []byte
	}{}
	for i, u := range users {
		id := i + 1
		salt := hash.GenerateSalt()
		hashedPw := hash.HashPassword(string(u.password), salt)
		m[u.email] = struct {
			id       int
			password []byte
			salt     []byte
		}{id, hashedPw, salt}
	}
	return fakeGetUserByEmail{m}
}

func (s fakeGetUserByEmail) GetUserByEmail(
	ctx context.Context,
	email string,
) (id int, password []byte, salt []byte, found bool) {
	u, ok := s.users[email]
	if !ok {
		found = false
		return
	}
	id = u.id
	password = u.password
	salt = u.salt
	found = true
	return
}

type fakeAddToken struct {
	tokens map[string]int
}

func newFakeAddToken() fakeAddToken {
	return fakeAddToken{map[string]int{}}
}

func (s *fakeAddToken) AddToken(ctx context.Context, userID int, token string) {}

func TestCreateToken(t *testing.T) {
	gube := newFakeGetUserByEmail(struct {
		email    string
		password string
	}{
		"test@example.com", "hunter2",
	})
	at := newFakeAddToken()

	token, err := CreateToken(context.Background(), gube, &at, "test@example.com", "hunter2")

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestCreateTokenInvalidPassword(t *testing.T) {
	gube := newFakeGetUserByEmail(struct {
		email    string
		password string
	}{
		"test@example.com", "hunter2",
	})
	at := newFakeAddToken()

	_, err := CreateToken(context.Background(), gube, &at, "test@example.com", "hunter3")

	assert.Equal(t, ErrUserNotFound, err)
}

func TestCreateTokenAccountNotExist(t *testing.T) {
	gube := newFakeGetUserByEmail()
	at := newFakeAddToken()

	_, err := CreateToken(context.Background(), gube, &at, "test@example.com", "hunter2")

	assert.Equal(t, ErrUserNotFound, err)
}

func TestCreateTokenEmptyEmail(t *testing.T) {
	gube := newFakeGetUserByEmail()
	at := newFakeAddToken()

	_, err := CreateToken(context.Background(), gube, &at, "", "hunter2")

	assert.Equal(t, ErrEmptyEmail, err)
}

func TestCreateTokenEmptyPassword(t *testing.T) {
	gube := newFakeGetUserByEmail()
	at := newFakeAddToken()

	_, err := CreateToken(context.Background(), gube, &at, "test@example.com", "")

	assert.Equal(t, ErrEmptyPassword, err)
}
