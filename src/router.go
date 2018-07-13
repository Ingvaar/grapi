package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"os"
)

func NewRouter(opt Options) *mux.Router {
	routeur := mux.NewRouter().StrictSlash(true);
	var handler http.Handler;
	var routes = create_routes(opt);

	if routes != nil {
		for _, route := range routes {
			handler = HandlerFunc[route.HandlerFunc];
			handler = Logger(handler, route.Name);
			routeur.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler);
		}
	} else {
		fmt.Println("Error: Config file incorrect\nExiting...");
		os.Exit(1);
	}
	return (routeur);
}
