package apiv1

import (
	"fmt"
	"net/http"

	"github.com/ad9311/hitomgr/internal/db"
)

var database *db.Database

// Setup ...
func Setup(dbts *db.Database) {
	database = dbts
}

func writeResponse(w http.ResponseWriter, code int, res []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(res)
	if err != nil {
		fmt.Println(err)
	}
}
