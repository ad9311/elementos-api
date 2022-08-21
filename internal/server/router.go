package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const (
	root      = "/"
	signIn    = "/sign_in"
	signOut   = "/sign_out"
	signUp    = "/sign_up"
	dashboard = "/dashboard"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(loadSession)
	mux.Use(newCSRF)

	mux.Get(root, getRoot)

	mux.Get(dashboard, getDashboard)

	mux.Get(signIn, getSignIn)
	mux.Post(signIn, postSignIn)

	mux.Get(signUp, getSignUp)
	mux.Post(signUp, postSignUp)

	mux.Post(signOut, postSignOut)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
