package repository

import (
	"estore/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository struct {
	C *mgo.Collection
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	err := r.C.Insert(&product)

	return err

}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	err := r.C.Update(bson.M{"_id": product.Id}, bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"price":       product.Price,
			"description": product.Description,
			"properties":  product.Properties,
			"files":       product.Files,
			"store_id":    product.StoreID,
			"updated_at":  product.UpdatedAt,
		},
	})

	return err

}

func (r *ProductRepository) SearchProduct(name string) (*model.Product, error) {
	var product model.Product
	err := r.C.Find(bson.M{"name": name}).One(&product)
	return &product, err
}

func (r *ProductRepository) GetAllProducts() []model.Product {
	var products []model.Product
	var product model.Product
	iter := r.C.Find(nil).Iter()

	if iter.Next(&product) {
		products = append(products, product)

	}

	return products
}
