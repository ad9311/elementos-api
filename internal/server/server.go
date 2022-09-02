package server

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
)

var serverPort string
var serverSecure bool
var session *scs.SessionManager

// SetUp ...
func SetUp(port string, secure bool) *scs.SessionManager {
	serverPort = port
	serverSecure = secure

	session = scs.New()
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = secure

	return session
}

// New ...
func New() *http.Server {
	return &http.Server{
		Addr:    ":" + serverPort,
		Handler: routes(),
	}
}
