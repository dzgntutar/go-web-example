package api

import (
	"net/http"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type ICagoryService interface {
	GetAllCategory() ([]model.Category, error)
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
func (c CategoryApi) GetAllCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		categories, err := c.service.GetAllCategory()
		if err != nil {
			model.RespWithError(w, http.StatusNotFound, err.Error())
			return
		}

		model.RespWithJSON(w, http.StatusOK, categories)

	}
}
