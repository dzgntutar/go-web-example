package main

import (
	"database/sql"
	"fmt"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/model"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/repository"
	_ "github.com/lib/pq"
)

var (
	dbPstgr *sql.DB
	dbError error
)

func main() {

	var host, port, user, password, dbName = "localhost", "5432", "postgres", "Password1*", "my-db"

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	dbPstgr, dbError = sql.Open("postgres", dbInfo)

	if dbError != nil {
		panic(dbError)
	}

	categoryRepo := repository.NewCategoryRepositoryDb(dbPstgr)
	//categoryRepo := repository.CategoryRepository{Db: dbPstgr}

	categoryRepo.Insert(model.Category{Title: "C sharp", Description: "c sharp dili"})

	var result = categoryRepo.Get()
	fmt.Println(result)

}
