package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// NewRouter : creates a new router
func NewRouter(opt Options) *mux.Router {
	routeur := mux.NewRouter().StrictSlash(true)
	var handler http.Handler
	var routes = createRoutes(opt)

	if routes != nil {
		for _, route := range routes {
			handler = HandlerFunc[route.HandlerFunc]
			handler = Logger(handler, route.Name)
			routeur.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	} else {
		log.Fatal("Error: Routes config file incorrect")
		os.Exit(1)
	}
	return (routeur)
}
