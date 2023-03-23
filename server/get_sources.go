package server

import (
	"net/http"

	"github.com/go-chi/render"
	"kafji.net/buma/services"
)

type getSourcesResponse struct {
	Sources []services.UserFeedSource `json:"sources"`
}

func getSourcesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	sources := services.GetSources(ctx, getDB(ctx), getUserID(ctx))

	res := getSourcesResponse{sources}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &res)
}
