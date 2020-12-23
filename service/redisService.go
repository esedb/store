package service

import (
	"encoding/json"
	"estore/common"
	"time"
)

type RedisService struct{}

func (redis *RedisService) Set(key string, value interface{}, t time.Duration) error {
	var client = common.RedisConn()
	json.Marshal(&value)
	ctx := client.Context()
	err := client.Set(ctx, key, value, t).Err()
	return err
}

func (redis *RedisService) Get(key string) (string, error) {
	var client = common.RedisConn()
	ctx := client.Context()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	client.Close()

	return val, nil
}
