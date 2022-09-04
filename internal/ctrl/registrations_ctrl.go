package ctrl

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/cnsl"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/val"
	"github.com/justinas/nosurf"
)

// GetSignUp ...
func GetSignUp(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["Alert"] = alert(r)
		appMap["Notice"] = notice(r)
		if err := render.WriteView(w, "registrations_new", appMap); err != nil {
			cnsl.Error(err)
		}
	}
}

// PostSignUp ...
func PostSignUp(w http.ResponseWriter, r *http.Request) {
	err := val.ValidateUserSignUp(database, r)
	if err != nil {
		cnsl.Log(err)
		session.Put(r.Context(), "alert", err.Error())
		http.Redirect(w, r, "/sign_up", http.StatusSeeOther)
	} else {
		notif := fmt.Sprintf("user %s created successfully", r.PostFormValue("username"))
		session.Put(r.Context(), "notice", notif)
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}
