package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/val"
	"github.com/justinas/nosurf"
)

// GetSignIn ...
func GetSignIn(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		App.CSRFToken = nosurf.Token(r)
		App.URL = r.URL.String()
		if err := render.WriteView(w, "users_sign_in"); err != nil {
			fmt.Println(err)
		}
	}
}

// PostSignIn ...
func PostSignIn(w http.ResponseWriter, r *http.Request) {
	user, err := val.ValidateUserSignIn(database, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	} else {
		App.CurrentUser = &user
		App.CurrentUser.HashedPassword = ""
		if err = App.Session.RenewToken(r.Context()); err != nil {
			fmt.Println(err)
		}
		App.CurrentUser.SignedIn = true
		App.CurrentUser.LastLogin = time.Now()
		App.Session.Put(r.Context(), "signedIn", true)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

// GetSignUp ...
func GetSignUp(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		App.CSRFToken = nosurf.Token(r)
		App.URL = r.URL.String()
		if err := render.WriteView(w, "users_sign_up"); err != nil {
			fmt.Println(err)
		}
	}
}

// PostSignUp ...
func PostSignUp(w http.ResponseWriter, r *http.Request) {
	err := val.ValidateUserSignUp(database, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_up", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostSignOut ...
func PostSignOut(w http.ResponseWriter, r *http.Request) {
	App.CurrentUser = &db.User{}
	_ = App.Session.Destroy(r.Context())
	_ = App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
}
