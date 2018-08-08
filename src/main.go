package main

import (
	"database/sql"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
	"net/http"
)

var dbSQL *sql.DB
var cfg Config
var redisCli *redis.Client

func main() {
	opt := ParsCmdline()
	cfg = GetConfig(opt)
	dbSQL = openSQLDatabase(cfg)
	redisCli = openRedis(cfg)
	router := NewRouter(opt)

	defer dbSQL.Close()
	defer redisCli.Close()
	log.Fatal(http.ListenAndServe(":8080", router))
}
