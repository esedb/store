package repository

import (
	"estore/model"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CategoryRepository struct {
	C *mgo.Collection
}

func (r *CategoryRepository) CreateCategory(category *model.Category) error {
	_id := bson.NewObjectId()
	category.Id = _id
	err := r.C.Insert(&category)

	return err

}

func (r *CategoryRepository) UpdateCategory(category *model.Category) error {
	fmt.Println("Category ID: ", category.Id)

	err := r.C.Update(bson.M{"_id": category.Id}, bson.M{
		"$set": bson.M{
			"name": category.Name,
		},
	})

	return err

}

func (r *CategoryRepository) SearchCategoryByName(name string) (*model.Category, error) {
	var category model.Category
	err := r.C.Find(bson.M{"name": name}).One(&category)
	return &category, err
}
