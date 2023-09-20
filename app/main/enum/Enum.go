package enum

import (
	"encoding/json"
	"fmt"
	"strings"
)

type REQ_TYPE int32
type RESP_CODE int32

const (
	// enum request and response type
	TYPE_GENERATE_TOKEN         REQ_TYPE = 200
	TYPE_REQUEST_OCR            REQ_TYPE = 201
	TYPE_REQUEST_HTTP_GET_TEST  REQ_TYPE = 202
	TYPE_REQUEST_HTTP_POST_TEST REQ_TYPE = 203

	// generic response code
	SUCCESS       RESP_CODE = 200
	FAILED        RESP_CODE = 301
	SETTLED       RESP_CODE = 302
	PAID          RESP_CODE = 303
	UNKNOWN_ERROR RESP_CODE = 304
	BAD_REQUEST   RESP_CODE = 305
	UNAUTHORIZED  RESP_CODE = 306

	// response code auth
	AUTH_ERROR_DESERIALIZE_JSON_REQUEST RESP_CODE = 505
	AUTH_ERROR_INVALID_CLIENT_ID        RESP_CODE = 506
	AUTH_ERROR_INVALID_MSG_ID           RESP_CODE = 507
)

var (
	STRING_TO_REQ_TYPE = map[string]REQ_TYPE{
		"TYPE_GENERATE_TOKEN":         TYPE_GENERATE_TOKEN,
		"TYPE_REQUEST_OCR":            TYPE_REQUEST_OCR,
		"TYPE_REQUEST_HTTP_GET_TEST":  TYPE_REQUEST_HTTP_GET_TEST,
		"TYPE_REQUEST_HTTP_POST_TEST": TYPE_REQUEST_HTTP_POST_TEST,
	}

	REQ_TYPE_TO_STRING = map[REQ_TYPE]string{
		TYPE_GENERATE_TOKEN:         "TYPE_GENERATE_TOKEN",
		TYPE_REQUEST_OCR:            "TYPE_REQUESTS_OCR",
		TYPE_REQUEST_HTTP_GET_TEST:  "TYPE_REQUESTS_HTTP_GET_TEST",
		TYPE_REQUEST_HTTP_POST_TEST: "TYPE_REQUESTS_HTTP_POST_TEST",
	}

	STRING_TO_RESPONSE_CODE = map[string]RESP_CODE{
		"SUCCESS":                             SUCCESS,
		"FAILED":                              FAILED,
		"SETTLED":                             SETTLED,
		"PAID":                                PAID,
		"UNKNOWN_ERROR":                       UNKNOWN_ERROR,
		"BAD_REQUEST":                         BAD_REQUEST,
		"UNAUTHORIZED":                        UNAUTHORIZED,
		"AUTH_ERROR_DESERIALIZE_JSON_REQUEST": AUTH_ERROR_DESERIALIZE_JSON_REQUEST,
		"AUTH_ERROR_INVALID_CLIENT_ID":        AUTH_ERROR_INVALID_CLIENT_ID,
		"AUTH_ERROR_INVALID_MSG_ID":           AUTH_ERROR_INVALID_MSG_ID,
	}

	RESPONSE_CODE_TO_STRING = map[RESP_CODE]string{
		SUCCESS:                             "SUCCESS",
		FAILED:                              "FAILED",
		SETTLED:                             "SETTLED",
		PAID:                                "PAID",
		UNKNOWN_ERROR:                       "UNKNOWN_ERROR",
		BAD_REQUEST:                         "BAD_REQUEST",
		UNAUTHORIZED:                        "UNAUTHORIZED",
		AUTH_ERROR_DESERIALIZE_JSON_REQUEST: "AUTH_ERROR_DESERIALIZE_JSON_REQUEST",
		AUTH_ERROR_INVALID_CLIENT_ID:        "AUTH_ERROR_INVALID_CLIENT_ID",
		AUTH_ERROR_INVALID_MSG_ID:           "AUTH_ERROR_INVALID_MSG_ID",
	}
)

func (reqType REQ_TYPE) String() string {
	c, _ := REQ_TYPE_TO_STRING[reqType]
	return c
}

func (respCode RESP_CODE) String() string {
	c, _ := RESPONSE_CODE_TO_STRING[respCode]
	return c
}

func (s REQ_TYPE) MarshalJSON() ([]byte, error) {
	// It is assumed Suit implements fmt.Stringer.
	return json.Marshal(s.String())
}

// UnmarshalJSON must be a *pointer receiver* to ensure that the indirect from the
// parsed value can be set on the unmarshaling object. This means that the
// ParseSuit function must return a *value* and not a pointer.
func (s *REQ_TYPE) UnmarshalJSON(data []byte) (err error) {
	var suits string
	if err := json.Unmarshal(data, &suits); err != nil {
		return err
	}
	if *s, err = ParseSuit(suits); err != nil {
		return err
	}
	return nil
}

func ParseSuit(s string) (REQ_TYPE, error) {
	s = strings.TrimSpace(strings.ToUpper(s))
	value, ok := STRING_TO_REQ_TYPE[s]
	if !ok {
		return REQ_TYPE(0), fmt.Errorf("%q is not a valid card suit", s)
	}
	return REQ_TYPE(value), nil
}
