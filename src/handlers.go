package main

import (
	"strconv"
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

func getLine(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r)
	tab_name := pathVars["table"]
	id := pathVars["id"]
	id_num, err_atoi := strconv.Atoi(id)

	statement := fmt.Sprintf("SELECT * FROM %s WHERE id=%d", tab_name, id_num)
	content, err := db.Exec(statement)
	if err != nil || err_atoi != nil {
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
		}
		if err_atoi != nil {
			fmt.Fprintf(w, "Error: invalid id '%s'\n", id)
		}
	} else {
		fmt.Fprintf(w, "Table : %s\n", tab_name)
		fmt.Fprintf(w, "Line id : %d\n", id_num)
		fmt.Fprintf(w, "Content : %s\n", content)
	}
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
