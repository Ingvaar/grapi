package nosql

import (
	"net/http"

	"github.com/gorilla/mux"

	"grapi/db"
	"grapi/utils"
)

// Exists : do an exists on the id passed in the url
// and return an http response code
func Exists(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["type"] + ":" + mux.Vars(r)["id"]

	reply, err := db.Nosql.Cmd("EXISTS", id).Int()
	if err != nil {
		utils.SendResponse(w, err, http.StatusBadRequest)
	} else if reply == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
