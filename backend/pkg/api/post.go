package api

import (
	"net/http"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/service"
)

type PostApi struct {
	service service.PostService
}

func NewPostApi(service service.PostService) PostApi {
	return PostApi{}
}

func (p PostApi) All() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		posts, err := p.service.All()
		if err != nil {
			model.RespWithError(w, http.StatusNotFound, err.Error())
			return
		}

		model.RespWithJSON(w, http.StatusOK, posts)

	}
}
