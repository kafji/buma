package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"kafji.net/buma/services"
)

type createAccountRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func createAccountHandler(c echo.Context) error {
	ctx := c.Request().Context()
	env := getEnv(c)

	var req createAccountRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := services.CreateAccount(ctx, env.getDB(), req.Email, req.Password)
	if err != nil {
		switch err {
		case services.ErrEmptyEmail:
			return echo.NewHTTPError(http.StatusBadRequest)

		case services.ErrEmptyPassword:
			return echo.NewHTTPError(http.StatusBadRequest)

		case services.ErrNonUniqueEmail:
			return echo.NewHTTPError(http.StatusConflict)

		default:
			log.Panic().Err(err).Msg("server: failed to create account")
		}
	}

	return c.JSON(http.StatusOK, nil)
}
