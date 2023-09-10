package handler

type Response struct {
	ResponseId      string         `json:"responseId"`
	Type            REQ_TYPE_ENUM  `json:"type"`
	ResponseCode    RESP_CODE_ENUM `json:"responseCode"`
	ResponseMessage string         `json:"responseMessage"`
}
