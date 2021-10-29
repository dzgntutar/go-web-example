package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type IPostService interface {
	GetAll() ([]model.Post, error)
	Insert(postDto model.PostDto) error
}

type PostApi struct {
	service IPostService
}

func NewPostApi(service IPostService) PostApi {
	return PostApi{
		service: service,
	}
}

func (p PostApi) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := p.service.GetAll()
		if err != nil {
			model.RespWithError(w, http.StatusNotFound, err.Error())
			return
		}

		model.RespWithJSON(w, http.StatusOK, posts)
	}
}

func (p PostApi) Insert() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var postDto model.PostDto
		err := json.NewDecoder(r.Body).Decode(&postDto)

		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else {
			err := p.service.Insert(postDto)
			if err != nil {
				model.RespWithError(w, 400, "Error")
			} else {
				model.RespWithError(w, 201, "Post Created")
			}
		}
	}
}
