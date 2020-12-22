package service

import (
	"estore/model"
)

type CartService struct{}

func (c *CartService) AddToCart(cart model.Cart) error {
	redis := &RedisService{}
	key := cart.UserId
	carts := []model.Cart{
		cart,
	}
	err := redis.Set(key, carts, 0)

	return err
}

func (c *CartService) GetFromCart(key string) (interface{}, error) {
	redis := &RedisService{}
	val, err := redis.Get(key)
	if err != nil {
		return nil, err
	}

	return val, nil
}
