package server

import (
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func userLoggedIn(r *http.Request) bool {
	return app.session.GetBool(r.Context(), "signedIn")
}

func validateForm(r *http.Request, fields []string) error {
	for _, v := range fields {
		if r.PostFormValue(v) == "" {
			return errors.New("invalid form")
		}
	}

	return nil
}

func validatePassword(r *http.Request, encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(encryptedPassword),
		[]byte(r.PostFormValue("password")),
	)
	if err != nil {
		return err
	}

	return nil
}
