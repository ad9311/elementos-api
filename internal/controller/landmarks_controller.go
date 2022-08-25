package controller

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/val"
	"github.com/justinas/nosurf"
)

// GetDashboard ...
func GetDashboard(w http.ResponseWriter, r *http.Request) {
	if App.IsUserSignedIn(r) {
		App.CSRFToken = nosurf.Token(r)
		App.URL = r.URL.String()
		lms, err := database.SelectLandmarks()
		if err != nil {
			fmt.Println(err)
		}
		App.Landmarks = lms
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
		lm, err := val.ValidateShowLandmark(database, App.URL)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
		App.Landmark = &lm
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
	if r.PostFormValue("user-id") == fmt.Sprintf("%d", App.CurrentUser.ID) {
		lm, err := val.ValidateNewLandmark(database, r)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/landmarks/new", http.StatusSeeOther)
		} else {
			App.Landmark = &lm
			path := fmt.Sprintf("/landmarks/%d", App.Landmark.ID)
			http.Redirect(w, r, path, http.StatusSeeOther)
		}
	} else {
		fmt.Println(fmt.Errorf("incorrect user"))
		http.Redirect(w, r, "/landmarks/new", http.StatusSeeOther)
	}
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
	if r.PostFormValue("user-id") == fmt.Sprintf("%d", App.CurrentUser.ID) {
		err := val.ValidateEditLandmark(database, r)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
		} else {
			path := fmt.Sprintf("/landmarks/%s", r.PostFormValue("landmark-id"))
			http.Redirect(w, r, path, http.StatusSeeOther)
		}
	} else {
		fmt.Println(fmt.Errorf("incorrect user"))
		http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
	}
}

// PostDeleteLandmark ...
func PostDeleteLandmark(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("user-id") == fmt.Sprintf("%d", App.CurrentUser.ID) {
		err := val.ValidateDeleteLandmark(database, r)
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		fmt.Println(fmt.Errorf("incorrect user"))
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}
