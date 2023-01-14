package server

import (
	"github.com/ad9311/hitomgr/internal/api/apiv1"
	"github.com/go-chi/chi"
)

func apiv1Routes(r chi.Router) {
	r.Route("/api/v1/landmarks", func(r chi.Router) {
		r.Get("/", apiv1.GetLandmarks)
	})

	r.Route("/api/v1/categories", func(r chi.Router) {
		r.Get("/", apiv1.GetCategories)
	})
}
