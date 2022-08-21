package server

import (
	"html/template"
	"net/http"

	"github.com/ad9311/hitomgr/internal/cfg"
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/alexedwards/scs/v2"
)

type data struct {
	CurrentUser *db.User
	CSRFToken   string
	StringMap   map[string]string
}

type application struct {
	config     *cfg.Config
	database   *db.Database
	session    *scs.SessionManager
	viewsCache map[string]*template.Template
	Data       data
}

var app application

// SetUp set ups the server with the loaded configuration
// and the database.
func SetUp(conf *cfg.Config, dtbs *db.Database) error {
	app.config = conf
	app.database = dtbs
	viewsCache, err := deafultViewsCache()
	if err != nil {
		return err
	}
	app.viewsCache = viewsCache
	app.session = scs.New()
	app.session.Cookie.Persist = true
	app.session.Cookie.SameSite = http.SameSiteLaxMode
	app.session.Cookie.Secure = app.config.SeverSecure

	return nil
}

// New returns a new server with the loaded configuration.
func New() *http.Server {
	return &http.Server{
		Addr:    ":" + app.config.ServerPort,
		Handler: routes(),
	}
}
