package server

import (
	"database/sql"
	"net/http"

	"github.com/ad9311/elementos_mgr/internal/cfg"
)

type application struct {
	config   *cfg.Config
	database *sql.DB
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
