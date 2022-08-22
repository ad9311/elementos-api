package sess

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
	"github.com/alexedwards/scs/v2"
)

// Data ...
type Data struct {
	CurrentUser *db.User
	Session     *scs.SessionManager
	CSRFToken   string
	StringMap   map[string]string
}

var sessionData Data

// Init ...
func Init(secure bool) *Data {
	sessionData.Session = scs.New()
	sessionData.Session.Cookie.Persist = true
	sessionData.Session.Cookie.SameSite = http.SameSiteLaxMode
	sessionData.Session.Cookie.Secure = secure

	return &sessionData
}
