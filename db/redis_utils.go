package db

import (
	"fmt"
	"log"

	"github.com/mediocregopher/radix.v2/redis"

	c "grapi/config"
)

// OpenRedis : Connect to the redis server with config file infos
func OpenRedis() {
	connectionStr := fmt.Sprintf("%s:%s",
		c.Cfg.RedisAddress,
		c.Cfg.RedisPort)
	var err error

	if c.Cfg.Redis == 0 {
		Db.Redis = nil
		return
	}
	Db.Redis, err = redis.Dial(c.Cfg.RedisConnType, connectionStr)
	if err != nil {
		defer Db.Redis.Close()
		log.Fatal("Cannot connect to redis")
	}
	log.Printf("Redis connected with address: %s\n", connectionStr)
}
