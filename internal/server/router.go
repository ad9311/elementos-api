package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	signIn    = "/sign_in"
	signOut   = "/sign_out"
	dashboard = "/dashboard"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(loadSession)
	mux.Use(newCSRF)

	mux.Get(dashboard, getDashboard)

	mux.Get(signIn, getSignIn)
	mux.Post(signIn, postSignIn)

	mux.Post(signOut, postSignOut)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
