package service

import (
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type CategoryService struct {
	repository CategoryRepo
}

type CategoryRepo interface {
	GetAllCategory() ([]model.Category, error)
}

func NewCategoryService(repo CategoryRepo) CategoryService {
	return CategoryService{
		repository: repo,
	}
}

func (c CategoryService) GetAllCategory() ([]model.Category, error) {
	return c.repository.GetAllCategory()
}
