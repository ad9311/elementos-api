package ctrl

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/cnsl"
	"github.com/ad9311/hitomgr/internal/render"
	"github.com/ad9311/hitomgr/internal/val"
	"github.com/justinas/nosurf"
)

// GetCategories ...
func GetCategories(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		categories, err := database.SelectCategories()
		if err != nil {
			cnsl.Log(err)
		}
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		appMap["Categories"] = categories
		appMap["Alert"] = alert(r)
		appMap["Notice"] = notice(r)
		if err := render.WriteView(w, "categories_index", appMap); err != nil {
			cnsl.Error(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// GetNewCategory ...
func GetNewCategory(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		appMap["Alert"] = alert(r)
		if err := render.WriteView(w, "categories_new", appMap); err != nil {
			cnsl.Error(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostCategory ...
func PostCategory(w http.ResponseWriter, r *http.Request) {
	err := val.ValidateNewCategory(database, r)
	if err != nil {
		cnsl.Log(err)
		session.Put(r.Context(), "alert", err.Error())
		http.Redirect(w, r, "/categories/new", http.StatusSeeOther)
	} else {
		notif := fmt.Sprintf("category %s created successfully", r.PostFormValue("name"))
		session.Put(r.Context(), "notice", notif)
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

// GetEditCategory ...
func GetEditCategory(w http.ResponseWriter, r *http.Request) {
	if session.GetBool(r.Context(), "user_signed_in") {
		category, err := val.ValidateShowCategory(database, r.URL.String())
		if err != nil {
			cnsl.Log(err)
		}
		appMap := make(map[string]interface{})
		appMap["CSRFToken"] = nosurf.Token(r)
		appMap["CurrentUser"] = currentUser(r)
		appMap["Category"] = category
		appMap["Alert"] = alert(r)
		if err := render.WriteView(w, "categories_edit", appMap); err != nil {
			cnsl.Error(err)
		}
	} else {
		http.Redirect(w, r, "/sign_in", http.StatusSeeOther)
	}
}

// PostEditCategory ...
func PostEditCategory(w http.ResponseWriter, r *http.Request) {
	err := val.ValidateEditCategory(database, r)
	if err != nil {
		cnsl.Log(err)
		session.Put(r.Context(), "alert", err.Error())
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	} else {
		notif := fmt.Sprintf("category %s updated successfully", r.PostFormValue("name"))
		session.Put(r.Context(), "notice", notif)
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

// PostDeleteCategory ...
func PostDeleteCategory(w http.ResponseWriter, r *http.Request) {
	err := val.ValidateDeleteCategory(database, r)
	if err != nil {
		cnsl.Log(err)
		session.Put(r.Context(), "alert", err.Error())
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	} else {
		notif := fmt.Sprintf("category %s deleted", r.PostFormValue("name"))
		session.Put(r.Context(), "notice", notif)
		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}
