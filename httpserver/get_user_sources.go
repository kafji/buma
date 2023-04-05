package httpserver

import (
	"net/http"

	"github.com/go-chi/render"
	"kafji.net/buma/models"
	getusersources "kafji.net/buma/services/get_user_sources"
)

type getUserSourcesResponse struct {
	Sources []models.UserSource `json:"sources"`
}

func getUserSourcesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	srcs := getusersources.GetUserSources(ctx, getDB(ctx), getUserID(ctx))

	res := getUserSourcesResponse{srcs}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &res)
}
