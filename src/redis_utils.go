package main

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
	"os"
)

// openRedis : Connect to the redis server with config file infos
func openRedis(cfg Config) *redis.Client {
	connectionStr := fmt.Sprintf("%s:%s",
		cfg.RedisAddress,
		cfg.RedisPort)

	if cfg.Redis == 0 {
		return (nil)
	}
	db, err := redis.Dial(cfg.RedisConnType, connectionStr)
	if err != nil {
		defer db.Close()
		log.Fatal("Cannot connect to redis")
		os.Exit(1)
	}
	fmt.Println("Redis connected")
	return (db)
}
