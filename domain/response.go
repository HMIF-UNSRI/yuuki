package domain

import "net/http"

type WebResponse struct {
	Code   uint16      `json:"code"`
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func NewResponse200(data interface{}) WebResponse {
	return WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   data,
	}
}
