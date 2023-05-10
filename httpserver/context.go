package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/exp/slog"
	"kafji.net/buma/database"
)

type ctxKey int

const (
	dbKey ctxKey = iota
	userIDKey
)

func getDB(ctx context.Context) database.Database {
	return MustValue[database.Database](ctx, dbKey)
}

func withDB(db database.Database) func(http.Handler) http.Handler {
	return WithValue(dbKey, db)
}

func getUserID(ctx context.Context) int {
	return MustValue[int](ctx, userIDKey)
}

func withUserID(userID int) func(http.Handler) http.Handler {
	return WithValue(userIDKey, userID)
}

func MustValue[T any](ctx context.Context, key any) T {
	v, ok := ctx.Value(key).(T)
	if !ok {
		msg := fmt.Sprintf("failed to get value from context with key `%v`", key)
		slog.Error(msg)
		panic(msg)
	}
	return v
}

func WithValue[T any](key any, val T) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx2 := context.WithValue(ctx, key, val)
			r2 := r.WithContext(ctx2)
			next.ServeHTTP(w, r2)
		})
	}
}
