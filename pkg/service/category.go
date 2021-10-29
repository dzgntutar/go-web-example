package service

import (
	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type ICategoryRepository interface {
	GetAllCategory() ([]model.Category, error)
	InsertCategory(category model.Category) (model.Category, error)
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

func (c CategoryService) GetAllCategory() ([]model.Category, error) {

	return c.repository.GetAllCategory()
}

func (c CategoryService) InsertCategory(categoryDto model.CategoryDto) (model.Category, error) {
	return c.repository.InsertCategory(categoryDto.ToCategory())
}

func (c CategoryService) GetById(id int32) *model.Category {
	return c.repository.GetById(id)
}
