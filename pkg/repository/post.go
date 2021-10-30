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

func (repository PostRepository) Insert(post model.Post) error {
	stmt, err := repository.db.Prepare("insert into post (title,body,createdate ,updatedate ,isshare ,isactive ,categoryid) values($1,$2,$3,$4,$5,$6,$7)")
	defer stmt.Close()
	if err != nil {
		return err
	} else {
		_, err := stmt.Exec(post.Title, post.Body, post.CreateDate, post.UpdateDate, post.IsShare, post.IsActive, post.CategoryId)
		return err
	}
}

func (repository PostRepository) GetAll() ([]model.Post, error) {

	var post model.Post
	var posts []model.Post

	rows, err := repository.db.Query("select id,title,body,createdate,updatedate,isshare,isactive ,categoryid from post")
	if err != nil {
		return posts, err
	} else {
		for rows.Next() {
			err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.CreateDate, &post.UpdateDate, &post.IsShare, &post.IsActive, &post.CategoryId)
			if err != nil {
				fmt.Println(err)
			} else {
				posts = append(posts, post)
			}
		}
		rows.Close()
		return posts, nil
	}
}

func (repository PostRepository) GetById(id int32) (model.Post, error) {
	stmp, err := repository.db.Prepare("select id,title,body,createdate ,updatedate ,isshare ,isactive ,categoryid from post where id = $1")
	var post model.Post
	if err != nil {
		fmt.Println(err)
		return post, err
	} else {
		err := stmp.QueryRow(id).Scan(&post.Id, &post.Title, &post.Body, &post.CreateDate, &post.UpdateDate, &post.IsShare, &post.IsActive, &post.CategoryId)
		defer stmp.Close()
		if err != nil {
			fmt.Println(err)
			return post, err
		} else {
			fmt.Println(post)
			return post, nil
		}
	}
}
