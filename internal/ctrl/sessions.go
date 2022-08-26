package ctrl

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/render"
	"github.com/justinas/nosurf"
)

// GetSignIn ...
func GetSignIn(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		appMap := map[string]interface{}{}
		appMap["CSRFToken"] = nosurf.Token(r)
		if err := render.WriteView(w, "sessions_new", appMap); err != nil {
			fmt.Println(err)
		}
	}
}
