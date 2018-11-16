package app

import (
	"grapi/config"
	"grapi/core"
	"grapi/router"

	"github.com/gorilla/mux"
)

// App : app interface
type App struct {
	Database Database
	Cache    Cache
	Server   Server
}

// Database : database interface
type Database interface {
	Connect(*core.Config)
	Register(*core.Handlers)
}

// Cache : cache interface
type Cache interface {
	Connect(*core.Config)
	Register(*core.Handlers)
}

// Server : output interface
type Server interface {
	Start(*mux.Router, *core.Config)
}

// Run : run the app
func (app *App) Run() {
	config := config.GetConfig()

	handlers := new(core.Handlers)
	app.Database.Connect(config)
	app.Database.Register(handlers)
	app.Cache.Connect(config)
	app.Cache.Register(handlers)

	routes := router.CreateRoutes(config)
	router := router.NewRouter(routes, *handlers, *config)
	app.Server.Start(router, config)

}
