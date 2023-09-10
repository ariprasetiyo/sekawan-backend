package handlerAuth

import (
	"sekawan-backend/app/main/handler"
)

type AuthResponse struct {
	handler.Response
	body AuthBodyResponse `json:"body"`
}

type AuthBodyResponse struct {
	token string `json:"token"`
}
