package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"kafji.net/buma/services"
)

type getFeedResponse struct {
	Feed services.UserFeed `json:"feed"`
}

func getFeedHandler(c echo.Context) error {
	ctx := c.Request().Context()

	env := getEnv(c)

	feed := services.GetFeed(ctx, env.getDB(), env.getUserID())

	res := getFeedResponse{feed}
	return c.JSON(http.StatusOK, res)
}
