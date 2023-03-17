package server

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"kafji.net/buma/database"
)

func StartServer(port int, db database.Database) error {
	r := echo.New()

	r.HideBanner = true
	r.HidePort = true

	envf := NewEnvFactory(db)
	SetupRouter(r, envf)

	addr := fmt.Sprintf(":%d", port)

	routes, err := json.Marshal(r.Routes())
	if err != nil {
		log.Panic().Msg("server: failed to generate routes")
	}

	log.Info().Str("addr", addr).RawJSON("routes", routes).Msg("server: starting")

	return r.Start(addr)
}

func SetupRouter(r *echo.Echo, envf func() Environment) {
	r.Use(requestLogger())
	r.Use(middleware.Recover())
	r.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(120)))
	r.Use(requestEnvironment(envf))

	account := r.Group("/account")
	account.POST("", createAccountHandler)
	account.POST("/token", createTokenHandler)

	user := r.Group("/user")
	user.Use(authorization)
	user.POST("/source", addSourceHandler)
	user.GET("/sources", getSourcesHandler)
	user.DELETE("/source/:id", deleteSourceHandler)
	user.GET("/feed", getFeedHandler)
}
