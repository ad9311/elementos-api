package server

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/controller"
	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// mux.Use(middleware.Recoverer)
	mux.Use(loadSession)
	mux.Use(newCSRF)

	// Root
	mux.Get("/", controller.GetRoot)

	// Users
	mux.Route("/sign_in", func(r chi.Router) {
		r.Get("/", controller.GetSignIn)
		r.Post("/", controller.PostSignIn)
	})
	mux.Route("/sign_up", func(r chi.Router) {
		r.Get("/", controller.GetSignUp)
		r.Post("/", controller.PostSignUp)
	})
	mux.Post("/sign_out", controller.PostSignOut)

	// Landmarks
	mux.Get("/dashboard", controller.GetDashboard)
	mux.Route("/landmarks", func(r chi.Router) {
		r.Post("/", controller.PostNewLandmark)
		r.Get("/new", controller.GetNewLandmark)
		r.Route("/{landmarkID}", func(r chi.Router) {
			r.Get("/", controller.GetShowLandmark)
			r.Get("/edit", controller.GetEditLandmark)
			r.Post("/", controller.PostEditLandmark)
			r.Post("/delete", controller.PostDeleteLandmark)
		})
	})

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
