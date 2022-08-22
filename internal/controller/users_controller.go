package controller

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/render"
	"github.com/justinas/nosurf"
)

// GetSignIn ...
func GetSignIn(w http.ResponseWriter, r *http.Request) {
	if Data.Session.GetBool(r.Context(), "signedIn") {
		http.Redirect(w, r, "/dasboard", http.StatusSeeOther)
	} else {
		Data.CSRFToken = nosurf.Token(r)
		if err := render.WriteView(w, "sign_in.view.html"); err != nil {
			fmt.Println(err)
		}
	}
}
