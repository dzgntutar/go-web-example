package repository

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type PostRepository struct {
	db *sql.DB
}

func (repository PostRepository) All() ([]model.Post, error) {

	var post model.Post
	var posts []model.Post

	rows, err := repository.db.Query("select * from id,Title,Body,CreateAt,UpdateAt,UpdateAt,IsShare,IsActive ,CategoryId")
	if err != nil {
		return posts, err
	} else {
		for rows.Next() {
			err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreateAt, &post.UpdateAt, &post.IsShare, &post.IsActive, &post.CategoryId)
			if err != nil {
				fmt.Println(err)
			} else {
				posts = append(posts, post)
			}

			rows.Close()
		}

		return posts, nil
	}
}
