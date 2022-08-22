package controller

import "net/http"

// GetRoot ...
func GetRoot(w http.ResponseWriter, r *http.Request) {
	if Data.Session.GetBool(r.Context(), "signedIn") {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "sign_in", http.StatusSeeOther)
	}
}
