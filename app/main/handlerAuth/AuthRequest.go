package handlerAuth

import (
	"sekawan-backend/app/main/handler"
)

type AuthRequest struct {
	handler.Request
	Body AuthBodyRequest `json:"body"`
}

/*
cred : json in encryption AES 256
*/
type AuthBodyRequest struct {
	Cred string `json:"cred"`
}

type AuthCredRequest struct {
	FullName string `json:"fullName"`
	PhoneNo  string `json:"phoneNo"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
