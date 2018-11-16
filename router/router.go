package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	c "grapi/core"
	m "grapi/middlewares"
)

// NewRouter : creates a new router
func NewRouter(routes []c.Route, handlers c.Handlers, config c.Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	var handler http.Handler

	if routes != nil {
		for _, route := range routes {
			handler = m.ValidateMiddleware(config, route.Level, handlers[route.HandlerFunc])
			handler = m.Logger(handler, route.Name, config)
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(handler)
		}
	} else {
		log.Fatal("Error: Routes config file incorrect")
		os.Exit(1)
	}
	return (router)
}
