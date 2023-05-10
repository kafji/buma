package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"kafji.net/buma/database"
)

func setupRouter(r *chi.Mux, db database.Database) {
	r.Use(middleware.RequestLogger(newSlogFormatter()))
	r.Use(middleware.Recoverer)

	r.Use(withDB(db))

	r.Post("/signup", createAccountHandler)
	r.Post("/login", createTokenHandler)

	me := chi.NewRouter()
	me.Use(authorization)
	me.Post("/source", addSourceHandler)
	me.Get("/sources", getUserSourcesHandler)
	me.Get("/feed", getUserFeedHandler)
	r.Mount("/me", me)
}
