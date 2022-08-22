package controller

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/render"
	"github.com/justinas/nosurf"
)

// GetSignIn ...
func GetSignIn(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		App.CSRFToken = nosurf.Token(r)
		if err := render.WriteView(w, "sign_in.view.html"); err != nil {
			fmt.Println(err)
		}
	}
}

// PostSignIn ...
func PostSignIn(w http.ResponseWriter, r *http.Request) {
	params := []string{"username", "password"}
	err := validateFormParams(r, params)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
		return
	}

	user, err := database.SelectUserByUsername(r)
	App.CurrentUser = user
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
		return
	}

	err = validatePassword(r.PostFormValue("password"), user.EncryptedPassword)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
		return
	}
	user.EncryptedPassword = ""

	err = database.UpdateUserLastLogin(user)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
		return
	}

	_ = App.Session.RenewToken(r.Context())
	App.Session.Put(r.Context(), "signedIn", true)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

// GetSignUp ...
func GetSignUp(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		App.CSRFToken = nosurf.Token(r)
		if err := render.WriteView(w, "sign_up.view.html"); err != nil {
			fmt.Println(err)
		}
	}
}

// PostSignUp ...
func PostSignUp(w http.ResponseWriter, r *http.Request) {
	err := validateSignUpForm(r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_up", http.StatusSeeOther)
		return
	}

	ic, err := database.SelectInvitationCode(r.PostFormValue("code"))
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_up", http.StatusSeeOther)
		return
	}

	err = validateDate(ic.Validity)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_up", http.StatusSeeOther)
		return
	}

	ep, err := App.EncryptPassword(r.PostFormValue("password"))
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_up", http.StatusSeeOther)
		return
	}

	err = database.InsertUser(r, ep)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/sign_up", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
}

// PostSignOut ...
func PostSignOut(w http.ResponseWriter, r *http.Request) {
	App.CurrentUser = nil
	_ = App.Session.Destroy(r.Context())
	_ = App.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
}
