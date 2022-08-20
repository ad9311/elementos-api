package server

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

const (
	signInTemplate    = "sign_in.template.html"
	signUpTemplate    = "sign_up.template.html"
	dashboardTemplate = "dashboard.template.html"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	if userSignedIn(r) {
		http.Redirect(w, r, dashboard, http.StatusSeeOther)
	} else {
		http.Redirect(w, r, signIn, http.StatusSeeOther)
	}
}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	if userSignedIn(r) {
		app.Data.CSRFToken = nosurf.Token(r)
		if err := writeTemplate(w, dashboardTemplate); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, signIn, http.StatusSeeOther)
	}
}

func getSignIn(w http.ResponseWriter, r *http.Request) {
	if userSignedIn(r) {
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

	user, err := app.database.SelectUser(r)
	app.Data.CurrentUser = user
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signIn, http.StatusSeeOther)
		return
	}

	err = validatePassword(r.PostFormValue("password"), user.EncryptedPassword)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signIn, http.StatusSeeOther)
		return
	}
	user.EncryptedPassword = ""

	err = app.database.UpdateLastLogin(user)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signIn, http.StatusSeeOther)
		return
	}

	_ = app.session.RenewToken(r.Context())
	app.session.Put(r.Context(), "signedIn", true)
	http.Redirect(w, r, dashboard, http.StatusSeeOther)
}

func getSignUp(w http.ResponseWriter, r *http.Request) {
	if userSignedIn(r) {
		http.Redirect(w, r, dashboard, http.StatusSeeOther)
	} else {
		app.Data.CSRFToken = nosurf.Token(r)
		if err := writeTemplate(w, signUpTemplate); err != nil {
			fmt.Println(err)
		}
	}
}

func postSignUp(w http.ResponseWriter, r *http.Request) {
	fields := []string{"first_name", "last_name", "username", "email", "password"}
	err := validateForm(r, fields)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signUp, http.StatusSeeOther)
		return
	}

	err = confirmPasswords(
		r.PostFormValue("password"),
		r.PostFormValue("password-confirmation"),
	)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signUp, http.StatusSeeOther)
		return
	}

	ep, err := encryptPassword(r.PostFormValue("password"))
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signUp, http.StatusSeeOther)
		return
	}

	err = app.database.InsertUser(r, ep)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, signUp, http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, signIn, http.StatusSeeOther)
}

func postSignOut(w http.ResponseWriter, r *http.Request) {
	_ = app.session.Destroy(r.Context())
	_ = app.session.RenewToken(r.Context())
	http.Redirect(w, r, signIn, http.StatusSeeOther)
}
