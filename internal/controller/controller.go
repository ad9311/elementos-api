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
	App.Landmark = &db.Landmark{
		ID:          1,
		Name:        "San Felipe Castle",
		NativeName:  "Castillo San Felipe",
		Class:       "Fortification",
		Description: "Something...",
		WikiURL:     "https://wikipedia.org",
		Location:    []string{"Colombia", "Cartagena"},
		ImgURLs:     []string{"img1.png", "img2.png"},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   "ad9311",
	}
	App.Landmarks = []*db.Landmark{App.Landmark}
}
