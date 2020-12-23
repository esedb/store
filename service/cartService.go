package service

import (
	"encoding/json"
	"estore/model"
	"fmt"
)

type CartService struct{}

func (c *CartService) AddToCart(cart model.Cart) error {
	redis := &RedisService{}
	key := cart.UserId
	fmt.Println("keys: ", key)
	val, err := c.GetFromCart(key)
	if err != nil {
		return err
	}
	if val != nil {
		val = append(val, cart)
	}
	carts := []model.Cart{
		cart,
	}
	err = redis.Set(key, carts, 0)

	return err
}

func (c *CartService) GetFromCart(key string) ([]model.Cart, error) {
	redis := &RedisService{}
	val, err := redis.Get(key)
	if err != nil {
		return nil, err
	}
	var carts []model.Cart
	json.Unmarshal([]byte(val), &carts)

	return carts, nil
}
