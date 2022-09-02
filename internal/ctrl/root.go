package ctrl

import (
	"net/http"
)

// GetRoot ...
func GetRoot(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}
