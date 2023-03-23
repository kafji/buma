package server

import (
	"net/http"
	"strings"

	"kafji.net/buma/services"
)

func authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		header := r.Header.Get("authorization")
		if header == "" {
			msg := "missing authorization header"
			badRequest(w, r, &msg)
			return
		}

		token, ok := strings.CutPrefix(header, "Bearer ")
		if !ok {
			msg := "invalid authorization header format"
			badRequest(w, r, &msg)
			return
		}

		userID, ok := services.Authenticate(ctx, getDB(ctx), token)
		if !ok {
			msg := "invalid credentials"
			forbidden(w, r, &msg)
			return
		}

		withUserID(userID)(next).ServeHTTP(w, r)
	})
}
