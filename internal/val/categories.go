package val

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
)

// ValidateNewCategory ...
func ValidateNewCategory(dtbs *db.Database, r *http.Request) error {
	params := []string{"name"}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	formMap := formToMap(r, params)

	err := dtbs.InsertCategory(formMap)
	if err != nil {
		return err
	}

	return nil
}

// ValidateShowCategory ...
func ValidateShowCategory(dtbs *db.Database, urlStr string) (db.Category, error) {
	i, err := retrieveIDFromURL(urlStr, "categories")
	if err != nil {
		return db.Category{}, err
	}

	cat, err := dtbs.SelectCategory(i)
	if err != nil {
		return db.Category{}, err
	}

	return cat, nil
}

// ValidateEditCategory ...
func ValidateEditCategory(dtbs *db.Database, r *http.Request) error {
	id, err := retrieveIDFromURL(r.URL.String(), "categories")
	if err != nil {
		return err
	}

	params := []string{"name", "category_id"}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	formMap := formToMap(r, params)

	if formMap["category_id"] != fmt.Sprintf("%d", id) {
		return fmt.Errorf(
			"ids %d and %s do not match",
			id,
			formMap["category_id"],
		)
	}

	if err := dtbs.UpdateCategory(formMap); err != nil {
		return err
	}

	return nil
}

// ValidateDeleteCategory ...
func ValidateDeleteCategory(dtbs *db.Database, r *http.Request) error {
	id, err := retrieveIDFromURL(r.URL.String(), "categories")
	if err != nil {
		return err
	}

	params := []string{"category_id"}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	if r.PostFormValue("category_id") != fmt.Sprintf("%d", id) {
		return fmt.Errorf(
			"ids %d and %s do not match",
			id,
			r.PostFormValue("category_id"),
		)
	}

	if err := dtbs.DeleteCategory(id); err != nil {
		return err
	}

	return err
}
