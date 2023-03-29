package httpserver

import (
	"net/http"

	"github.com/go-chi/render"
	"kafji.net/buma/services"
)

type getFeedResponse struct {
	Feed services.UserFeed `json:"feed"`
}

func getFeedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	feed := services.GetFeed(ctx, getDB(ctx), getUserID(ctx))

	res := getFeedResponse{feed}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &res)
}
