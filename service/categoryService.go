package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
)

type categoryService struct {
	r *repo.CategoryRepository
}

var CategoryService = &categoryService{}

func (service *categoryService) NewCategoryService() *repo.CategoryRepository {
	context := common.NewContext()
	c := context.DbCollection("category")
	service.r = &repo.CategoryRepository{
		C: c,
	}

	return service.r

}

func (service *categoryService) CreateCategory(category *model.Category) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("category")
	r := &repo.CategoryRepository{
		C: c,
	}
	err := r.CreateCategory(category)

	return err
}

func (service *categoryService) UpdateCategory(category *model.Category) error {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("category")
	r := &repo.CategoryRepository{
		C: c,
	}
	err := r.UpdateCategory(category)

	return err

}

func (service *categoryService) SearchCategoryByName(name string) (*model.Category, error) {
	context := common.NewContext()
	defer context.Close()
	c := context.DbCollection("category")
	r := &repo.CategoryRepository{
		C: c,
	}
	category, err := r.SearchCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return category, err
}

func (service *categoryService) GetCollection() *repo.CategoryRepository {
	return service.r
}
