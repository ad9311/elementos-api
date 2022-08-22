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
		App.URL = r.URL.String()
		if err := render.WriteView(w, "landmarks_index"); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// GetShowLandmark ...
func GetShowLandmark(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		App.URL = r.URL.String()
		if err := render.WriteView(w, "landmarks_show"); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// GetNewLandmark ...
func GetNewLandmark(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		App.CSRFToken = nosurf.Token(r)
		App.URL = r.URL.String()
		if err := render.WriteView(w, "landmarks_new"); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostNewLandmark ...
func PostNewLandmark(w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprintf("/landmarks/%d", App.Landmark.ID)
	http.Redirect(w, r, path, http.StatusSeeOther)
}

// GetEditLandmark ...
func GetEditLandmark(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		App.CSRFToken = nosurf.Token(r)
		App.URL = r.URL.String()
		if err := render.WriteView(w, "landmarks_edit"); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostEditLandmark ...
func PostEditLandmark(w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprintf("/landmarks/%d", App.Landmark.ID)
	http.Redirect(w, r, path, http.StatusSeeOther)
}
