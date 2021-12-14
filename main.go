package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	deliveryHttp "yuuki/app/category/delivery/http"
	"yuuki/app/category/repository"
	"yuuki/app/category/usecase"
	"yuuki/app/mariadb"
	"yuuki/pkg/config"
	"yuuki/pkg/exception"
)

func main() {
	configuration := config.NewConfiguration(`./.env`)
	database := mariadb.GetConnection(configuration)
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository(database)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository, validate)

	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	deliveryHttp.RegisterProductHandler(router, categoryUsecase)

	log.Fatal(http.ListenAndServe(":8080", router))
}
