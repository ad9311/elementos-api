package server

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/ctrl"
	"github.com/go-chi/chi"
)

func managerRoutes(r chi.Router) {
	r.Use(loadSession)
	r.Use(newCSRF)

	// Root
	r.Get("/", ctrl.GetRoot)

	// Sessions
	r.Route("/sign_in", func(r chi.Router) {
		r.Get("/", ctrl.GetSignIn)
		r.Post("/", ctrl.PostSignIn)
	})
	r.Post("/sign_out", ctrl.PostSignOut)

	// Registrations
	r.Route("/sign_up", func(r chi.Router) {
		r.Get("/", ctrl.GetSignUp)
		r.Post("/", ctrl.PostSignUp)
	})

	// Landmarks
	r.Get("/dashboard", ctrl.GetDashboard)
	r.Route("/landmarks", func(r chi.Router) {
		r.Post("/", ctrl.PostNewLandmark)
		r.Get("/new", ctrl.GetNewLandmark)
		r.Route("/{landmarkID:[\\d]+}", func(r chi.Router) {
			r.Get("/", ctrl.GetShowLandmark)
			r.Get("/edit", ctrl.GetEditLandmark)
			r.Post("/", ctrl.PostEditLandmark)
			r.Post("/delete", ctrl.PostDeleteLandmark)
		})
	})

	// Categories
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", ctrl.GetCategories)
		r.Post("/", ctrl.PostCategory)
		r.Get("/new", ctrl.GetNewCategory)
		r.Route("/{categoryID:[\\d]+}", func(r chi.Router) {
			r.Get("/edit", ctrl.GetEditCategory)
			r.Post("/", ctrl.PostEditCategory)
			r.Post("/delete", ctrl.PostDeleteCategory)
		})
	})

	fileServer := http.FileServer(http.Dir("./web/static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
}
