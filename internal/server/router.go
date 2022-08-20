package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	login     = "/login"
	dashboard = "/dashboard"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(loadSession)
	mux.Use(newCSRF)

	mux.Get(dashboard, getDashboard)

	mux.Get(login, getLogin)
	mux.Post(login, postLogin)

	mux.Post("/logout", postLogout)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
