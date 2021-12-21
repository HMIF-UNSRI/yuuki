package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
	_categoryHandler "yuuki/app/category/delivery/http"
	_categoryRepo "yuuki/app/category/repository"
	_categoryUsecase "yuuki/app/category/usecase"
	"yuuki/app/mariadb"
	_uploadHandler "yuuki/app/upload/delivery/http"
	_uploadRepo "yuuki/app/upload/repository"
	_uploadUsecase "yuuki/app/upload/usecase"
	"yuuki/middleware"
	"yuuki/pkg/config"
	"yuuki/pkg/exception"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
}

func main() {
	configuration := config.NewConfiguration(`./.env`)
	database := mariadb.GetConnection(configuration)
	validate := validator.New()

	categoryRepository := _categoryRepo.NewCategoryRepository(database)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryRepository, validate)

	uploadRepository := _uploadRepo.NewUploadRepository(database)
	uploadUsecase := _uploadUsecase.NewUploadUsecase(uploadRepository)

	router := httprouter.New()
	//logMiddleware := &middleware.LogMiddleware{Handler: router}
	router.PanicHandler = exception.ErrorHandler

	_categoryHandler.RegisterProductHandler(router, categoryUsecase)
	_uploadHandler.RegisterUploadHandler(router, uploadUsecase)

	log.Fatal(http.ListenAndServe(":8080", &middleware.LogMiddleware{Handler: router}))
}
