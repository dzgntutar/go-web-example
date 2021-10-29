package repository

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-web-example/pkg/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repository CategoryRepository) GetAll() ([]model.Category, error) {

	var category model.Category
	var categorylist []model.Category

	rows, err := repository.db.Query("select id,title,description,createdate,updatedate from category")

	if err != nil {
		return categorylist, err
	} else {
		for rows.Next() {
			err := rows.Scan(&category.Id, &category.Title, &category.Description, &category.CreateDate, &category.UpdateDate)

			if err != nil {
				fmt.Println(err)
			} else {
				categorylist = append(categorylist, category)
			}
		}

		rows.Close()

		return categorylist, nil
	}
}

func (repository CategoryRepository) Insert(category model.Category) (model.Category, error) {

	stmt, err := repository.db.Prepare("insert into category(title,description,createdate,updatedate) values($1,$2,$3,$4)")

	if err != nil {
		return category, err
	} else {
		result, err := stmt.Exec(category.Title, category.Description, category.CreateDate, category.UpdateDate)
		defer stmt.Close()

		if err != nil {
			return category, err
		} else {
			lastInserted, err := result.LastInsertId()
			if err != nil {
				fmt.Println(err)
			}
			category.Id = lastInserted
			return category, nil
		}
	}
}

func (repository CategoryRepository) GetById(id int32) *model.Category {
	fmt.Println("id:", id)
	stmp, err := repository.db.Prepare("select id,title,description,createdate,updatedate from category where id = $1")

	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		var category model.Category
		err := stmp.QueryRow(id).Scan(&category.Id, &category.Title, &category.Description, &category.CreateDate, &category.UpdateDate)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			return &category
		}
	}
}
