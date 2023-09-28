package handlerTest

import (
	"sekawan-backend/app/main/handler"
	"sekawan-backend/app/main/handlerAuth"
)

type PostTestResponse struct {
	handler.Response
	Body *PostTestBodyResponse `json:"body,omitempty"`
}

type PostTestBodyResponse struct {
	JwtDecoding *handlerAuth.JWTToken `json:"jwtDecoding,omitempty"`
	Name        string                `json:"name,omitempty"`
	UserId      string                `json:"userId,omitempty"`
	Acl         string                `json:"acl,omitempty"`
}
