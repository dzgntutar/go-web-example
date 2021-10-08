package service

import (
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type PostRepo interface {
	GetAllPost() ([]model.Post, error)
}

type PostService struct {
	repository PostRepo
}

func NewPostService(repository PostRepo) PostService {
	return PostService{
		repository: repository,
	}
}

func (s PostService) GetAllPost() ([]model.Post, error) {
	return s.repository.GetAllPost()
}
