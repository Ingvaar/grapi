package main

import (
	a "grapi/app"
	c "grapi/cache"
	d "grapi/database"
	se "grapi/server"
)

func main() {
	app := a.App{
		Database: &d.SQL{},
		Cache:    &c.Redis{},
		Server:   &se.Server{},
	}
	app.Run()
}
