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

func init() {
	CategoryService.NewCategoryService()
}

func (service *categoryService) NewCategoryService() *categoryService {
	context := common.NewContext()
	c := context.DbCollection("category")
	service.r = &repo.CategoryRepository{
		C: c,
	}

	return service

}

func (service *categoryService) CreateCategory(category *model.Category) error {
	err := service.r.CreateCategory(category)

	return err
}

func (service *categoryService) UpdateCategory(category *model.Category) error {
	err := service.r.UpdateCategory(category)

	return err

}

func (service *categoryService) SearchCategoryByName(name string) (*model.Category, error) {
	category, err := service.r.SearchCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return category, err
}

func (service *categoryService) GetCollection() *repo.CategoryRepository {
	return service.r
}
