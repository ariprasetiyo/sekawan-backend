package handlerAuth

import (
	"sekawan-backend/app/main/handler"
)

type AuthResponse struct {
	handler.Response
	Body AuthBodyResponse `json:"body"`
}

type AuthBodyResponse struct {
	Token  string `json:"token"`
	UserId string `json:"userId"`
}
