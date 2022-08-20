package server

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func newCSRF(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.config.SeverSecure,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func loadSession(next http.Handler) http.Handler {
	return app.session.LoadAndSave(next)
}
