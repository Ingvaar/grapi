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
	db.OpenSQLDatabase()
	db.OpenRedis()
	r.NewRouter()

	defer db.Db.SQL.Close()
	defer db.Db.Redis.Close()
	log.Printf("Server started on port %v", c.Cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":8080", r.Router))
}
