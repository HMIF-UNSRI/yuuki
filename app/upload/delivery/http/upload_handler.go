package http

import (
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type uploadHandler struct {
	uploadUsecase domain.UploadUsecase
}

func RegisterUploadHandler(router *httprouter.Router, usecase domain.UploadUsecase) {
	handler := &uploadHandler{uploadUsecase: usecase}
	directory := http.Dir("./resources")

	router.ServeFiles("/api/resources/*filepath", directory)
	router.POST("/api/uploads", handler.Upload)
}

func (handler *uploadHandler) Upload(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	altText := request.FormValue("alt_text")
	if altText == "" {
		panic(domain.NewBadRequestError("alt text is required"))
	}

	err := request.ParseMultipartForm(domain.MaxImageSize)
	helper.PanicIfErr(err)

	// Save image
	file, fileHeader, err := request.FormFile("image")
	helper.PanicIfErr(err)
	defer file.Close()

	// Validation
	if fileHeader.Filename == "" {
		panic(domain.NewBadRequestError("image is empty"))
	}

	if fileHeader.Size > domain.MaxImageSize {
		panic(domain.NewBadRequestError("image size exceeds the capacity of 1MB"))
	}

	extension := filepath.Ext(fileHeader.Filename)
	if extension != ".jpg" && extension != ".jpeg" && extension != ".png" {
		panic(domain.NewBadRequestError("only jpg, jpeg and png formats are accepted"))
	}

	// Generate new image name
	imageName := uuid.NewString() + extension
	fileDestination, err := os.Create("resources/" + imageName)
	helper.PanicIfErr(err)

	_, err = io.Copy(fileDestination, file)
	helper.PanicIfErr(err)

	payload := domain.UploadPayload{
		ImageName: imageName,
		AltText:   altText,
	}

	// Save to database
	payload = handler.uploadUsecase.Create(request.Context(), payload)
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}
