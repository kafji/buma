package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"kafji.net/buma/services"
)

type createAccessTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createAccessTokenResponse struct {
	Token string `json:"token"`
}

func createTokenHandler(c echo.Context) error {
	ctx := c.Request().Context()

	env := getEnv(c)

	var req createAccessTokenRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	token, ok, err := services.CreateToken(ctx, env.getDB(), env.getDB(), req.Email, req.Password)
	if err != nil {
		switch err {
		case services.ErrEmptyEmail:
			return echo.NewHTTPError(http.StatusBadRequest)
		case services.ErrEmptyPassword:
			return echo.NewHTTPError(http.StatusBadRequest)
		default:
			log.Panic().
				Str("tag", "server").
				Err(err).
				Msg("create token error")
		}
	}
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	res := createAccessTokenResponse{token}
	return c.JSON(http.StatusOK, res)
}
