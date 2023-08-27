package handler

import "time"

type Request struct {
	RequestId   string `json:"requestId"`
	Type        string `json:"type"`
	RequestTime int64  `json:"requestTime"`
}

func (rm *Request) RequestModelDefault() {
	if rm.RequestTime == 0 {
		rm.RequestTime = time.Now().UnixMilli()
	}
}
