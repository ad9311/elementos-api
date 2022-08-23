package validation

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ValidateSignUpForm ...
func ValidateSignUpForm(r *http.Request) error {
	params := []string{
		"first-name",
		"last-name",
		"username",
		"email",
		"password",
		"code",
	}
	err := ValidateFormParams(r, params)
	if err != nil {
		return err
	}

	err = ValidatePasswordConfirmation(
		r.PostFormValue("password"),
		r.PostFormValue("password-confirmation"),
	)
	if err != nil {
		return err
	}

	return nil
}

// ValidateFormParams ...
func ValidateFormParams(r *http.Request, fields []string) error {
	for _, v := range fields {
		if r.PostFormValue(v) == "" {
			return errors.New("invalid form or required fields not present")
		}
	}

	return nil
}

// ValidatePassword ...
func ValidatePassword(password string, encryptedPassword string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(encryptedPassword),
		[]byte(password),
	)
	if err != nil {
		return err
	}

	return nil
}

// ValidatePasswordConfirmation ...
func ValidatePasswordConfirmation(password string, passwordConfirmation string) error {
	if password != passwordConfirmation {
		return fmt.Errorf("passwords don't match")
	}
	return nil
}

// ValidateDateAfter ...
func ValidateDateAfter(date time.Time) error {
	if time.Now().After(date) {
		return fmt.Errorf("date already passed")
	}

	return nil
}
