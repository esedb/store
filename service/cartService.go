package service

import (
	"encoding/json"
	"estore/model"
	"fmt"
)

type CartService struct{}

func (c *CartService) AddToCart(cart model.Cart) (string, error) {
	redis := &RedisService{}
	key := cart.UserId
	carts, _ := c.GetFromCart(key)
	fmt.Println("carts: ", carts)
	if carts != nil && len(carts) > 0 {
		for _, c := range carts {
			if c.ItemId == cart.ItemId {
				continue
			} else {
				carts = append(carts, cart)
			}
		}
	} else {
		carts = append(carts, cart)
	}

	result, err := redis.Set(key, carts, 0)

	return result, err
}

func (c *CartService) GetFromCart(key string) ([]model.Cart, error) {
	redis := &RedisService{}
	val, err := redis.Get(key)
	if err != nil {
		return nil, err
	}
	var carts []model.Cart
	json.Unmarshal([]byte(val), &carts)
	fmt.Println("carts: ", carts)

	return carts, nil
}