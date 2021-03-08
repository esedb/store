package common

import (
	"os"

	"github.com/go-redis/redis/v8"
)

func RedisConn() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	address := ""
	if host != "" {
		address = host + ":6379"
	} else {
		address = "localhost:6379"

	}
	rdb := redis.NewClient(&redis.Options{
		Addr: address,
		// Password: "", // no password set
		// DB:       0,  // use default DB
	})

	return rdb
}
