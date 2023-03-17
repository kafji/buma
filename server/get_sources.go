package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"kafji.net/buma/services"
)

type getSourcesResponse struct {
	Sources []services.UserFeedSource `json:"sources"`
}

func getSourcesHandler(c echo.Context) error {
	ctx := c.Request().Context()

	env := getEnv(c)

	sources := services.GetSources(ctx, env.getDB(), env.getUserID())

	res := getSourcesResponse{sources}
	return c.JSON(http.StatusOK, res)
}
