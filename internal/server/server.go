package server

import (
	"net/http"

	"github.com/ad9311/hitomgr/internal/cfg"
	"github.com/alexedwards/scs/v2"
)

var config *cfg.Config
var session *scs.Session

// Init ...
func Init(conf *cfg.Config, ssn *scs.SessionManager) {
	config = conf
	session = ssn
}

// New returns a new server with the loaded configuration.
func New() *http.Server {
	return &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: routes(),
	}
}
