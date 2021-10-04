package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/api"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/repository"
	"github.com/dzqnTtr/go-and-react-blog-example/backend/pkg/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	Db     *sql.DB
}

var (
	dbError error
)

func main() {

	app := App{}

	var host, port, user, password, dbName = "localhost", "5432", "postgres", "Password1*", "my-db"

	app.Initialize(host, port, user, password, dbName)

	app.routes()

	app.Run(":8080")
}

func (app *App) Initialize(host, port, user, password, dbName string) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	app.Db, dbError = sql.Open("postgres", connectionString)

	if dbError != nil {
		panic(dbError)
	}

	app.Router = mux.NewRouter()
}

func (app *App) routes() {
	categoryApi := initCategoryApi(app.Db)
	app.Router.HandleFunc("/category", categoryApi.All()).Methods("GET")
}

func (app *App) Run(port string) {
	fmt.Printf("Server started at %s\n", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func initCategoryApi(db *sql.DB) api.CategoryApi {
	categoryRepo := repository.NewCategoryRepositoryDb(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryApi := api.NewApi(categoryService)
	return categoryApi
}
