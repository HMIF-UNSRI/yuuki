package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type categoryHandler struct {
	categoryUsecase domain.CategoryUsecase
}

func RegisterProductHandler(router *httprouter.Router, usecase domain.CategoryUsecase) {
	handler := &categoryHandler{categoryUsecase: usecase}
	router.POST("/api/categories", handler.Create)
	router.PUT("/api/categories/:id", handler.Update)
}

func (handler *categoryHandler) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	payload := domain.CategoryPayload{}
	helper.ReadFromRequestBody(request, &payload)

	payload = handler.categoryUsecase.Create(request.Context(), payload)
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}

func (handler *categoryHandler) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(domain.NewBadRequestError("failed convert id from string to int"))
	}

	payload := domain.CategoryPayload{}
	helper.ReadFromRequestBody(request, &payload)

	payload.ID = id
	payload = handler.categoryUsecase.Update(request.Context(), payload)
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}
