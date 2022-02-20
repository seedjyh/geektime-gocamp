package echo

import (
	"geektime-gocamp/week4/homework/internal/student"
)

func NewRegister(service *student.Service) *register {
	return &register{service: service}
}

func NewUnregister(service *student.Service) *unregister {
	return &unregister{service: service}
}

type ResponseBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var defaultSuccessResponse = &ResponseBody{
	Code:    0,
	Message: "success",
}
