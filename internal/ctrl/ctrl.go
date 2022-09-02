package ctrl

import (
	"encoding/gob"
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var database *db.Database

// SetUp ...
func SetUp(dtbs *db.Database, sess *scs.SessionManager) {
	database = dtbs
	session = sess
	gob.Register(db.User{})
}

func currentUser(r *http.Request) db.User {
	i := session.Get(r.Context(), "current_user")
	user, ok := i.(db.User)
	if ok {
		return user
	}

	return db.User{}
}

func alert(r *http.Request) string {
	return session.PopString(r.Context(), "alert")
}

func notice(r *http.Request) string {
	return session.PopString(r.Context(), "notice")
}
