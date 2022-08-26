package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseAPI(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	responseApi := Response{
		Meta: meta,
		Data: data,
	}

	return responseApi
}

func FormatError(err error) []string {
	var newError []string

	for _, e := range err.(validator.ValidationErrors) {
		newError = append(newError, e.Error())
	}

	return newError
}
