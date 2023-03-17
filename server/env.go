package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"kafji.net/buma/database"
)

// Request environment.
type Environment interface {
	getDB() database.Database
	setUserID(int)
	getUserID() int
}

type environment struct {
	db     database.Database
	userID int
}

func NewEnvFactory(db database.Database) func() Environment {
	return func() Environment {
		return &environment{db, 0}
	}
}

func (s *environment) getDB() database.Database {
	return s.db
}

func (s *environment) setUserID(id int) {
	s.userID = id
}

func (s *environment) getUserID() int {
	return s.userID
}

const (
	envKey = "environment"
)

// requestEnvironment returns a middleware that will set a request environment into context.
func requestEnvironment(envf func() Environment) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			env := envf()
			c.Set(envKey, env)
			return next(c)
		}
	}
}

// getEnv returns a request environment.
func getEnv(c echo.Context) Environment {
	env, ok := c.Get(envKey).(Environment)
	if !ok {
		log.Panic().Msg("server: missing env from context")
	}
	return env
}
