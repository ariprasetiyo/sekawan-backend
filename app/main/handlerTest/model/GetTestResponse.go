package handlerTest

import (
	"sekawan-backend/app/main/handler"
)

type GetTestResponse struct {
	handler.Response
	Body GetTestBodyResponse `json:"body"`
}

type GetTestBodyResponse struct {
	Name   string `json:"name"`
	UserId string `json:"userId"`
	Acl    string `json:"acl"`
}
