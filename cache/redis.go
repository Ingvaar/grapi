package redis

import (
	"fmt"
	"grapi/core"
	"log"

	"github.com/mediocregopher/radix.v2/redis"
)

// Redis :
type Redis struct {
	RC *redis.Client
}

// Connect :
func (rd *Redis) Connect(config *core.Config) {
	connectionStr := fmt.Sprintf("%s:%s",
		config.CacheAddress,
		config.CachePort)

	if config.Cache != 0 {
		var err error
		rd.RC, err = redis.Dial(config.CacheConnType, connectionStr)
		if err != nil {
			defer rd.RC.Close()
			log.Fatal("Cannot connect to redis")
		}
		log.Printf("Redis connected with address: %s\n", connectionStr)
	}
}

// Register : register the functions to handler map
func (rd *Redis) Register(handlers *core.Handlers) {
	redis := Redis{RC: rd.RC}
	temp := core.Handlers{}

	for key, value := range *handlers {
		temp[key] = value
	}
	temp["set"] = redis.Set
	temp["read"] = redis.Read
	temp["remove"] = redis.Delete
	*handlers = temp
}
