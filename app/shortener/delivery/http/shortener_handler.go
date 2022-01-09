package http

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"yuuki/domain"
	"yuuki/pkg/helper"
)

type shortenerHandler struct {
	shortenerUsecase domain.ShortenerUsecase
}

func RegisterShortenerHandler(router *httprouter.Router, usecase domain.ShortenerUsecase) {
	handler := &shortenerHandler{shortenerUsecase: usecase}

	router.POST("/api/shorteners", handler.Create)
	router.GET("/api/shorteners/:slug", handler.GetBySlug)
}

func (handler *shortenerHandler) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	payload := domain.ShortenerPayload{}
	helper.ReadFromRequestBody(request, &payload)

	payload = handler.shortenerUsecase.Create(request.Context(), payload)
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}

func (handler *shortenerHandler) GetBySlug(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	payload := domain.ShortenerPayload{}
	payload.Slug = params.ByName("slug")

	payload = handler.shortenerUsecase.GetBySlug(request.Context(), payload)
	helper.WriteToResponseBody(writer, domain.NewResponse200(payload))
}
