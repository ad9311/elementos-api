package ctrl

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/val"
	"github.com/justinas/nosurf"
)

// GetDashboard ...
func GetDashboard(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		if err := render.WriteView(w, "landmarks_index", appMap); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// GetNewLandmark ...
func GetNewLandmark(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		if err := render.WriteView(w, "landmarks_new", appMap); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostNewLandmark ...
func PostNewLandmark(w http.ResponseWriter, r *http.Request) {
	lm, err := val.ValidateNewLandmark(database, r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/landmarks/new", http.StatusSeeOther)
	} else {
		path := fmt.Sprintf("/landmarks/%d", lm.ID)
		http.Redirect(w, r, path, http.StatusSeeOther)
	}
}

// GetShowLandmark ...
func GetShowLandmark(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		lm, err := val.ValidateShowLandmark(database, r.URL.String())
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
		appMap := make(map[string]interface{})
		appMap["Landmark"] = lm
		if err := render.WriteView(w, "landmarks_show", appMap); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}
