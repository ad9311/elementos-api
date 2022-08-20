package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(sessionsLoad)
	// mux.Use(newCsrf)

	// fileServer := http.FileServer(http.Dir("./web/static/"))
	// mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
