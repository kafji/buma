package httpserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"kafji.net/buma/database"
)

func StartServer(port int, db database.Database) {
	r := chi.NewRouter()

	setupRouter(r, db)

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Panic(err)
	}
}
