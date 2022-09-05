package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/hitomgr/internal/cnsl"
	"github.com/ad9311/hitomgr/internal/db"
	"github.com/ad9311/hitomgr/internal/val"
)

// GetLandmarks ...
func GetLandmarks(w http.ResponseWriter, r *http.Request) {
	landmarks, err := val.ValidateGetLandmarks(database, r)
	if err != nil {
		cnsl.Error(err)
		response := db.Response{Message: "INTERNAL ERROR"}
		res, _ := json.Marshal(response)
		writeResponse(w, http.StatusInternalServerError, res)
	} else {
		response := db.ResponseWithData{Message: "OK", Data: landmarks}
		res, _ := json.Marshal(response)
		writeResponse(w, http.StatusOK, res)
	}
}
