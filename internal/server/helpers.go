package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func userSignedIn(r *http.Request) bool {
	return app.session.GetBool(r.Context(), "signedIn")
}

func validateForm(r *http.Request, fields []string) error {
	for _, v := range fields {
		if r.PostFormValue(v) == "" {
			return errors.New("invalid form or required fields not present")
		}
	}

	return nil
}

func validatePassword(password string, encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(encryptedPassword),
		[]byte(password),
	)
	if err != nil {
		return err
	}

	return nil
}

func confirmPasswords(password string, passwordConfirmation string) error {
	if password != passwordConfirmation {
		return fmt.Errorf("passwords don't match")
	}
	return nil
}

func encryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(encryptedPassword), err
	}

	return string(encryptedPassword), nil
}

func validateDate(date time.Time) error {
	if time.Now().After(date) {
		return fmt.Errorf("date already passed")
	}
	return nil
}
