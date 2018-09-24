package main

import (
	c "grapi/config"
	"grapi/db"
	r "grapi/router"
	s "grapi/server"
)

func main() {
	c.ParsCmdline()
	c.GetConfig()
	db.OpenSQL()
	db.OpenNoSQL()
	r.NewRouter()

	defer db.SQL.Close()
	defer db.Nosql.Close()
	s.StartServer()
}
