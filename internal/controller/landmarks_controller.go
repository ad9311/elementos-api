package controller

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/render"
	"github.com/justinas/nosurf"
)

// GetDashboard ...
func GetDashboard(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		App.CSRFToken = nosurf.Token(r)
		if err := render.WriteView(w, "dashboard.view.html"); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "sign_in", http.StatusSeeOther)
	}
}
