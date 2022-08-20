package server

import "net/http"

func userLoggedIn(r *http.Request) bool {
	return app.session.GetBool(r.Context(), "login")
}
