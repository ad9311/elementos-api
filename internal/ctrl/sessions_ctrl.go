package ctrl

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/cnsl"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/val"
	"github.com/justinas/nosurf"
)

// GetSignIn ...
func GetSignIn(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["Alert"] = alert(r)
		appMap["Notice"] = notice(r)
		if err := render.WriteView(w, "sessions_new", appMap); err != nil {
			cnsl.Error(err)
		}
	}
}

// PostSignIn ...
func PostSignIn(w http.ResponseWriter, r *http.Request) {
	user, err := val.ValidateUserSignIn(database, r)
	if err != nil {
		cnsl.Log(err)
		session.Put(r.Context(), "alert", err.Error())
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	} else {
		session.Put(r.Context(), "user_signed_in", true)
		session.Put(r.Context(), "current_user", user)
		cnsl.Log(fmt.Sprintf("user %s has logged in", user.Username))
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

// PostSignOut ...
func PostSignOut(w http.ResponseWriter, r *http.Request) {
	user := currentUser(r)
	session.Destroy(r.Context())
	session.RenewToken(r.Context())
	session.Put(r.Context(), "notice", "sign out successfully")
	cnsl.Log(fmt.Sprintf("user %s has logged out", user.Username))
	http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
}
