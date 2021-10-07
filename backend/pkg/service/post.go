package service

import (
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/repository"
)

type PostService struct {
	repository *repository.PostRepository
}

func (s PostService) All() ([]model.Post, error) {
	return s.repository.All()
}
