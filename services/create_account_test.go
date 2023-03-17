package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeAddUser struct {
	users map[string]struct {
		password []byte
		salt     []byte
	}
}

func newFakeAddUser() fakeAddUser {
	return fakeAddUser{map[string]struct {
		password []byte
		salt     []byte
	}{}}
}

func (s fakeAddUser) AddUser(ctx context.Context, email string, password []byte, salt []byte) bool {
	if _, found := s.users[email]; found {
		return false
	}

	s.users[email] = struct {
		password []byte
		salt     []byte
	}{password, salt}

	return true
}

func TestCreateAccount(t *testing.T) {
	au := newFakeAddUser()

	err := CreateAccount(context.Background(), au, "test@example.com", "hunter2")

	assert.Nil(t, err)
}

func TestCreateAccountNonUniqueEmail(t *testing.T) {
	au := newFakeAddUser()

	err := CreateAccount(context.Background(), au, "test@example.com", "hunter2")
	assert.Nil(t, err)

	err = CreateAccount(context.Background(), au, "test@example.com", "hunter2")
	assert.Equal(t, ErrNonUniqueEmail, err)
}

func TestCreateAccountEmptyEmail(t *testing.T) {
	au := newFakeAddUser()

	err := CreateAccount(context.Background(), au, "", "hunter2")

	assert.Equal(t, ErrEmptyEmail, err)
}

func TestCreateAccountEmptyPassword(t *testing.T) {
	au := newFakeAddUser()

	err := CreateAccount(context.Background(), au, "test@example.com", "")

	assert.Equal(t, ErrEmptyPassword, err)
}
