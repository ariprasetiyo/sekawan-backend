package handlerAuth

import (
	"sekawan-backend/app/main/handler"
)

type AuthRequest struct {
	handler.Request
	body AuthBodyRequest `json:"body"`
}

/*
cred : json authModel in encryption AES 256
*/
type AuthBodyRequest struct {
	cred string `json:"cred"`
}
