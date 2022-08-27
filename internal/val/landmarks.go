package val

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
)

// ValidateNewLandmark ...
func ValidateNewLandmark(dtbs *db.Database, r *http.Request, userID int64) (db.Landmark, error) {
	params := []string{
		"user_id",
		"name",
		"native_name",
		"category",
		"description",
		"wiki_url",
		"location",
		"img_urls",
	}
	if err := checkFormParams(r, params); err != nil {
		return db.Landmark{}, err
	}

	formMap := formToMap(r, params)
	err := checkUserID(formMap["user_id"], userID)
	if err != nil {
		return db.Landmark{}, err
	}

	formMap["location"] = "{" + formMap["location"] + "}"
	formMap["img_urls"] = "{" + formMap["img_urls"] + "}"

	err = dtbs.InsertLandmark(formMap)
	if err != nil {
		return db.Landmark{}, err
	}

	landmark, err := dtbs.SelectLandmarkByName(formMap["name"])
	if err != nil {
		return db.Landmark{}, err
	}

	return landmark, nil
}

// ValidateShowLandmark ...
func ValidateShowLandmark(dtbs *db.Database, urlStr string) (db.Landmark, error) {
	i, err := retrieveIDFromURL(urlStr, "landmarks")
	if err != nil {
		return db.Landmark{}, err
	}

	lm, err := dtbs.SelectLandmarkByID(int64(i))
	if err != nil {
		return db.Landmark{}, err
	}

	return lm, nil
}

// ValidateEditLandmark ...
func ValidateEditLandmark(dtbs *db.Database, r *http.Request) error {
	id, err := retrieveIDFromURL(r.URL.String(), "landmarks")
	if err != nil {
		return err
	}

	params := []string{
		"landmark_id",
		"name",
		"native_name",
		"class",
		"description",
		"wiki_url",
		"location",
		"img_urls",
	}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	formMap := formToMap(r, params)
	formMap["location"] = "{" + formMap["location"] + "}"
	formMap["img_urls"] = "{" + formMap["img_urls"] + "}"

	if formMap["landmark_id"] != fmt.Sprintf("%d", id) {
		return fmt.Errorf("form error")
	}

	if err := dtbs.UpdateLandmarkByID(formMap); err != nil {
		return err
	}

	return nil
}

// ValidateDeleteLandmark ...
func ValidateDeleteLandmark(dtbs *db.Database, r *http.Request) error {
	id, err := retrieveIDFromURL(r.URL.String(), "landmarks")
	if err != nil {
		return err
	}

	params := []string{"landmark_id"}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	if r.PostFormValue("landmark_id") != fmt.Sprintf("%d", id) {
		return fmt.Errorf("form error")
	}

	if err := dtbs.DeleteLandmarkByID(id); err != nil {
		return err
	}

	return nil
}
