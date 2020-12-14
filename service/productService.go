package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
	"time"
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
	product.CreatedAt = time.Now().Unix()
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
	product.UpdatedAt = time.Now().Unix()
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

func (service *productService) GetAllProducts(store_id string) []model.Product {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("product")
	r := &repo.ProductRepository{
		C: c}
	product := r.GetAllProducts(&store_id)

	return product

}
