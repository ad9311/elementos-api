package server

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func userSignedIn(r *http.Request) bool {
	return app.session.GetBool(r.Context(), "signedIn")
}

func encryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(encryptedPassword), err
	}

	return string(encryptedPassword), nil
}
