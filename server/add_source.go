package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"kafji.net/buma/services"
)

type addUserSourceRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func addSourceHandler(c echo.Context) error {
	ctx := c.Request().Context()

	env := getEnv(c)

	var req addUserSourceRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	_, err := services.AddSource(ctx, env.getDB(), env.getUserID(), req.Name, req.URL)
	if err != nil {
		switch err {
		case services.ErrNonUniqueSourceName:
			return echo.NewHTTPError(http.StatusBadRequest)
		default:
			log.Panic().Str("tag", "server").Err(err).Msg("add source error")
		}
	}

	return c.NoContent(http.StatusOK)
}
