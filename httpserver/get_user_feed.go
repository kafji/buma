package httpserver

import (
	"net/http"

	"github.com/go-chi/render"
	getuserfeed "kafji.net/buma/services/get_user_feed"
)

func getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	feed := getuserfeed.GetUserFeed(ctx, getDB(ctx), getUserID(ctx))

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &feed)
}
