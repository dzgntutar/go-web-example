package repository

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (repository PostRepository) GetAllPost() ([]model.Post, error) {

	var post model.Post
	var posts []model.Post

	// ,createdate,updatedate,isshare,isactive,categoryid
	rows, err := repository.db.Query("select id,title,body from post")
	fmt.Println(rows)
	if err != nil {
		return posts, err
	} else {
		for rows.Next() {
			// , &post.CreateDate, &post.UpdateDate, &post.IsShare, &post.IsActive, &post.CategoryId
			err = rows.Scan(&post.Id, &post.Title, &post.Body)
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
