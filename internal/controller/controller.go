package controller

import (
	"time"

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
	App.Landmark = &db.Landmark{}
	App.Landmarks = []*db.Landmark{{
		ID:           1,
		Name:         "San Felipe Castle",
		NativeName:   "Castillo San Felipe",
		Class:        "Fortification",
		Description:  "Something...",
		StartingYear: "1608",
		EndingYear:   "1705",
		WikiURL:      "https://wikipedia.org",
		Location:     []string{"Colombia", "Cartagena"},
		ImgURLs:      []string{"...", "..."},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    "ad9311",
	}}
}
