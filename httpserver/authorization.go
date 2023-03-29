package httpserver

import (
	"net/http"
	"strings"

	"kafji.net/buma/services/authenticate"
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

		userID, err := authenticate.Authenticate(ctx, getDB(ctx), token)
		if err != nil && err == authenticate.ErrInvalidToken {
			msg := "invalid credentials"
			forbidden(w, r, &msg)
		}

		withUserID(userID)(next).ServeHTTP(w, r)
	})
}
