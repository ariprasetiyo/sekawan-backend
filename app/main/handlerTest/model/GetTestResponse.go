package handlerTest

import (
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/handlerAuth"
)

type GetTestResponse struct {
	handler.Response
	Body GetTestBodyResponse `json:"body"`
}

type GetTestBodyResponse struct {
	JwtDecoding handlerAuth.JWTToken `json:"jwtDecoding"`
	Name        string               `json:"name"`
	UserId      string               `json:"userId"`
	Acl         string               `json:"acl"`
}
