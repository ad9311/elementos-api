package server

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/ctrl"
	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// mux.Use(middleware.Recoverer)
	mux.Use(loadSession)
	mux.Use(newCSRF)

	// Root
	mux.Get("/", ctrl.GetRoot)

	// Sessions
	mux.Route("/sign_in", func(r chi.Router) {
		r.Get("/", ctrl.GetSignIn)
		r.Post("/", ctrl.PostSignIn)
	})
	mux.Post("/sign_out", ctrl.PostSignOut)

	// Registrations
	mux.Route("/sign_up", func(r chi.Router) {
		r.Get("/", ctrl.GetSignUp)
		r.Post("/", ctrl.PostSignUp)
	})

	// // Landmarks
	mux.Get("/dashboard", ctrl.GetDashboard)
	mux.Route("/landmarks", func(r chi.Router) {
		r.Post("/", ctrl.PostNewLandmark)
		r.Get("/new", ctrl.GetNewLandmark)
		r.Route("/{landmarkID}", func(r chi.Router) {
			r.Get("/", ctrl.GetShowLandmark)
			// 	r.Get("/edit", controller.GetEditLandmark)
			// 	r.Post("/", controller.PostEditLandmark)
			// 	r.Post("/delete", controller.PostDeleteLandmark)
		})
	})

	fileServer := http.FileServer(http.Dir("./web/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
