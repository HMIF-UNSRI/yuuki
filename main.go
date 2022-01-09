package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"net/http"
	_categoryHandler "yuuki/app/category/delivery/http"
	_categoryRepo "yuuki/app/category/repository"
	_categoryUsecase "yuuki/app/category/usecase"

	"yuuki/app/mariadb"
	_shortenerHandler "yuuki/app/shortener/delivery/http"
	_shortenerRepo "yuuki/app/shortener/repository"
	_shortenerUsecase "yuuki/app/shortener/usecase"

	_uploadHandler "yuuki/app/upload/delivery/http"
	_uploadRepo "yuuki/app/upload/repository"
	_uploadUsecase "yuuki/app/upload/usecase"

	"yuuki/middleware"
	"yuuki/pkg/config"
	"yuuki/pkg/exception"
	"yuuki/pkg/helper"
)

func main() {
	configuration := config.NewConfiguration(`./.env`)
	database := mariadb.GetConnection(configuration)
	validate := validator.New()

	categoryRepository := _categoryRepo.NewCategoryRepository(database)
	categoryUsecase := _categoryUsecase.NewCategoryUsecase(categoryRepository, validate)

	uploadRepository := _uploadRepo.NewUploadRepository(database)
	uploadUsecase := _uploadUsecase.NewUploadUsecase(uploadRepository)

	shortenerRepository := _shortenerRepo.NewShortenerRepository(database)
	shortenerUsecase := _shortenerUsecase.NewShortenerUsecase(shortenerRepository, validate)

	logger, err := zap.NewProduction()
	helper.PanicIfErr(err)
	defer logger.Sync()

	sugar := logger.Sugar()

	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	_categoryHandler.RegisterCategoryHandler(router, categoryUsecase)
	_uploadHandler.RegisterUploadHandler(router, uploadUsecase)
	_shortenerHandler.RegisterShortenerHandler(router, shortenerUsecase)

	sugar.Infow("listening and serving http on :8080")
	sugar.Fatal(http.ListenAndServe(":8080", &middleware.LogMiddleware{Handler: router, Logger: sugar}))
}
