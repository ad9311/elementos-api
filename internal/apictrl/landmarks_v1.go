package apictrl

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
)

// GetLandmarks ...
func GetLandmarks(w http.ResponseWriter, r *http.Request) {
	landmarks, err := database.SelectLandmarks()
	if err != nil {
		response := db.Response{Message: "INTERNAL ERROR"}
		res, _ := json.Marshal(response)
		writeResponse(w, http.StatusInternalServerError, res)
	} else {
		response := db.ResponseWithData{Message: "OK", Data: landmarks}
		res, _ := json.Marshal(response)
		writeResponse(w, http.StatusOK, res)
	}
}
