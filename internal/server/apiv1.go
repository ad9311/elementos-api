package server

import (
	"github.com/ad9311/hitomgr/internal/apictrl"
	"github.com/go-chi/chi"
)

func apiv1Routes(r chi.Router) {
	r.Route("/api/v1/landmarks", func(r chi.Router) {
		r.Get("/", apictrl.GetLandmarks)
	})
}
