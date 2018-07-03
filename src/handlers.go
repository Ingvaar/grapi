package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome !");
}

func getTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	tab_name := vars["table"];
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
	fmt.Fprintln(w, "getTable");
	fmt.Fprintf(w, "Table : %s\n", tab_name);
}

func getLine(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r);
	tab_name := pathVars["table"];
	id := pathVars["id"];

	fmt.Fprintln(w, "getLineContent");
	fmt.Fprintf(w, "Table : %s\n", tab_name);
	fmt.Fprintf(w, "Line id : %s\n", id);
}

func createLine(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r);
	tab_name := pathVars["table"];
	id := pathVars["id"];

	fmt.Fprintln(w, "createLine");
	fmt.Fprintf(w, "Table : %s\n", tab_name);
	fmt.Fprintf(w, "Line id : %s\n", id);
}

func deleteLine(w http.ResponseWriter, r *http.Request) {
	pathVars := mux.Vars(r);
	tab_name := pathVars["table"];
	id := pathVars["id"];

	fmt.Fprintln(w, "deleteLine");
	fmt.Fprintf(w, "Table : %s\n", tab_name);
	fmt.Fprintf(w, "Line id : %s\n", id);
}
