package handlerTest

import "sekawan-backend/app/main/handler"

type PostTestRequest struct {
	handler.Request
	Body PostTestBodyRequest `json:"body"`
}

type PostTestBodyRequest struct {
	Name string `json:"name"`
}
