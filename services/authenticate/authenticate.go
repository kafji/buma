package authenticate

import "context"

type AuthenticateError string

func (s AuthenticateError) Error() string {
	return string(s)
}

const ErrInvalidToken = AuthenticateError("user not found or invalid token")

type GetUserIDByToken interface {
	GetUserIDByToken(ctx context.Context, token string) (id int, found bool)
}

func Authenticate(ctx context.Context, guibt GetUserIDByToken, token string) (userID int, err error) {
	userID, ok := guibt.GetUserIDByToken(ctx, token)
	if !ok {
		return 0, ErrInvalidToken
	}
	return userID, nil
}
