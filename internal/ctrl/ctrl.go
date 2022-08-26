package ctrl

import (
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var database *db.Database

// SetUp ...
func SetUp(dtbs *db.Database, sess *scs.SessionManager) {
	database = dtbs
	session = sess
}
