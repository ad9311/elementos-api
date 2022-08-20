package server

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/ad9311/elementos_mgr/internal/cfg"
)

type data struct {
	StringMap map[string]string
	CSRFToken string
}

type application struct {
	config        *cfg.Config
	database      *sql.DB
	sessionsData  *data
	templateCache map[string]*template.Template
}

var app application

// SetUp set ups the server with the loaded configuration
// and the database.
func SetUp(conf *cfg.Config, conn *sql.DB) {
	app.config = conf
	app.database = conn
}

// New returns a new server with the loaded configuration.
func New() *http.Server {
	return &http.Server{
		Addr:    ":" + app.config.ServerPort,
		Handler: routes(),
	}
}
