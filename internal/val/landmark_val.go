package val

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ad9311/hitomgr/internal/db"
)

// ValidateNewLandmark ...
func ValidateNewLandmark(dtbs *db.Database, r *http.Request, user *db.User) (db.Landmark, error) {
	var landmark db.Landmark

	params := []string{
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

	location := strings.Split(r.PostFormValue("location"), ",")
	imgURLs := strings.Split(r.PostFormValue("img-urls"), ",")
	strMap := map[string][]string{"img-urls": imgURLs, "location": location}

	err := dtbs.InsertLandmark(r, user.ID, strMap)
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
