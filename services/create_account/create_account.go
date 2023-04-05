package createaccount

import (
	"context"

	"golang.org/x/exp/slog"
	"kafji.net/buma/hash"
)

type CreateAccountError string

func (s CreateAccountError) Error() string {
	return string(s)
}

const (
	ErrEmptyEmail           = CreateAccountError("email must not be empty")
	ErrEmptyPassword        = CreateAccountError("password must not be empty")
	ErrAccountAlreadyExists = CreateAccountError("account with the specified email already exist")
)

type AddUser interface {
	// AddUser adds new user information into persistent data store.
	AddUser(ctx context.Context, email string, password []byte, salt []byte) bool
}

// CreateAccount creates a new user account.
//
// Returns [`CreateAccountError`] if error occured.
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
		return ErrAccountAlreadyExists
	}

	slog.Info("user account created", "email", email)

	return nil
}
