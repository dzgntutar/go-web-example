package api

import (
	"net/http"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type CategoryApi struct {
	service CagoryService
}

type CagoryService interface {
	GetAllCategory() ([]model.Category, error)
}

func NewCategoryApi(s CagoryService) CategoryApi {
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
