package services

import (
	"context"

	"github.com/google/uuid"
	"kafji.net/buma/hash"
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
) (token string, ok bool, err error) {
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
		ok = false
		return
	}

	ok = hash.VerifyPassword(password, salt, hashedPw)
	if !ok {
		return
	}

	token = uuid.New().String()
	at.AddToken(ctx, userID, token)

	return
}
