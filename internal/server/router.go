package server

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/controller"
	"github.com/go-chi/chi"
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

	// mux.Use(middleware.Recoverer)
	mux.Use(loadSession)
	mux.Use(newCSRF)

	mux.Get(root, controller.GetRoot)

	mux.Get(dashboard, controller.GetDashboard)

	mux.Get(signIn, controller.GetSignIn)
	mux.Post(signIn, controller.PostSignIn)

	mux.Get(signUp, controller.GetSignUp)
	mux.Post(signUp, controller.PostSignUp)

	mux.Post(signOut, controller.PostSignOut)

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
