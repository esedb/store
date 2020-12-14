package repository

import (
	"estore/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//StoreRepository encapsulate store crude operations
type StoreRepository struct {
	C *mgo.Collection
}

//CreateStore creat User in Collection
func (r *StoreRepository) CreateStore(store *model.Store) error {
	err := r.C.Insert(&store)

	return err

}

//UpdateStore store collection
func (r *StoreRepository) UpdateStore(store *model.Store) error {
	err := r.C.Update(bson.M{"_id": store.Id}, bson.M{
		"$set": bson.M{
			"name":     store.Name,
			"location": store.Location,
			"Address":  store.Address,
			"Phone":    store.Phone,
		},
	})

	return err

}

//UpdatePhone update store phone number
func (r *StoreRepository) UpdatePhone(store *model.Store) error {
	err := r.C.Update(bson.M{"_id": store.Id}, bson.M{
		"$set": bson.M{
			"Phone": store.Phone,
		},
	})

	return err
}
