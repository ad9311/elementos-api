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
		landmarks, err := database.SelectLandmarks()
		if err != nil {
			fmt.Println(err)
		}
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		appMap["Landmarks"] = landmarks
		appMap["Alert"] = alert(r)
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
		appMap["Alert"] = alert(r)
		if err := render.WriteView(w, "landmarks_new", appMap); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostNewLandmark ...
func PostNewLandmark(w http.ResponseWriter, r *http.Request) {
	user := currentUser(r)
	landmark, err := val.ValidateNewLandmark(database, r, user.ID)
	if err != nil {
		fmt.Println(err)
		session.Put(r.Context(), "alert", err.Error())
		http.Redirect(w, r, "/landmarks/new", http.StatusSeeOther)
	} else {
		path := fmt.Sprintf("/landmarks/%d", landmark.ID)
		http.Redirect(w, r, path, http.StatusSeeOther)
	}
}

// GetShowLandmark ...
func GetShowLandmark(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		landmark, err := val.ValidateShowLandmark(database, r.URL.String())
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		appMap["Landmark"] = landmark
		appMap["Alert"] = alert(r)
		if err := render.WriteView(w, "landmarks_show", appMap); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// GetEditLandmark ...
func GetEditLandmark(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		landmark, err := val.ValidateShowLandmark(database, r.URL.String())
		if err != nil {
			fmt.Println(err)
		}
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		appMap["Landmark"] = landmark
		appMap["Alert"] = alert(r)
		if err := render.WriteView(w, "landmarks_edit", appMap); err != nil {
			fmt.Println(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostEditLandmark ...
func PostEditLandmark(w http.ResponseWriter, r *http.Request) {
	err := val.ValidateEditLandmark(database, r)
	if err != nil {
		fmt.Println(err)
		session.Put(r.Context(), "alert", err.Error())
		http.Redirect(w, r, r.URL.String(), http.StatusSeeOther)
	} else {
		path := fmt.Sprintf("/landmarks/%s", r.PostFormValue("landmark_id"))
		http.Redirect(w, r, path, http.StatusSeeOther)
	}
}

// PostDeleteLandmark ...
func PostDeleteLandmark(w http.ResponseWriter, r *http.Request) {
	err := val.ValidateDeleteLandmark(database, r)
	if err != nil {
		fmt.Println(err)
		session.Put(r.Context(), "alert", err.Error())
	}
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
