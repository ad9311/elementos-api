package ctrl

import (
	"fmt"
	"net/http"

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
		if err := render.WriteView(w, "registrations_new", appMap); err != nil {
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
