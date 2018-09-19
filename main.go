package main

import (
	"log"
	"net/http"

	c "grapi/config"
	"grapi/db"
	r "grapi/router"
)

func main() {
	c.ParsCmdline()
	c.GetConfig()
	db.OpenSQL()
	db.OpenNoSQL()
	r.NewRouter()

	defer db.SQL.Close()
	defer db.Nosql.Close()
	log.Printf("Server started on port %v", c.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":8080", r.Router))
}
