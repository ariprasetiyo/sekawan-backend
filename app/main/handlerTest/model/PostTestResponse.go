package handlerTest

import (
	"sekawan-backend/app/main/handler"
)

type PostTestResponse struct {
	handler.Response
	Body *PostTestBodyResponse `json:"body,omitempty"`
}

type PostTestBodyResponse struct {
	Name   string `json:"name,omitempty"`
	UserId string `json:"userId,omitempty"`
	Acl    string `json:"acl,omitempty"`
}
