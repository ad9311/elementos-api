package val

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ad9311/hitomgr/internal/db"
)

// ValidateNewLandmark ...
func ValidateNewLandmark(dtbs *db.Database, r *http.Request) (db.Landmark, error) {
	var landmark db.Landmark

	params := []string{
		"user-id",
		"name",
		"native-name",
		"class",
		"description",
		"wiki-url",
		"location",
		"img-urls",
	}
	if err := checkFormParams(r, params); err != nil {
		return landmark, err
	}

	formMap := formToMap(r, params)
	location := strings.Split(r.PostFormValue("location"), ",")
	imgURLs := strings.Split(r.PostFormValue("img-urls"), ",")
	formMap["location"] = location
	formMap["img-urls"] = imgURLs

	err := dtbs.InsertLandmark(r, formMap)
	if err != nil {
		return landmark, err
	}

	lm, err := dtbs.SelectLandmarkByName(r.PostFormValue("name"))
	if err != nil {
		return landmark, err
	}

	return lm, nil
}

// ValidateShowLandmark ...
func ValidateShowLandmark(dtbs *db.Database, urlStr string) (db.Landmark, error) {
	landmark := db.Landmark{}
	url, err := url.Parse(urlStr)
	if err != nil {
		return landmark, err
	}

	urlSlice := strings.Split(url.Path, "/")
	id := urlSlice[len(urlSlice)-1]
	i, err := strconv.Atoi(id)
	if err != nil {
		return landmark, err
	}

	lm, err := dtbs.SelectLandmarkByID(int64(i))
	if err != nil {
		return landmark, err
	}

	return lm, nil
}

// ValidateEditLandmark ...
func ValidateEditLandmark(dtbs *db.Database, r *http.Request) error {
	params := []string{
		"user-id",
		"landmark-id",
		"name",
		"native-name",
		"class",
		"description",
		"wiki-url",
		"location",
		"img-urls",
	}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	formMap := formToMap(r, params)
	location := strings.Split(r.PostFormValue("location"), ",")
	imgURLs := strings.Split(r.PostFormValue("img-urls"), ",")
	formMap["location"] = location
	formMap["img-urls"] = imgURLs

	if err := dtbs.UpdateLandmarkByID(formMap); err != nil {
		return err
	}

	return nil
}

// ValidateDeleteLandmark ...
func ValidateDeleteLandmark(dtbs *db.Database, r *http.Request) error {
	params := []string{
		"user-id",
		"landmark-id",
	}
	if err := checkFormParams(r, params); err != nil {
		return err
	}

	if err := dtbs.DeleteLandmarkByID(r.PostFormValue("landmark-id")); err != nil {
		return err
	}

	return nil
}
