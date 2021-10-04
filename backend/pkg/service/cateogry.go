package service

import (
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/repository"
)

type CategoryService struct {
	repository *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) CategoryService {
	return CategoryService{
		repository: repo,
	}
}

func (c CategoryService) All() ([]model.Category, error) {
	return c.repository.All()
}
