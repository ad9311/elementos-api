package server

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

const (
	loginTemplate     = "login.template.html"
	dashboardTemplate = "dashboard.template.html"
)

func getDashboard(w http.ResponseWriter, r *http.Request) {
	if userLoggedIn(r) {
		app.CSRFToken = nosurf.Token(r)
		if err := writeTemplate(w, dashboardTemplate); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, login, http.StatusSeeOther)
	}
}

func getLogin(w http.ResponseWriter, r *http.Request) {
	if userLoggedIn(r) {
		http.Redirect(w, r, dashboard, http.StatusSeeOther)
	} else {
		app.CSRFToken = nosurf.Token(r)
		if err := writeTemplate(w, loginTemplate); err != nil {
			fmt.Println(err)
		}
	}
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	_ = app.session.RenewToken(r.Context())
	app.session.Put(r.Context(), "login", true)
	http.Redirect(w, r, dashboard, http.StatusSeeOther)
}

func postLogout(w http.ResponseWriter, r *http.Request) {
	_ = app.session.Destroy(r.Context())
	_ = app.session.RenewToken(r.Context())
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
