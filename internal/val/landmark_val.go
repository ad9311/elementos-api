package val

import (
	"net/http"
	"strings"

	"github.com/ad9311/hitomgr/internal/db"
)

// ValidateNewLandmark ...
func ValidateNewLandmark(dtbs *db.Database, r *http.Request, user *db.User) (*db.Landmark, error) {
	var lm *db.Landmark

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
		return lm, err
	}

	location := strings.Split(r.PostFormValue("location"), ",")
	imgURLs := strings.Split(r.PostFormValue("img-urls"), ",")
	strMap := map[string][]string{"img-urls": imgURLs, "location": location}

	if err := dtbs.InsertLandmark(r, user.ID, strMap); err != nil {
		return lm, err
	}

	return lm, nil
}
