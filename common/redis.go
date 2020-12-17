package common

import (
	"github.com/go-redis/redis/v8"
)

func RedisObj () {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}