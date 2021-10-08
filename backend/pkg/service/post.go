package service

import (
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/repository"
)

type PostService struct {
	repository *repository.PostRepository
}

func NewPostService(repository *repository.PostRepository) PostService {
	return PostService{
		repository: repository,
	}
}

func (s PostService) GetAllPost() ([]model.Post, error) {
	return s.repository.GetAllPost()

}
