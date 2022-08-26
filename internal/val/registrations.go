package val

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
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

	if err := checkPasswordConfirmation(
		formMap["password"],
		formMap["password_confirmation"],
	); err != nil {
		return err
	}

	// invCode, err := dtbs.SelectInvitationCode(formMap["invitation_code"])
	// if err != nil {
	// 	return err
	// }

	// if err = checkDateAfter(invCode.Validity); err != nil {
	// 	return err
	// }

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(formMap["password"]),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}
	formMap["hashed_password"] = string(hashedPassword)

	if err = dtbs.InsertUser(formMap); err != nil {
		return err
	}

	return nil
}
