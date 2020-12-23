package service

import (
	"encoding/json"
	"estore/common"
	"time"
)

type RedisService struct{}

func (redis *RedisService) Set(key string, value interface{}, t time.Duration) (string, error) {
	var client = common.RedisConn()
	defer client.Close()
	_str, _ := json.Marshal(value)
	ctx := client.Context()
	result, err := client.Set(ctx, key, _str, t).Result()

	return result, err
}

func (redis *RedisService) Get(key string) (string, error) {
	var client = common.RedisConn()
	defer client.Close()
	ctx := client.Context()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
