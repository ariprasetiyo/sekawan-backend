package handler

import (
	"sekawan-backend/app/main/enum"
	"sekawan-backend/app/main/util"
	"time"

	"github.com/google/uuid"
)

type Request struct {
	RequestId   string        `json:"requestId"`
	Type        enum.REQ_TYPE `json:"type"`
	RequestTime int64         `json:"requestTime"`
}

func (rm *Request) RequestModelDefault() {
	if rm.RequestTime == 0 {
		rm.RequestTime = time.Now().UnixMilli()
	}

	if util.IsEmptyObject(rm.RequestId) || util.IsEmptyString(rm.RequestId) {
		rm.RequestId = uuid.New().String()
	}
}
