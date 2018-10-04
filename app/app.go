package app

import (
	"grapi/config"
	"grapi/core"
	"grapi/router"

	"github.com/gorilla/mux"
)

// App : app interface
type App struct {
	SQL    SQL
	Redis  NoSQL
	Server Server
}

// SQL : sql interface
type SQL interface {
	Connect(*core.Config)
	Register(*core.Handlers)
}

// NoSQL : nosql interface
type NoSQL interface {
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
	app.SQL.Connect(config)
	app.SQL.Register(handlers)
	app.Redis.Connect(config)
	app.Redis.Register(handlers)

	routes := router.CreateRoutes(config)
	router := router.NewRouter(routes, *handlers, *config)
	app.Server.Start(router, config)

}
