package api

import (
	"net/http"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/service"
)

type CategoryApi struct {
	CategoryService service.CategoryService
}

func NewApi(s service.CategoryService) CategoryApi {
	return CategoryApi{
		CategoryService: s,
	}
}

// FindAllCategories ...
func (c CategoryApi) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		categories, err := c.CategoryService.All()
		if err != nil {
			model.RespWithError(w, http.StatusNotFound, err.Error())
			return
		}

		model.RespWithJSON(w, http.StatusOK, categories)

	}
}
