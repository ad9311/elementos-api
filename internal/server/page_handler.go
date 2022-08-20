package server

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

const (
	signInTemplate    = "sign_in.template.html"
	dashboardTemplate = "dashboard.template.html"
)

func getDashboard(w http.ResponseWriter, r *http.Request) {
	if userLoggedIn(r) {
		app.Data.CSRFToken = nosurf.Token(r)
		if err := writeTemplate(w, dashboardTemplate); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, signIn, http.StatusSeeOther)
	}
}

func getSignIn(w http.ResponseWriter, r *http.Request) {
	if userLoggedIn(r) {
		http.Redirect(w, r, dashboard, http.StatusSeeOther)
	} else {
		app.Data.CSRFToken = nosurf.Token(r)
		if err := writeTemplate(w, signInTemplate); err != nil {
			fmt.Println(err)
		}
	}
}

func postSignIn(w http.ResponseWriter, r *http.Request) {
	fields := []string{"username", "password"}
	err := validateForm(r, fields)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signIn, http.StatusSeeOther)
		return
	}

	user, err := app.database.GetUser(r)
	app.Data.CurrentUser = user
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signIn, http.StatusSeeOther)
		return
	}

	err = validatePassword(r, user.EncryptedPassword)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signIn, http.StatusSeeOther)
		return
	}

	user.EncryptedPassword = ""
	_ = app.session.RenewToken(r.Context())
	app.session.Put(r.Context(), "signedIn", true)
	http.Redirect(w, r, dashboard, http.StatusSeeOther)
}

func postSignOut(w http.ResponseWriter, r *http.Request) {
	_ = app.session.Destroy(r.Context())
	_ = app.session.RenewToken(r.Context())
	http.Redirect(w, r, signIn, http.StatusSeeOther)
}
