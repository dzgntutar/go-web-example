package service

import (
	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type ICategoryRepository interface {
	GetAll() ([]model.Category, error)
	Insert(category model.Category) (model.Category, error)
	GetById(id int32) *model.Category
}

type CategoryService struct {
	repository ICategoryRepository
}

func NewCategoryService(repo ICategoryRepository) CategoryService {
	return CategoryService{
		repository: repo,
	}
}

func (c CategoryService) GetAll() ([]model.Category, error) {

	return c.repository.GetAll()
}

func (c CategoryService) Insert(categoryDto model.CategoryDto) (model.Category, error) {
	return c.repository.Insert(categoryDto.ToCategory())
}

func (c CategoryService) GetById(id int32) *model.Category {
	return c.repository.GetById(id)
}
