package server

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/ad9311/elementos_mgr/internal/cfg"
	"github.com/alexedwards/scs/v2"
)

type data struct {
}

type application struct {
	config        *cfg.Config
	database      *sql.DB
	session       *scs.SessionManager
	templateCache map[string]*template.Template
	StringMap     map[string]string
	CSRFToken     string
	sessionsData  *data
	templateCache map[string]*template.Template
}

var app application

// SetUp set ups the server with the loaded configuration
// and the database.
func SetUp(conf *cfg.Config, conn *sql.DB) error {
	app.config = conf
	app.database = conn
	templateCache, err := defaultTemplateCache()
	if err != nil {
		return err
	}
	app.templateCache = templateCache
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
