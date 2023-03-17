package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"kafji.net/buma/services"
)

func authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		env := getEnv(c)

		header := c.Request().Header.Get("authorization")
		if header == "" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		token, ok := strings.CutPrefix(header, "Bearer ")
		if !ok {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		userID, ok := services.Authenticate(ctx, env.getDB(), token)
		if !ok {
			return echo.NewHTTPError(http.StatusForbidden)
		}

		env.setUserID(userID)

		return next(c)
	}
}
