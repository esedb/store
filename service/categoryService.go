package service

import (
	"estore/common"
	"estore/model"
	repo "estore/repository"
)

func init() {

}

type CategoryService struct {
	r *repo.CategoryRepository
}

func (service *CategoryService) NewCategoryService() *CategoryService {
	context := common.NewContext()
	c := context.DbCollection("category")
	service.r = &repo.CategoryRepository{
		C: c,
	}

	return service

}

func (service *CategoryService) CreateCategory(category *model.Category) error {
	err := service.r.CreateCategory(category)
	return err
}

func (service *CategoryService) UpdateCategory(category *model.Category) error {
	err := service.r.UpdateCategory(category)

	return err

}

func (service *CategoryService) SearchCategoryByName(name string) (*model.Category, error) {
	category, err := service.r.SearchCategoryByName(name)
	if err != nil {
		return nil, err
	}
	return category, err
}

func (service *CategoryService) GetCollection() *repo.CategoryRepository {
	return service.r
}
