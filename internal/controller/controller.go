package controller

import (
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/sess"
)

// App ...
var App *sess.App
var database *db.Database

// Init ...
func Init(dtbs *db.Database, sessionData *sess.App) {
	database = dtbs
	App = sessionData
}
