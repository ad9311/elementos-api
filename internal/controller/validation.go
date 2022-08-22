package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func validateSignUpForm(r *http.Request) error {
	params := []string{
		"first-name",
		"last-name",
		"username",
		"email",
		"password",
		"code",
	}
	err := validateFormParams(r, params)
	if err != nil {
		return err
	}

	err = validatePasswordConfirmation(
		r.PostFormValue("password"),
		r.PostFormValue("password-confirmation"),
	)
	if err != nil {
		return err
	}

	return nil
}

func validateFormParams(r *http.Request, fields []string) error {
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

func validatePasswordConfirmation(password string, passwordConfirmation string) error {
	if password != passwordConfirmation {
		return fmt.Errorf("passwords don't match")
	}
	return nil
}

func validateDate(date time.Time) error {
	if time.Now().After(date) {
		return fmt.Errorf("date already passed")
	}
	return nil
}
