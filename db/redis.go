package db

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/redis"

	c "grapi/config"
)

// OpenNoSQL : Connect to the redis server with config file infos
func OpenNoSQL() {
	connectionStr := fmt.Sprintf("%s:%s",
		c.Cfg.NoSQLAddress,
		c.Cfg.NoSQLPort)
	var err error

	if c.Cfg.NoSQL == 0 {
		Nosql = nil
		return
	}
	Nosql, err = redis.Dial(c.Cfg.NoSQLConnType, connectionStr)
	if err != nil {
		defer Nosql.Close()
		log.Fatal("Cannot connect to redis")
	}
	log.Printf("Redis connected with address: %s\n", connectionStr)
}
