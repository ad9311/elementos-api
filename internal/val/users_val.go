package val

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
	"golang.org/x/crypto/bcrypt"
)

// ValidateUserSignIn ...
func ValidateUserSignIn(dtbs *db.Database, r *http.Request) (*db.User, error) {
	var user *db.User

	params := []string{"username", "password"}
	if err := checkFormParams(r, params); err != nil {
		return user, err
	}

	u, err := dtbs.SelectUserByUsername(r)
	user = u
	if err != nil {
		return user, err
	}

	err = checkPassword(r.PostFormValue("password"), u.HashedPassword)
	if err != nil {
		return user, err
	}

	err = dtbs.UpdateUserLastLogin(u)
	if err != nil {
		return user, err
	}

	return user, nil
}

// ValidateUserSignUp ...
func ValidateUserSignUp(dtbs *db.Database, r *http.Request) error {
	params := []string{
		"first-name",
		"last-name",
		"username",
		"email",
		"password",
		"password-confirmation",
		"code",
	}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	if err := checkPasswordConfirmation(
		r.PostFormValue("password"),
		r.PostFormValue("password-confirmation"),
	); err != nil {
		return err
	}

	ic, err := dtbs.SelectInvitationCode(r.PostFormValue("code"))
	if err != nil {
		return err
	}

	if err = checkDateAfter(ic.Validity); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(r.PostFormValue("password")),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	if err = dtbs.InsertUser(r, string(hashedPassword)); err != nil {
		return err
	}

	return nil
}
