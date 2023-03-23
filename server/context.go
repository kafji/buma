package server

import (
	"context"
	"fmt"
	"net/http"

	"kafji.net/buma/database"
)

type contextKey int

func (s contextKey) String() string {
	return ""
}

const (
	dbCtxKey contextKey = iota
	userIDCtxKey
)

func getVal[T any](ctx context.Context, key contextKey) T {
	v, ok := ctx.Value(key).(T)
	if !ok {
		panic(fmt.Sprintf("missing %s from context", key.String()))
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
