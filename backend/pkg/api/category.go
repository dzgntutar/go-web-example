package api

import (
	"encoding/json"
	"net/http"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type ICagoryService interface {
	GetAllCategory() ([]model.Category, error)
	InsertCategory(categoryDto model.CategoryDto)
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

		categories, err := c.service.GetAllCategory()
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
		c.service.InsertCategory(categoryDto)
		model.RespWithJSON(w, http.StatusOK, "Create is successful")
	}
}
