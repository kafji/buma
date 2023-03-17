package services

import (
	"context"

	"kafji.net/buma/hash"
)

type AddUser interface {
	// AddUser adds new user information into persistent data store.
	AddUser(ctx context.Context, email string, password []byte, salt []byte) bool
}

// CreateAccount creates a new user account.
func CreateAccount(ctx context.Context, au AddUser, email, password string) error {
	if email == "" {
		return ErrEmptyEmail
	}

	if password == "" {
		return ErrEmptyPassword
	}

	salt := hash.GenerateSalt()
	hashedPw := hash.HashPassword(password, salt)
	ok := au.AddUser(ctx, email, hashedPw, salt)
	if !ok {
		return ErrNonUniqueEmail
	}

	return nil
}
