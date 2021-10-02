package repository

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func SetCategoryRepositoryDb(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repository CategoryRepository) Insert(category model.Category) {
	stmt, err := repository.db.Prepare("insert into category(title,description) values($1,$2)")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := stmt.Exec(category.Title, category.Description)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result.RowsAffected())
		}
	}
}

func (repository CategoryRepository) Get() []model.Category {

	var category model.Category
	var categorylist []model.Category

	rows, err := repository.db.Query("select id,title,description from category")

	if err != nil {
		return categorylist
	} else {
		for rows.Next() {
			err := rows.Scan(&category.Id, &category.Title, &category.Description)

			if err != nil {
				fmt.Println(err)
			} else {
				categorylist = append(categorylist, category)
			}

			rows.Close()
		}

		return categorylist
	}
}

func (repository CategoryRepository) GetById(id int32) *model.Category {
	stmp, err := repository.db.Prepare("select id, title , description from category where id = $1")

	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		var category model.Category
		err := stmp.QueryRow(id).Scan(&category.Id, &category.Title, &category.Description)
		if err != nil {
			fmt.Println(err)
			return nil
		} else {
			return &category
		}

	}
}
