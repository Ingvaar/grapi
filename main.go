package main

import (
	a "grapi/app"
	r "grapi/redis"
	s "grapi/sql"
	se "grapi/server"
)

func main() {
	app := a.App{
		SQL:	&s.Database{},
		Redis:	&r.Database{},
		Server: &se.Server{},
	}
	app.Run()
}
