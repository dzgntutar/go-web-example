package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dzqnTtr/go-web-example/pkg/model"
	"github.com/gorilla/mux"
)

type IPostService interface {
	GetAll() ([]model.Post, error)
	Insert(postDto model.PostDto) error
	GetById(id int32) (model.PostDto, error)
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

func (p PostApi) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		post, err := p.service.GetById(int32(idInt))
		if err != nil {
			model.RespWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		model.RespWithJSON(w, http.StatusOK, post)
	}
}
