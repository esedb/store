package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
	"time"
)

type itemService struct{}

var ItemService = &itemService{}

func (service *itemService) CreateItem(item *model.Item) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("items")
	r := &repo.ItemRepository{
		C: c,
	}
	item.CreatedAt = time.Now().Unix()
	err := r.CreateItem(item)

	return err

}

func (service *itemService) UpdateItem(item *model.Item) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("items")
	r := &repo.ItemRepository{
		C: c,
	}
	item.UpdatedAt = time.Now().Unix()
	err := r.UpdateItem(item)

	return err

}

func (service *itemService) SearchItemByName(name string) (*model.Item, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("items")
	r := &repo.ItemRepository{
		C: c}
	item, err := r.SearchItemByName(name)
	if err != nil {
		return nil, err
	}
	return item, err

}

func (service *itemService) GetItem(item_id string) (*model.Item, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("items")
	r := &repo.ItemRepository{
		C: c}
	item, err := r.GetItem(item_id)
	if err != nil {
		return nil, err

	}

	return item, err

}

func (service *itemService) GetAllItemsByStore(store_id string) []model.Item {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("items")
	r := &repo.ItemRepository{
		C: c}
	item := r.GetAllItemsByStore(&store_id)

	return item

}

func (service *itemService) GetAllItems() ([]model.Item, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("items")
	r := &repo.ItemRepository{
		C: c}
	items, err := r.GetAllItems()
	if err != nil {
		return nil, err

	}

	return items, err

}
