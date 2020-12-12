package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
)

func CreateStore(store *model.Store) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("store")
	r := &repo.StoreRepository{
		C: c,
	}
	err := r.CreateStore(store)

	return err

}

func UpdateStore(store *model.Store) (*model.Store, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("store")
	r := &repo.StoreRepository{
		C: c,
	}
	err := r.UpdateStore(store)
	if err != nil {
		return nil, err
	}

	return store, err

}
