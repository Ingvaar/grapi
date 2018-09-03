package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"

	"grapi/handlers"
)

// Route : struct for the routes config file
type Route struct {
	Name        string `json:"name"`
	Method      string `json:"method"`
	Pattern     string `json:"pattern"`
	HandlerFunc string `json:"handler"`
}

// Routes : declares the type of an array of Route
type Routes []Route

// Router : global var of the router
var Router *mux.Router

// NewRouter : creates a new router
func NewRouter() {
	Router = mux.NewRouter().StrictSlash(true)
	var handler http.Handler
	var routes = createRoutes()

	if routes != nil {
		for _, route := range routes {
			handler = handlers.HandlerFunc[route.HandlerFunc]
			handler = Logger(handler, route.Name)
			Router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	} else {
		log.Fatal("Error: Routes config file incorrect")
		os.Exit(1)
	}
}
