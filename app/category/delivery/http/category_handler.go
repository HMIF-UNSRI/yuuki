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

	router.GET("/api/categories", handler.List)
	router.POST("/api/categories", handler.Create)
	router.GET("/api/categories/:id", handler.GetByID)
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

func (handler *categoryHandler) GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		panic(domain.NewBadRequestError("failed convert id from string to int"))
	}

	payload := domain.CategoryPayload{}
	payload.ID = id
	payload = handler.categoryUsecase.GetBy(request.Context(), payload)
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}

func (handler *categoryHandler) List(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	payload := handler.categoryUsecase.List(request.Context())
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}
