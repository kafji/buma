package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"kafji.net/buma/database"
)

type contextKey int

const (
	dbCtxKey contextKey = iota
	userIDCtxKey
)

func getVal[T any](ctx context.Context, key contextKey) T {
	v, ok := ctx.Value(key).(T)
	if !ok {
		panic(fmt.Sprintf("missing %v from context", key))
	}
	return v
}

func getDB(ctx context.Context) database.Database {
	return getVal[database.Database](ctx, dbCtxKey)
}

func withDB(db database.Database) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), dbCtxKey, db)))
		})
	}
}

func getUserID(ctx context.Context) int {
	return getVal[int](ctx, userIDCtxKey)
}

func withUserID(userID int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), userIDCtxKey, userID)))
		})
	}
}
