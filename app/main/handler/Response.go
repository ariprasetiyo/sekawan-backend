package handler

import "time"

type Response struct {
	ResponseId  string `json:"responseId"`
	Type        string `json:"type"`
	RequestTime int64  `json:"requestTime"`
}

func (rm *Request) ResponsetModelDefault() {
	if rm.RequestTime == 0 {
		rm.RequestTime = time.Now().UnixMilli()
	}
}
