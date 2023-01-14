package apiv1

import (
	"encoding/json"
	"net/http"

	"github.com/ad9311/hitomgr/internal/cnsl"
	"github.com/ad9311/hitomgr/internal/db"
)

// GetCategories ...
func GetCategories(w http.ResponseWriter, r *http.Request) {
	categogies, err := database.SelectCategories()
	if err != nil {
		cnsl.Error(err)
		response := db.Response{Message: "INTERNAL ERROR"}
		res, _ := json.Marshal(response)
		writeResponse(w, http.StatusInternalServerError, res)
	} else {
		response := db.ResponseWithData{Message: "OK", Data: categogies}
		res, _ := json.Marshal(response)
		writeResponse(w, http.StatusOK, res)
	}
}
