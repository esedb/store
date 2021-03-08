package repository

import (
	"estore/model"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ItemRepository struct {
	C *mgo.Collection
}

func (r *ItemRepository) CreateItem(item *model.Item) error {
	item.Name = strings.ToLower(item.Name)
	_id := bson.NewObjectId()
	item.Id = _id
	item.CreatedAt = time.Now().Unix()
	err := r.C.Insert(&item)

	return err

}

func (r *ItemRepository) UpdateItem(item *model.Item) error {
	item.Name = strings.ToLower(item.Name)
	err := r.C.Update(bson.M{"_id": item.Id}, bson.M{
		"$set": bson.M{
			"name":        item.Name,
			"price":       item.Price,
			"description": item.Description,
			"properties":  item.Properties,
			"files":       item.Files,
			"updated_at":  time.Now().Unix(),
			"qty":         item.Qty,
		},
	})

	return err

}

func (r *ItemRepository) SearchItemByName(name string) (*model.Item, error) {
	var item model.Item
	name = strings.ToLower(name)
	err := r.C.Find(bson.M{"name": name}).One(&item)
	return &item, err
}

func (r *ItemRepository) GetItem(item_id string) (*model.Item, error) {
	var item model.Item
	err := r.C.Find(bson.M{"_id": item_id}).One(&item)
	if err != nil {
		return nil, err
	}

	return &item, err
}

func (r *ItemRepository) GetAllItemsByStore(store_id *string) []model.Item {
	var items []model.Item
	var item model.Item
	iter := r.C.Find(bson.M{"store_id": store_id}).Iter()

	for iter.Next(&item) {
		items = append(items, item)

	}

	return items
}

func (r *ItemRepository) GetAllItems() ([]model.Item, error) {
	var items []model.Item
	var item model.Item
	iter := r.C.Find(nil).Iter()
	err := iter.Err()
	if err != nil {
		return nil, err
	}

	for iter.Next(&item) {
		items = append(items, item)

	}

	return items, nil
}
