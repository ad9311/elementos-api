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
		Secure:   serverSecure,
		SameSite: http.SameSiteLaxMode,
		Name:     "hitomgr",
	})

	return csrfHandler
}

func loadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
