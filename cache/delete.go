package redis

import (
	"net/http"

	"github.com/gorilla/mux"

	"grapi/utils"
)

// Delete : delete the passed id
func (rd *Redis) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["type"] + ":" + mux.Vars(r)["id"]

	reply, err := rd.RC.Cmd("DEL", id).Int()
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
	} else if reply == 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
