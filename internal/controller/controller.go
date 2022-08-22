package controller

import (
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/sess"
)

// Data ...
var Data *sess.Data
var database *db.Database

// Init ...
func Init(dtbs *db.Database, sessionData *sess.Data) {
	database = dtbs
	Data = sessionData
}
