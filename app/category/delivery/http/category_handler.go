package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type categoryHandler struct {
	categoryUsecase domain.CategoryUsecase
}

func RegisterProductHandler(router *httprouter.Router, usecase domain.CategoryUsecase) {
	handler := &categoryHandler{categoryUsecase: usecase}
	router.POST("/api/categories", handler.Create)
}

func (handler *categoryHandler) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	payload := domain.CategoryPayload{}
	helper.ReadFromRequestBody(request, &payload)

	payload = handler.categoryUsecase.Create(request.Context(), payload)
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}
