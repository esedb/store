package service

import (
	"estore/common"
	"time"
)

type RedisService struct{}

func (redis *RedisService) Set(key string, value interface{}, t time.Duration) error {
	var client = common.RedisConn()
	err := client.Set(nil, key, value, t).Err()
	return err
}

func (redis *RedisService) Get(key string) (interface{}, error) {
	var client = common.RedisConn()
	val, err := client.Get(nil, key).Result()
	if err != nil {
		return nil, err
	}
	client.Close()

	return val, nil
}
