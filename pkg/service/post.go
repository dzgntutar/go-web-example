package service

import (
	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type IPostRepository interface {
	GetAll() ([]model.Post, error)
	Insert(post model.Post) error
	GetById(id int32) (model.Post, error)
}

type PostService struct {
	repository IPostRepository
}

func NewPostService(repository IPostRepository) PostService {
	return PostService{
		repository: repository,
	}
}

func (s PostService) GetAll() ([]model.Post, error) {
	return s.repository.GetAll()
}

func (s PostService) Insert(postDto model.PostDto) error {
	return s.repository.Insert(postDto.ToPost())
}

func (s PostService) GetById(id int32) (model.PostDto, error) {
	post, err := s.repository.GetById(id)

	return post.ToPostDto(), err
}
