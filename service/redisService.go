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
	err := client.Set(nil, key, value, t).Err()
	return err
}

func (redis *RedisService) Get(key string) (string, error) {
	var client = common.RedisConn()
	val, err := client.Get(nil, key).Result()
	if err != nil {
		return "", err
	}
	client.Close()

	return val, nil
}
