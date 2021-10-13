package repository

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repository CategoryRepository) GetAllCategory() ([]model.Category, error) {

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

func (repository CategoryRepository) InsertCategory(category model.Category) {

	stmt, err := repository.db.Prepare("insert into category(title,description,createdate,updatedate) values($1,$2,$3,$4)")

	if err != nil {
		fmt.Println(err)
	} else {
		result, err := stmt.Exec(category.Title, category.Description, category.CreateDate, category.UpdateDate)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result.RowsAffected())
		}
	}
}

// func (repository CategoryRepository) GetById(id int32) *model.Category {
// 	stmp, err := repository.db.Prepare("select id, title , description from category where id = $1")

// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	} else {
// 		var category model.Category
// 		err := stmp.QueryRow(id).Scan(&category.Id, &category.Title, &category.Description)
// 		if err != nil {
// 			fmt.Println(err)
// 			return nil
// 		} else {
// 			return &category
// 		}

// 	}
// }
