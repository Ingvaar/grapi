package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func status(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()

	if err == nil {
		fmt.Fprintln(w, "Database connected")
	} else {
		fmt.Fprintf(w, "%s\n", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome !")
}

func createLine(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	tab_name := pathVars["table"]
	id := pathVars["id"]

	fmt.Fprintln(w, "createLine")
	fmt.Fprintf(w, "Table : %s\n", tab_name)
	fmt.Fprintf(w, "Line id : %s\n", id)
}

func deleteLine(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	tab_name := pathVars["table"]
	id := pathVars["id"]

	fmt.Fprintln(w, "deleteLine")
	fmt.Fprintf(w, "Table : %s\n", tab_name)
	fmt.Fprintf(w, "Line id : %s\n", id)
}
