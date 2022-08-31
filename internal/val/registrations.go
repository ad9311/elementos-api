package val

import (
	"errors"
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/errs"
	"golang.org/x/crypto/bcrypt"
)

// ValidateUserSignUp ...
func ValidateUserSignUp(dtbs *db.Database, r *http.Request) error {
	params := []string{
		"first_name",
		"last_name",
		"username",
		"password",
		"password_confirmation",
		"invitation_code",
	}
	if err := checkFormParams(r, params); err != nil {
		return err
	}
	formMap := formToMap(r, params)

	inviation, err := dtbs.SelectInvitation(formMap["invitation_code"])
	if err != nil {
		return errors.New(errs.InvNotExists)
	}

	if err = checkDateAfter(inviation.ExpiresAt, "invitation"); err != nil {
		return err
	}

	if err := checkPasswordConfirmation(
		formMap["password"],
		formMap["password_confirmation"],
	); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(formMap["password"]),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return errors.New(errs.InternalErr)
	}
	formMap["hashed_password"] = string(hashedPassword)

	if err = dtbs.InsertUser(formMap); err != nil {
		return errors.New(errs.UserNotInserted)
	}

	return nil
}
