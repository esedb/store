package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
)

type storeService struct{}

var StoreService = &storeService{}

func (service *storeService) CreateStore(store *model.Store) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("store")
	r := &repo.StoreRepository{
		C: c,
	}
	err := r.CreateStore(store)

	return err

}

func (service *storeService) UpdateStore(store *model.Store) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("store")
	r := &repo.StoreRepository{
		C: c,
	}
	err := r.UpdateStore(store)
	if err != nil {
		return err
	}

	return err

}

func (service *storeService) GetAllStore(merchant_id string) ([]model.Store, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("store")
	r := &repo.StoreRepository{
		C: c,
	}

	stores, err := r.GetAllStores(&merchant_id)
	if err != nil {
		return nil, err
	}

	return stores, err

}
