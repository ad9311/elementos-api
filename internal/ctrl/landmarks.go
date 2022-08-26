package ctrl

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/render"
	"github.com/justinas/nosurf"
)

// GetDashboard ...
func GetDashboard(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		appMap := map[string]interface{}{}
		appMap["CSRFToken"] = nosurf.Token(r)
		if err := render.WriteView(w, "landmarks_index", appMap); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}
