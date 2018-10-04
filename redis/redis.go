package redis

import (
	"fmt"
	"grapi/core"
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

// Database :
type Database struct {
	DB *redis.Client
}

// Connect :
func (db *Database) Connect(config *core.Config) {
	connectionStr := fmt.Sprintf("%s:%s",
		config.NoSQLAddress,
		config.NoSQLPort)

	if config.NoSQL != 0 {
		var err error
		db.DB, err = redis.Dial(config.NoSQLConnType, connectionStr)
		if err != nil {
			defer db.DB.Close()
			log.Fatal("Cannot connect to redis")
		}
		log.Printf("Redis connected with address: %s\n", connectionStr)
	}
}

// Register : register the functions to handler map
func (db *Database) Register(handlers *core.Handlers) {
	redis := Database{DB: db.DB}
	temp := core.Handlers{}

	for key, value := range *handlers {
		temp[key] = value
	}
	temp["set"] = redis.Set
	temp["read"] = redis.Read
	temp["remove"] = redis.Delete
	*handlers = temp
}
