package handlerTest

import "sekawan-backend/app/main/handler"

type PostTestUploadMD5Request struct {
	handler.Request
	Body PostTestUploadMD5BodyRequest `json:"body"`
}

type PostTestUploadMD5BodyRequest struct {
	Base64Image string `json:"base64Image"`
}
