package nosql

import (
	"net/http"
)

// Update :
func Update(w http.ResponseWriter, r *http.Request) {
	Create(w, r)
}
