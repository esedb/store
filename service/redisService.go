import (
	"estore/common"
)

type RedisService struct{}

var client = common.RedisConn()

func (redis RedisService) Set(key string, value interface{}) {
	err := client.set(key, value).Err()
	if err {
		panic("An error occured on Redis")
	}
}

func (redis RedisService) Get(key string) interface{} {
	err := client.Get(key).Err()
	if err {
		panic("An error occured trying to retrieve value from redis")
	}
}
