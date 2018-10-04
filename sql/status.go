package sql

import (
	"fmt"
	"net/http"
)

// Status : print the status of the database as plain text
func (db *Database) Status(w http.ResponseWriter, r *http.Request) {
	sqlErr := db.DB.Ping()

	if sqlErr == nil {
		fmt.Fprintln(w, "SQL Database connected")
	} else if sqlErr != nil && db.config.UseSQL == 1 {
		fmt.Fprintln(w, sqlErr)
	}
}
