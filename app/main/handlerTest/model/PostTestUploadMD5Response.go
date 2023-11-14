package handlerTest

import (
	"sekawan-backend/app/main/handler"
)

type PostTestUploadMD5Response struct {
	handler.Response
	Body *PostTestUploadMD5BodyResponse `json:"body,omitempty"`
}

type PostTestUploadMD5BodyResponse struct {
	PlateNumber string `json:"plateNumber,omitempty"`
}
