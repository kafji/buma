package httpserver

import (
	"context"
	"net/http"

	"github.com/kafji/quari"
	"github.com/kafji/quari/httpserver/ctxval"
	"kafji.net/buma/database"
)

type ctxKey int

const (
	dbKey ctxKey = iota
	userIDKey
)

func getDB(ctx context.Context) database.Database {
	return quari.MustValue[database.Database](ctx, dbKey)
}

func withDB(db database.Database) func(http.Handler) http.Handler {
	return ctxval.WithValue(dbKey, db)
}

func getUserID(ctx context.Context) int {
	return quari.MustValue[int](ctx, userIDKey)
}

func withUserID(userID int) func(http.Handler) http.Handler {
	return ctxval.WithValue(userIDKey, userID)
}
