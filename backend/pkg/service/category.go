package service

import (
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type ICategoryRepository interface {
	GetAllCategory() ([]model.Category, error)
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
