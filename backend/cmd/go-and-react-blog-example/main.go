package main

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
	_ "github.com/lib/pq"
)

var (
	db      *sql.DB
	dbError error
)

func main() {

	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "ecetutar"
	dbName := "my-blog"

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, dbError = sql.Open("postgres", dbInfo)

	if dbError != nil {
		panic(dbError)
	}

	//insertCategory()
	selectCategory()
}

func insertCategory() {
	r, err := db.Exec("insert into category(title,description) values('db','aciklama')")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r.RowsAffected())
	}
}

func selectCategory() {

	var category model.Category
	var categorylist []model.Category

	rows, err := db.Query("select id,title,description from category")

	if err != nil {
		fmt.Println(err)
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

		fmt.Println(categorylist)
	}
}
