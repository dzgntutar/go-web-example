package api

import (
	"encoding/json"
	"net/http"

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

// FindAllPosts ...
func (c CategoryApi) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		categories, err := c.CategoryService.All()
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			return
		}

		response, _ := json.Marshal(categories)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
