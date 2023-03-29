package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"kafji.net/buma/database"
)

func SetupRouter(r *chi.Mux, db database.Database) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(withDB(db))

	r.Post("/signup", createAccountHandler)
	r.Post("/login", createTokenHandler)

	me := chi.NewRouter()
	me.Use(authorization)
	me.Post("/source", addSourceHandler)
	me.Get("/sources", getSourcesHandler)
	me.Get("/feed", getFeedHandler)
	r.Mount("/me", me)
}

func StartServer(port int, db database.Database) {
	r := chi.NewRouter()

	SetupRouter(r, db)

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Panic(err)
	}
}
