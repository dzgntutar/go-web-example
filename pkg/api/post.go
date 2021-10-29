package api

import (
	"net/http"

	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type IPostService interface {
	GetAllPost() ([]model.Post, error)
}

type PostApi struct {
	service IPostService
}

func NewPostApi(service IPostService) PostApi {
	return PostApi{
		service: service,
	}
}

func (p PostApi) GetAllPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := p.service.GetAllPost()
		if err != nil {
			model.RespWithError(w, http.StatusNotFound, err.Error())
			return
		}

		model.RespWithJSON(w, http.StatusOK, posts)
	}
}
