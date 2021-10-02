package main

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/repository"
	_ "github.com/lib/pq"
)

var (
	db      *sql.DB
	dbError error
)

func main() {

	var host, port, user, password, dbName = "localhost", "5432", "postgres", "ecetutar", "my-blog"

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, dbError = sql.Open("postgres", dbInfo)

	if dbError != nil {
		panic(dbError)
	}

	categoryRepository := repository.SetCategoryRepositoryDb(db)

	var result = categoryRepository.Get()
	fmt.Println(result)

}
