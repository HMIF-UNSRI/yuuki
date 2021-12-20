package domain

import "net/http"

type WebResponse struct {
	Code   uint16      `json:"code"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

type WebResponsePagination struct {
	Code       uint16      `json:"code"`
	Status     string      `json:"status"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

func NewResponse200(data interface{}) WebResponse {
	return WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   data,
	}
}

func NewResponsePagination200(data interface{}, pagination Pagination) WebResponsePagination {
	return WebResponsePagination{
		Code:       200,
		Status:     http.StatusText(200),
		Data:       data,
		Pagination: pagination,
	}
}
