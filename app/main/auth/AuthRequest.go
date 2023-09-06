package auth

import (
	"sekawan-backend/app/main/handler"
)

type AuthRequest struct {
	handler.Request
	body AuthBodyRequest `json:"body"`
}

type AuthBodyRequest struct {
	userId   string `json:"userId"`
	password string `json:"password"`
}
