package createtoken

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
	"kafji.net/buma/hash"
)

type CreateTokenError string

func (s CreateTokenError) Error() string {
	return string(s)
}

const (
	ErrEmptyEmail    = CreateTokenError("email must not be empty")
	ErrEmptyPassword = CreateTokenError("password must not be empty")
	ErrUserNotFound  = CreateTokenError("user not found")
)

type GetUserByEmail interface {
	GetUserByEmail(ctx context.Context, email string) (id int, password []byte, salt []byte, found bool)
}

type AddToken interface {
	AddToken(ctx context.Context, userID int, token string)
}

func CreateToken(
	ctx context.Context,
	gube GetUserByEmail,
	at AddToken,
	email,
	password string,
) (
	token string,
	err error,
) {
	if email == "" {
		err = ErrEmptyEmail
		return
	}

	if password == "" {
		err = ErrEmptyPassword
		return
	}

	userID, hashedPw, salt, found := gube.GetUserByEmail(ctx, email)
	if !found {
		err = ErrUserNotFound
		return
	}

	if ok := hash.VerifyPassword(password, salt, hashedPw); !ok {
		err = ErrUserNotFound
		return
	}

	token = uuid.New().String()
	at.AddToken(ctx, userID, token)

	slog.Info("token created", "email", email)

	return
}
