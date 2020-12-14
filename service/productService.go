package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
)

type productService struct{}

var ProductService = &productService{}

func (service *productService) CreateProduct(product *model.Product) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("product")
	r := &repo.ProductRepository{
		C: c,
	}
	err := r.CreateProduct(product)

	return err

}

func (service *productService) UpdateProduct(product *model.Product) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("product")
	r := &repo.ProductRepository{
		C: c,
	}
	err := r.UpdateProduct(product)

	return err

}

func (service *productService) SearchPRoductByName(name string) (*model.Product, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("product")
	r := &repo.ProductRepository{
		C: c}
	product, err := r.SearchProduct(name)
	if err != nil {
		return nil, err
	}
	return product, err

}
