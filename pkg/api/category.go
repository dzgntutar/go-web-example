package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dzqnTtr/go-web-example/pkg/model"
	"github.com/gorilla/mux"
)

type ICagoryService interface {
	GetAll() ([]model.Category, error)
	Insert(categoryDto model.CategoryDto) (model.Category, error)
	GetById(id int32) (*model.Category, error)
}

type CategoryApi struct {
	service ICagoryService
}

func NewCategoryApi(s ICagoryService) CategoryApi {
	return CategoryApi{
		service: s,
	}
}

// FindAllCategories ...
func (c CategoryApi) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		categories, err := c.service.GetAll()
		if err != nil {
			model.RespWithError(w, http.StatusNotFound, err.Error())
			return
		}

		model.RespWithJSON(w, http.StatusOK, categories)
	}
}

//insert category
func (c CategoryApi) Insert() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var categoryDto model.CategoryDto

		err := json.NewDecoder(r.Body).Decode(&categoryDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		category, err := c.service.Insert(categoryDto)
		if err != nil {
			model.RespWithError(w, 400, "Error")
		}
		model.RespWithJSON(w, http.StatusOK, category)
	}
}

func (c CategoryApi) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		category, err := c.service.GetById(int32(idInt))
		//fmt.Println("Category", category)
		if err != nil {
			model.RespWithError(w, http.StatusBadRequest, "Bad Request")
			return
		}
		model.RespWithJSON(w, http.StatusOK, category)
	}
}
