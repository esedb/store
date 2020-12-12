package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
)

func CreateProduct(product *model.Product) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("product")
	r := &repo.ProductRepository{
		C: c,
	}
	err := r.CreateProduct(product)

	return err

}

func UpdateProduct(product *model.Product) (*model.Product, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("product")
	r := &repo.ProductRepository{
		C: c,
	}
	err := r.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, err

}

func SearchPRoductByName(name *string) (*model.Product, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("product")
	r := &repo.ProductRepository{c}
	product, err := r.SearchProduct(*name)
	if err != nil {
		return nil, err
	}
	return product, err

}
