package val

import (
	"errors"
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/errs"
)

// ValidateUserSignIn ...
func ValidateUserSignIn(dtbs *db.Database, r *http.Request) (db.User, error) {
	params := []string{"username", "password"}
	if err := checkFormParams(r, params); err != nil {
		return db.User{}, err
	}

	user, err := dtbs.SelectUserByUsername(r.PostFormValue("username"))
	if err != nil {
		return db.User{}, errors.New(errs.WrongPswdOrUser)
	}

	err = checkFormPassword(r.PostFormValue("password"), user.HashedPassword)
	user.HashedPassword = ""
	if err != nil {
		return db.User{}, errors.New(errs.WrongPswdOrUser)
	}

	err = dtbs.UpdateUserLastLogin(user.ID)
	if err != nil {
		return db.User{}, err
	}

	return user, nil
}
