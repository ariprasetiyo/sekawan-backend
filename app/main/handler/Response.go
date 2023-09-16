package handler

import "sekawan-backend/app/main/enum"

type Response struct {
	ResponseId      string         `json:"responseId"`
	Type            enum.REQ_TYPE  `json:"type"`
	ResponseCode    enum.RESP_CODE `json:"responseCode"`
	ResponseMessage string         `json:"responseMessage"`
}
