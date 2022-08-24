package val

import (
	"net/http"
	"strings"

	"github.com/ad9311/hitomgr/internal/db"
)

// ValidateNewLandmark ...
func ValidateNewLandmark(dtbs *db.Database, r *http.Request, user *db.User) (*db.Landmark, error) {
	var landmark *db.Landmark

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

	_, err := dtbs.InsertLandmark(r, user.ID, strMap)
	if err != nil {
		return landmark, err
	}

	lm, err := dtbs.SelectLandmarkByName(r.PostFormValue("name"))
	if err != nil {
		return landmark, err
	}
	landmark = lm
	landmark.CreatedBy = "jummm"

	return landmark, nil
}
