package repository

import (
	"estore/model"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository struct {
	C *mgo.Collection
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	product.Name = strings.ToLower(product.Name)
	_id := bson.NewObjectId()
	product.Id = _id
	product.CreatedAt = time.Now().Unix()
	err := r.C.Insert(&product)

	return err

}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	product.Name = strings.ToLower(product.Name)
	err := r.C.Update(bson.M{"_id": product.Id}, bson.M{
		"$set": bson.M{
			"name":        product.Name,
			"price":       product.Price,
			"description": product.Description,
			"properties":  product.Properties,
			"files":       product.Files,
			"updated_at":  time.Now().Unix(),
			"qty":         product.Qty,
		},
	})

	return err

}

func (r *ProductRepository) SearchProductByName(name string) (*model.Product, error) {
	var product model.Product
	name = strings.ToLower(name)
	err := r.C.Find(bson.M{"name": name}).One(&product)
	return &product, err
}

func (r *ProductRepository) GetProduct(product_id string) (*model.Product, error) {
	var product model.Product
	err := r.C.Find(bson.M{"_id": product_id}).One(&product)
	if err != nil {
		return nil, err
	}

	return &product, err
}

func (r *ProductRepository) GetAllProducts(store_id *string) []model.Product {
	var products []model.Product
	var product model.Product
	iter := r.C.Find(bson.M{"store_id": store_id}).Iter()

	if iter.Next(&product) {
		products = append(products, product)

	}

	return products
}
