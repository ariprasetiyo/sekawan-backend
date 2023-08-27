package handle_auth

import (
	"sekawan-backend/app/main/handler"
)

type AuthTokenRequest struct {
	handler.Request
	Body Body `json:"body"`
}

type Body struct {
}
