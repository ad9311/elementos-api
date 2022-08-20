package server

import (
	"errors"
	"net/http"
)

func userLoggedIn(r *http.Request) bool {
	return app.session.GetBool(r.Context(), "login")
}

func validateForm(r *http.Request, fields []string) error {
	for _, v := range fields {
		if r.PostFormValue(v) == "" {
			return errors.New("invalid form")
		}
	}

	return nil
}
