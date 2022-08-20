package server

import (
	"net/http"

	"github.com/ad9311/elementos_mgr/internal/cfg"
)

// App contains all current server's configuration.
type App struct {
	config *cfg.Config
}

var app App

// SetUp set ups the server with the loaded configuration
// and the database.
func SetUp(conf *cfg.Config) {
	app.config = conf
}

// New returns a new server with the loaded configuration.
func New() *http.Server {
	return &http.Server{
		Addr:    ":" + app.config.ServerPort,
		Handler: routes(),
	}
}
