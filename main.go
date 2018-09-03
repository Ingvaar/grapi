package main

import (
	"log"
	"net/http"

	c "github.com/ingvaar/grapi/config"
	"github.com/ingvaar/grapi/db"
	r "github.com/ingvaar/grapi/router"
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
