package sess

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
	"github.com/alexedwards/scs/v2"
	"golang.org/x/crypto/bcrypt"
)

// App ...
type App struct {
	CurrentUser *db.User
	Landmark    *db.Landmark
	Landmarks   []*db.Landmark
	Session     *scs.SessionManager
	URL         string
	CSRFToken   string
	StringMap   map[string]string
}

var app App

// Init ...
func Init(secure bool) *App {
	app.Session = scs.New()
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = secure

	return &app
}

// IsUserSignedIn ...
func (d *App) IsUserSignedIn(r *http.Request) bool {
	if d.Session.GetBool(r.Context(), "signedIn") {
		return true
	}

	return false
}

// EncryptPassword ...
func (d *App) EncryptPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(encryptedPassword), err
	}

	return string(encryptedPassword), nil
}
