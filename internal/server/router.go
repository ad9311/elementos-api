package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Group(managerRoutes)
	mux.Group(apiv1Routes)

	return mux
}
